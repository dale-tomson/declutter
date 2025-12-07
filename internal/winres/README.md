# Windows Resource Files

This directory contains configuration files for embedding Windows-specific metadata into the executable.

## Files

- **`winres.json`** - Configuration for Windows resources (version info, icon, company, description)
- **`app.manifest`** - Windows application manifest (execution level, DPI awareness, OS compatibility)
- **`icon.png`** - Application icon (100x80 PNG, embedded in the executable)

## Purpose

These files help reduce Windows Defender false positives by:
- Adding proper version information to the executable
- Identifying the publisher and application details
- Embedding the application icon
- Specifying execution requirements and OS compatibility

## Build Process

During Windows builds, the `go-winres` tool:
1. Reads `winres.json` and `app.manifest`
2. Generates a `.syso` resource file (`../rsrc_windows_amd64.syso`)
3. Go compiler automatically embeds the `.syso` file into the executable

## Updating Version

The version numbers in `winres.json` are automatically updated during the build process based on the version in `internal/version/version.go`.

## Learn More

- [go-winres documentation](https://github.com/tc-hib/go-winres)
- [Windows Defender troubleshooting guide](../../docs/WINDOWS_DEFENDER.md)
