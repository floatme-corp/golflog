# Changelog

## [1.5.0](https://github.com/floatme-corp/golflog/compare/v1.4.0...v1.5.0) (2022-04-27)


### Features

* **log:** add `WarnWrap` helper ([e77b76b](https://github.com/floatme-corp/golflog/commit/e77b76b1aa1fa39a2b8ef5697ff4fbd74acceb14))


### Bug Fixes

* **devkit:** set safe directory ([83ab80a](https://github.com/floatme-corp/golflog/commit/83ab80a985270b02ee1934579490a892d568641e))
* **make:** use tabs for make indentation ([417510b](https://github.com/floatme-corp/golflog/commit/417510bc2258c03582523b01ca032141512ace08))


### Miscellaneous

* **actions:** rename commits.yml -> commits.yaml ([5c472b7](https://github.com/floatme-corp/golflog/commit/5c472b7d4c21e616d02ee417be3cbb0d4da0fec9))
* **actions:** use service account for release-please ([4ac251a](https://github.com/floatme-corp/golflog/commit/4ac251a5a07a6475c78026c2fc3777f6992d0165))
* **docker:** bump golang from 1.17.8-alpine to 1.18.0-alpine ([bce3761](https://github.com/floatme-corp/golflog/commit/bce37613e270e27bed0e007eaff0d3e9ddfa6ec1))
* **docker:** bump golang from 1.18.0-alpine to 1.18.1-alpine ([ada7218](https://github.com/floatme-corp/golflog/commit/ada72187a30b2f79a2f350583c5edf0ede7f47f3))
* **docker:** bump golangci/golangci-lint ([36cede7](https://github.com/floatme-corp/golflog/commit/36cede7dae54e9caab4d178248d6f4d94ff1ef06))
* **docker:** bump hadolint/hadolint ([6bed679](https://github.com/floatme-corp/golflog/commit/6bed6799d53981ed5b42384cae077b0a47871779))
* **docker:** bump vektra/mockery from v2.10.0 to v2.10.4 ([a1780df](https://github.com/floatme-corp/golflog/commit/a1780dfe142a624f4f07049e460b53b99a7982e1))
* **docker:** bump vektra/mockery from v2.10.4 to v2.12.1 ([a76da71](https://github.com/floatme-corp/golflog/commit/a76da71a02082a06dd83b2f1fea06d1ef50f02fa))
* **go:** bump github.com/spf13/viper from 1.10.1 to 1.11.0 ([94275a1](https://github.com/floatme-corp/golflog/commit/94275a1f51641dc5676bce9669bf8d086eb6e261))
* **tools:** remove cobertura and action ([ab9d684](https://github.com/floatme-corp/golflog/commit/ab9d6843eecd3795b4410b3231e4066ced5836aa))
* **tools:** use coveralls ([4da485d](https://github.com/floatme-corp/golflog/commit/4da485db603091f61c0afe8bedcd5bcfec670f4a))

## [1.4.0](https://github.com/floatme-corp/golflog/compare/v1.3.0...v1.4.0) (2022-03-23)


### Features

* **log:** add `logr.Discard` proxy ([fbb3ffd](https://github.com/floatme-corp/golflog/commit/fbb3ffd6c821115546b993b4050a6642a0bc196d))


### Miscellaneous

* **docker:** bump golangci/golangci-lint ([42e477e](https://github.com/floatme-corp/golflog/commit/42e477ef50100128431bf99603447a554dabefab))
* **docker:** bump hadolint/hadolint ([c280e8f](https://github.com/floatme-corp/golflog/commit/c280e8f8121ecfdc1e6c5a8eb7b29ea49ff5ea71))
* **go:** bump github.com/go-logr/logr from 1.2.2 to 1.2.3 ([21e971c](https://github.com/floatme-corp/golflog/commit/21e971c121914e6622f272cff61b67e1fa9f6b45))

## [1.3.0](https://github.com/floatme-corp/golflog/compare/v1.2.0...v1.3.0) (2022-03-16)


### Features

* **context:** add `ContextWithNameAndValues` helper ([7c598a2](https://github.com/floatme-corp/golflog/commit/7c598a2f780e0bcf4e01a8e17f2f48565a05c2ff))
* **log:** add `Debug` helpers ([e2982c3](https://github.com/floatme-corp/golflog/commit/e2982c3ad268c81aed82ed2372d62e8006bb1561))
* **log:** add `V` helper ([8cc6020](https://github.com/floatme-corp/golflog/commit/8cc602041f894e8e2079eb674d1f0da8cb6733ba))
* **log:** add `Warn`/`Warning` helpers ([bf2ff7c](https://github.com/floatme-corp/golflog/commit/bf2ff7c6bec829b284f5417262754f98c179d6b1))


### Bug Fixes

* **log:** add `severity` to `Error` ([f1af812](https://github.com/floatme-corp/golflog/commit/f1af8129a4e35a335a22100d59539446e84e7541))
* **log:** add `severity` to `Info` ([3695120](https://github.com/floatme-corp/golflog/commit/3695120c355c30e8bc8ce24ef973e17d363dcfb1))
* **log:** use `WithCallStackHelper` ([eb48817](https://github.com/floatme-corp/golflog/commit/eb488178237f8ae734b4da9a6f2895c6acf87186))
* **readme:** fix spelling ([4b2787d](https://github.com/floatme-corp/golflog/commit/4b2787d0ca9334856142a6e219feb57906ec36dd))
* **test:** fix `Error` test name ([8720459](https://github.com/floatme-corp/golflog/commit/8720459bf0004d933c08ff299a8077684d443ec6))


### Miscellaneous

* **readme:** add `severity` documentation ([5c50581](https://github.com/floatme-corp/golflog/commit/5c50581313448e720891546ee25254af751b6aa0))

## [1.2.0](https://github.com/floatme-corp/golflog/compare/v1.1.0...v1.2.0) (2022-03-15)


### Features

* **context:** add `WithName`, `WithValues`, and `WithNameAndValues` ([7ed8dcb](https://github.com/floatme-corp/golflog/commit/7ed8dcbb38b12f80e760cb39f2e511b160b77365))
* **log:** add `Error` helper ([90f6969](https://github.com/floatme-corp/golflog/commit/90f6969633bd68bac156dbce125c2a4550a03a46))


### Miscellaneous

* **doc:** add missing `"` to README.md ([4ea8838](https://github.com/floatme-corp/golflog/commit/4ea883842cd67d048ac213d2cef13f1ae65ae4dd))
* **log:** fix docstring ([26af93a](https://github.com/floatme-corp/golflog/commit/26af93ac87d1a8ae87e7d7b09a16c87e87ae772d))
* **test:** add test for `Info` and `Wrap` ([3871e19](https://github.com/floatme-corp/golflog/commit/3871e19c52e49acdb72e79f2c19f52543af3d2ae))

## [1.1.0](https://github.com/floatme-corp/golflog/compare/v1.0.0...v1.1.0) (2022-03-14)


### Features

* add log.wrap and log.info to goflog ([01eb20c](https://github.com/floatme-corp/golflog/commit/01eb20cbeb7daf1a2942d7bb1e848ff85683b2fd))


### Bug Fixes

* go string formatting ([fe32c77](https://github.com/floatme-corp/golflog/commit/fe32c77661aeea5bdedc071223c8a6f3edb17dee))
* **mod:** add stub for mock package ([b515539](https://github.com/floatme-corp/golflog/commit/b515539f8a8c7f7337d62c4a1a305c1039c28c7a))


### Miscellaneous

* **doc:** add highlighting to README.md ([d12604e](https://github.com/floatme-corp/golflog/commit/d12604ead64552be82b37fefa8d53cf3e9164b89))

## 1.0.0 (2022-03-12)


### Features

* initial release ([b1218ea](https://github.com/floatme-corp/golflog/commit/b1218ea89348f42165467e85fcb5aae9f93dec48))


### Miscellaneous

* **github:** bump actions/checkout from 2 to 3 ([b7d1436](https://github.com/floatme-corp/golflog/commit/b7d14361f0834a002fb83faa53af8ffdb640c1ea))
* **readme:** fix badges ([5cf3cf3](https://github.com/floatme-corp/golflog/commit/5cf3cf305d30ed47944f1e8e584fb43daf3cc4f2))
* **readme:** grammer ([dd2e78a](https://github.com/floatme-corp/golflog/commit/dd2e78a099eb798e4ad1471d27132c603e43e52c))
* **release:** unhide chores ([439f121](https://github.com/floatme-corp/golflog/commit/439f1216dd6b1c7faa9d040ee5734997d252595e))
