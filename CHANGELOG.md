# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com), and this
project adheres to [Semantic Versioning](https://semver.org).

## [Unreleased]

### Added

* Extension recommendations
* Workspace settings
* `.editorconfig` that matches the Go language
* `.markdownlint.jsonc` that supports this changelog

### Changed

* `.gitignore` now includes settings for Go, Linux, MacOS, Windows and VSCode

### Deprecated

### Removed

* `main` branch was removed

### Fixed

### Security

* The token is still in plain text inside `main.go`; removing it is of utmost priority

## Guiding Principles

* Changelogs are for humans, not machines.
* There should be an entry for every single version.
* The same types of changes should be grouped.
* Versions and sections should be linkable.
* The latest version comes first.
* The release date of each version is displayed.
* Mention whether you follow Semantic Versioning.

## Types of changes

* `Added` for new features.
* `Changed` for changes in existing functionality.
* `Deprecated` for soon-to-be removed features.
* `Removed` for now removed features.
* `Fixed` for any bug fixes.
* `Security` in case of vulnerabilities.
