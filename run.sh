#!/bin/bash

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
SCRIPTS_DIR="$SCRIPT_DIR/scripts"

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

show_help() {
    echo -e "${BLUE}Declutter${NC} - File Organization Tool"
    echo ""
    echo "Usage: ./run.sh <command> [options]"
    echo ""
    echo "Commands:"
    echo "  ${GREEN}deps${NC} [check|install|go|fyne|mingw|linux]"
    echo "        Check or install build dependencies"
    echo ""
    echo "  ${GREEN}bump${NC} [major|minor|patch|<version>]"
    echo "        Bump or set the version number"
    echo ""
    echo "  ${GREEN}build${NC} <platform> [--no-archive] [--clean]"
    echo "        Build for specific platform(s)"
    echo "        Platforms: linux, linux-amd64, linux-arm64,"
    echo "                   mac, mac-amd64, mac-arm64, windows, all"
    echo ""
    echo "  ${GREEN}release${NC} [--skip-build] [--skip-tag] [--draft] [--dry-run]"
    echo "        Create a new release"
    echo ""
    echo "  ${GREEN}run${NC}"
    echo "        Build and run the application locally"
    echo ""
    echo "  ${GREEN}test${NC}"
    echo "        Run all tests"
    echo ""
    echo "  ${GREEN}clean${NC}"
    echo "        Remove build artifacts"
    echo ""
    echo "Examples:"
    echo "  ./run.sh deps check          # Check dependencies"
    echo "  ./run.sh bump patch          # Bump patch version"
    echo "  ./run.sh build linux         # Build for Linux"
    echo "  ./run.sh build all --clean   # Build all platforms"
    echo "  ./run.sh release --dry-run   # Preview release"
    echo ""
    echo "For more details, see: scripts/README.md"
}

case "${1:-help}" in
    deps)
        shift
        "$SCRIPTS_DIR/buildDependencyChecker.sh" "${@:-check}"
        ;;
    bump)
        shift
        "$SCRIPTS_DIR/bump.sh" "$@"
        ;;
    build)
        shift
        "$SCRIPTS_DIR/build.sh" "$@"
        ;;
    release)
        shift
        "$SCRIPTS_DIR/release.sh" "$@"
        ;;
    run)
        echo -e "${BLUE}Building and running Declutter...${NC}"
        go build -o "$SCRIPT_DIR/declutter" .
        "$SCRIPT_DIR/declutter"
        ;;
    test)
        echo -e "${BLUE}Running tests...${NC}"
        go test ./... -v
        ;;
    clean)
        echo -e "${YELLOW}Cleaning build artifacts...${NC}"
        rm -rf "$SCRIPT_DIR/dist"
        rm -f "$SCRIPT_DIR/declutter"
        rm -f "$SCRIPT_DIR/declutter.exe"
        echo -e "${GREEN}✓${NC} Cleaned"
        ;;
    help|-h|--help)
        show_help
        ;;
    *)
        echo -e "${RED}Unknown command: $1${NC}"
        echo ""
        show_help
        exit 1
        ;;
esac
