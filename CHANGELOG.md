# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.1.1] - 2023-03-27

### Added

- Use the `catkin` tag inside C++ configuration to make `rtcg-gen` produce
  Catkin `package.xml` and `CMakeLists.txt` files for each test.  This is
  experimental.  Inside the `catkin` tag, the `package` sub-tag contains a
  simplified implementation of the Catkin v2 package format.

### Changed

- STM suites now contain unified inferred channel types as well as the
  generated tests.  This means there is now an extra level of nesting in the
  data structure: tests are under `tests`, types are under `types` (and are
  a map from channel names to type info).  See `examples/bmon/stms.json`.
- There is now only one `convert.h` generated, and it uses the unified types
  stored in the STM.
- Both `convert.cpp` and the newly unified `convert.h` now appear in the
  `convert` subdirectory of the source root (alongside `rtcg` etc).
- Input directory structures now more closely match output directory
  structures.  For example, `convert.cpp` must now be in
  `$INPUT/$VARIANT/src/convert/convert.cpp`, which matches where it will be
  generated in the output.

## [0.1.0] - 2023-03-14

### Added

- Initial release, with support for basic ROS1 C++ code generation from
  forbidden traces.

[unreleased]: https://github.com/UoY-RoboStar/rtcg/compare/v0.1.0...HEAD
[0.1.1]: https://github.com/UoY-RoboStar/rtcg/compare/v0.1.0...v0.1.1
[0.1.0]: https://github.com/UoY-RoboStar/rtcg/releases/tag/v0.1.0