# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).



## [Unreleased]

### Changed

- Update cluster-api to v1.0.5.

## [0.4.0] - 2021-09-23

### Changed

- Updated capi to v0.4.2

## [0.3.0] - 2021-01-19

### Added

- Check functions for `MachinePool` `ReplicasReady` condition

## [0.2.0] - 2020-12-04

### Added

- Add `UpgradePending` and `ControlPlaneReferenceNotSet` condition reasons.
- Update code docs.

## [0.1.0] - 2020-11-30

- Rename template repo.
- Add provider-independent conditions: Creating, Upgrading, InfrastructureReady, ControlPlaneReady, NodePoolsReady
- Add Ready condition check functions

[Unreleased]: https://github.com/giantswarm/conditions/compare/v0.4.0...HEAD
[0.4.0]: https://github.com/giantswarm/conditions/compare/v0.3.0...v0.4.0
[0.3.0]: https://github.com/giantswarm/conditions/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/giantswarm/conditions/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/giantswarm/conditions/releases/tag/v0.1.0
