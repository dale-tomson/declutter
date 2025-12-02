# Declutter Build Scripts

This folder contains all the build and release scripts for Declutter.

## Quick Start

From the project root, use the wrapper script:

```bash
./run.sh <command> [options]
```

## Available Commands

### Check/Install Dependencies

```bash
# Check if all dependencies are installed
./run.sh deps check

# Install all dependencies
./run.sh deps install

# Install specific dependency
./run.sh deps go      # Install Go
./run.sh deps fyne    # Install Fyne CLI
./run.sh deps mingw   # Install MinGW-w64 (Windows cross-compilation)
./run.sh deps linux   # Install Linux build dependencies
```

### Version Bumping

```bash
# Show current version
./run.sh bump

# Bump version
./run.sh bump patch   # 1.0.0 -> 1.0.1
./run.sh bump minor   # 1.0.0 -> 1.1.0
./run.sh bump major   # 1.0.0 -> 2.0.0

# Set specific version
./run.sh bump 2.0.0-beta.1
```

### Building

```bash
# Build for specific platform
./run.sh build linux-amd64    # Linux x86_64
./run.sh build linux-arm64    # Linux ARM64
./run.sh build linux          # Linux (both)
./run.sh build mac-amd64      # macOS Intel
./run.sh build mac-arm64      # macOS Apple Silicon
./run.sh build mac            # macOS (both)
./run.sh build windows        # Windows x86_64
./run.sh build all            # All platforms

# Build options
./run.sh build linux --no-archive   # Don't create tar.gz/zip
./run.sh build all --clean          # Clean dist folder first
```

### Releasing

```bash
# Full release (build + tag + GitHub release)
./run.sh release

# Release options
./run.sh release --skip-build     # Use existing dist files
./run.sh release --skip-tag       # Don't create git tag
./run.sh release --draft          # Create draft release
./run.sh release --dry-run        # Preview what would happen
./run.sh release --version 1.2.0  # Override version
```

## Cross-Compilation Notes

### Building on Linux

| Target Platform | Requirements |
|-----------------|--------------|
| Linux amd64     | Native build |
| Linux arm64     | Requires ARM64 toolchain or native ARM64 machine |
| Windows amd64   | Requires `mingw-w64` (`./run.sh deps mingw`) |
| macOS           | **Not possible** - requires macOS |

### Building on macOS

| Target Platform | Requirements |
|-----------------|--------------|
| macOS amd64     | Native or Rosetta |
| macOS arm64     | Native or cross-compile |
| Linux           | Requires Docker or VM |
| Windows         | Requires `mingw-w64` (brew install mingw-w64) |

### Building on Windows

For Windows, we recommend using WSL2 with Ubuntu and following the Linux instructions.

## File Structure

```
scripts/
├── README.md                  # This file
├── buildDependencyChecker.sh  # Dependency checking/installation
├── bump.sh                    # Version bumping
├── build.sh                   # Platform-specific builds
└── release.sh                 # Full release automation
```

## Output

Built files are placed in the `dist/` folder:

```
dist/
├── declutter-1.0.0-linux-amd64.tar.gz
├── declutter-1.0.0-linux-arm64.tar.gz
├── declutter-1.0.0-darwin-amd64.tar.gz
├── declutter-1.0.0-darwin-arm64.tar.gz
└── declutter-1.0.0-windows-amd64.zip
```
