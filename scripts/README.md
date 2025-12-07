# Build Scripts

Most build operations are now handled by the `Makefile` in the project root.

## Quick Start (Makefile)

```bash
make run    # Run locally
make build  # Build binary
make test   # Run tests
make clean  # Clean artifacts
make bump   # Bump version
```

## Helper Scripts

This folder contains internal helper scripts used by the Makefile or CI/CD files.

-   `bump.sh`: Logic for version bumping and changelog updates.
