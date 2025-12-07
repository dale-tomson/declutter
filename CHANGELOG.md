# Changelog

All notable changes to Declutter will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed
- Refactored project structure: moved `main.go` to `cmd/declutter/main.go`
- Replaced shell scripts with `Makefile` (`make build`, `make run`, `make test`)
- Updated Go version to 1.25 in workflows
- Added UPX compression for Linux and Windows builds in release workflow

## [1.1.2] - 2025-12-07

### Fixed
- Folder selection state after organizing files - users can now select a new folder immediately after completion
- CI/CD workflow file handling for Windows resource generation

## [1.1.1] - 2025-12-07

### Added
- Comprehensive Windows Defender troubleshooting guide (`docs/WINDOWS_DEFENDER.md`)
- Windows resource metadata embedded in executables (version info, company, description)
- Application icon embedded in Windows executable (100x80 PNG)
- Windows application manifest (execution level, DPI awareness, OS compatibility)
- Automated Windows resource generation in build process

### Changed
- Windows executable now runs in GUI-only mode (no console window)
- Updated README with Windows Defender warning notice
- Build scripts now automatically generate Windows resources using `go-winres`

### Fixed
- Windows Defender false positive warnings addressed with proper metadata and user documentation

## [1.1.0] - 2025-12-02

### Added
- Mobile-friendly navigation with hamburger menu
- Responsive slideshow for "How It Works" section
- GitHub Actions CI workflow for automated testing
- GitHub Actions release workflow for manual releases
- Test coverage for icon package
- Separate changelog files per version in `changelogs/` folder

### Changed
- Moved CSS and JavaScript to separate files in docs
- Moved version display from header to footer in UI
- Updated README with screenshots and project structure

### Removed
- Unused `Result` struct from organizer package
- Unused exported wrapper functions (`EnsureDir`, `MoveFile`, `CopyFile`)
- Unused `GetYearPath` function
- Redundant `SetCloseIntercept` in main.go
- Duplicate `logo.svg` from project root

### Fixed
- Theme test assertions for color names

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

[Unreleased]: https://github.com/dale-tomson/declutter/compare/v1.1.2...HEAD
[1.1.2]: https://github.com/dale-tomson/declutter/compare/v1.1.1...v1.1.2
[1.1.1]: https://github.com/dale-tomson/declutter/compare/v1.1.0...v1.1.1
[1.1.0]: https://github.com/dale-tomson/declutter/compare/v1.0.0...v1.1.0
[1.0.0]: https://github.com/dale-tomson/declutter/releases/tag/v1.0.0
