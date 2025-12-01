# Changelog

All notable changes to Declutter will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.0.0] - 2025-12-01

### Added
- Initial release of Declutter
- File organization into Year/Month folders based on modification timestamps
- Cross-platform support (Windows, macOS, Linux)
- Modern UI with custom color theme
- Real-time activity log showing file operations
- Progress indicator during organization
- Confirmation dialog before organizing
- Skip existing files to prevent overwrites
- Graceful application close handling
- Custom app icon
- GitHub Pages website with OS detection for downloads

### Technical
- Built with Go and Fyne v2.7.1
- Separated architecture: organizer logic, theme, UI, and icon packages
- Comprehensive test suite (28+ tests)
- Thread-safe UI updates with fyne.Do()

[Unreleased]: https://github.com/dale-tomson/declutter/compare/v1.0.0...HEAD
[1.0.0]: https://github.com/dale-tomson/declutter/releases/tag/v1.0.0
