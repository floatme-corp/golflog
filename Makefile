#

SHELL := /bin/bash
INTERACTIVE := $(shell [ -t 0 ] && echo 1)

root_mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
export REPO_ROOT_DIR := $(realpath $(dir $(root_mkfile_path)))

export DOCKER_REPOSITORY ?= ghcr.io/floatme-corp/golflog
export DOCKER_DEVKIT_IMG ?= $(DOCKER_REPOSITORY):latest-devkit
export DOCKER_DEVKIT_PHONY_FILE ?= .docker-$(shell echo '$(DOCKER_DEVKIT_IMG)' | tr '/:' '.')

export DOCKER_DEVKIT_GITHUB_ARGS ?= \
	$(if $(GITHUB_EVENT_PATH),--volume $(GITHUB_EVENT_PATH):$(GITHUB_EVENT_PATH)) \
	--env CI \
	--env-file <(env | grep '^GITHUB_') \
	--env-file <(env | grep '^GIT_')

export DOCKER_DEVKIT_ARGS ?= \
	--rm \
	$(if $(INTERACTIVE),--tty) \
	--interactive \
	--env DEVKIT=true \
	--volume $(REPO_ROOT_DIR):/code:Z \
	--workdir /code \
	$(DOCKER_DEVKIT_GITHUB_ARGS)

# NOTE(jkoelker) Comma must not appear in a funciton call, use a variable
#                as suggested by the documentation.
#  https://www.gnu.org/software/make/manual/html_node/Syntax-of-Functions.html
comma := ,
export DOCKER_DEVKIT_BUILDX_ARGS ?= \
	$(if $(GITHUB_ACTIONS),--cache-from type=gha) \
	$(if $(GITHUB_ACTIONS),--cache-to type=gha$(comma)mode=max)

# NOTE(jkoelker) Abuse ifeq and the junk variable to proxy docker image state
#                to the target file
ifneq ($(shell command -v docker),)
    ifeq ($(shell docker image ls --quiet "$(DOCKER_DEVKIT_IMG)"),)
        export junk := $(shell rm -rf $(DOCKER_DEVKIT_PHONY_FILE))
    endif
endif

$(DOCKER_DEVKIT_PHONY_FILE): $(shell find $(REPO_ROOT_DIR)/tools -type f -name '*'.go)
$(DOCKER_DEVKIT_PHONY_FILE): Dockerfile.devkit
	docker buildx build \
		$(DOCKER_DEVKIT_BUILDX_ARGS) \
		--file $(REPO_ROOT_DIR)/Dockerfile.devkit \
		--tag "$(DOCKER_DEVKIT_IMG)" \
		$(REPO_ROOT_DIR) \
	&& touch $(DOCKER_DEVKIT_PHONY_FILE)

.PHONY: devkit
devkit: $(DOCKER_DEVKIT_PHONY_FILE)

WHAT ?= /bin/bash -l
.PHONY: devkit.run
devkit.run: devkit
	docker run \
		$(DOCKER_DEVKIT_ARGS) \
		"$(DOCKER_DEVKIT_IMG)" \
		/bin/bash -c \
			'git config --global safe.directory /code \
			&& $(WHAT)'

.PHONY: dev
dev: devkit.run

.PHONY: shell
shell: devkit.run

.PHONY: clean
clean: clean-coverage
clean: clean-mocks

.PHONY: clean-coverage
clean-coverage:
	rm -f coverage.out coverage.xml coverage.html

.PHONY: clean-mocks
clean-mocks:
	rm -r mocks/build_info.go mocks/configurator.go

coverage.out: $(shell find $(REPO_ROOT_DIR) -type f -name '*'.go)
coverage.out: generate.host
coverage.out:
	gotestsum \
		--junitfile-testsuite-name=relative \
		--junitfile-testcase-classname=short \
		-- \
		-covermode=atomic \
		-coverprofile=coverage.out \
		-race \
		-short \
		-v \
		./...

coverage.html: coverage.out
coverage.html:
	go tool cover -html=coverage.out -o coverage.html

.PHONY: goveralls
goveralls: coverage.out
goveralls:
	# NOTE(jkoelker) Shenanigans to securly shuttle the token to goveralls
	token="$$(mktemp)" && \
		printenv "GITHUB_TOKEN" > "$${token}" && \
		goveralls \
			-repotokenfile="$${token}" \
			-coverprofile=coverage.out \
			-service=github && \
		rm "$${token}"

.PHONY: lint.host-docker
lint.host-docker:
	@echo "Linting Dockerfiles"
	find $(REPO_ROOT_DIR)/ -name 'Dockerfile*' -exec hadolint {} +
	@echo

.PHONY: lint.host-go
lint.host-go: generate.host
	@echo "Linting Go files"
	golangci-lint run --verbose
	@echo

.PHONY: lint.host
lint.host: lint.host-docker
lint.host: lint.host-go

.PHONY: lint
lint: WHAT=make lint.host
lint: devkit.run

.PHONY: generate
generate: WHAT=make generate.host
generate: devkit.run

.PHONY: generate.host
generate.host: generate.host-mocks

.PHONY: generate.host-mocks
generate.host-mocks: mocks/configurator.go
generate.host-mocks: mocks/build_info.go

mocks/configurator.go: log.go
	go generate log.go

mocks/build_info.go: log.go
	go generate log.go

.PHONY: test
test: WHAT=make test.host
test: devkit.run

.PHONY: test.host
test.host: generate.host
test.host: coverage.html
