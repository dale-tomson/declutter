#!/bin/bash

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

VERSION_FILE="$PROJECT_ROOT/internal/version/version.go"
VERSION=$(grep 'const Version' "$VERSION_FILE" | sed 's/.*"\(.*\)"/\1/')
APP_NAME="declutter"
DIST_DIR="$PROJECT_ROOT/dist"

show_help() {
    echo "Declutter Build Script"
    echo ""
    echo "Usage: $0 <platform> [options]"
    echo ""
    echo "Platforms:"
    echo "  linux-amd64    Build for Linux (x86_64)"
    echo "  linux-arm64    Build for Linux (ARM64)"
    echo "  linux          Build for Linux (both architectures)"
    echo "  mac-amd64      Build for macOS (Intel)"
    echo "  mac-arm64      Build for macOS (Apple Silicon)"
    echo "  mac            Build for macOS (both architectures)"
    echo "  windows        Build for Windows (x86_64)"
    echo "  all            Build for all platforms"
    echo ""
    echo "Options:"
    echo "  --no-archive   Don't create archives (tar.gz/zip)"
    echo "  --clean        Clean dist folder before building"
    echo ""
    echo "Current version: ${VERSION}"
}

check_deps() {
    if ! command -v go &> /dev/null; then
        echo -e "${RED}Error: Go is not installed${NC}"
        echo "Run: ./run.sh deps install"
        exit 1
    fi
}

build_binary() {
    local os="$1"
    local arch="$2"
    local output_name="${APP_NAME}-${VERSION}-${os}-${arch}"
    
    if [ "$os" = "windows" ]; then
        output_name="${output_name}.exe"
    fi
    
    echo ""
    echo -e "${YELLOW}📦 Building for ${os}/${arch}...${NC}"
    
    export GOOS="$os"
    export GOARCH="$arch"
    export CGO_ENABLED=1
    
    case "$os" in
        linux)
            if [ "$arch" = "arm64" ] && [ "$(uname -m)" != "aarch64" ]; then
                echo -e "   ${YELLOW}⚠️  Skipping linux/arm64 (cross-compilation requires arm64 toolchain)${NC}"
                return 1
            fi
            ;;
        darwin)
            if [ "$(uname)" != "Darwin" ]; then
                echo -e "   ${YELLOW}⚠️  Skipping darwin/${arch} (requires macOS)${NC}"
                return 1
            fi
            ;;
        windows)
            if [ "$(uname)" = "Linux" ]; then
                if ! command -v x86_64-w64-mingw32-gcc &> /dev/null; then
                    echo -e "   ${YELLOW}⚠️  Skipping windows/amd64 (install mingw-w64)${NC}"
                    echo "   Run: ./run.sh deps mingw"
                    return 1
                fi
                export CC=x86_64-w64-mingw32-gcc
            fi
            ;;
    esac
    
    cd "$PROJECT_ROOT"
    go build -ldflags="-s -w" -o "${DIST_DIR}/${output_name}" .
    
    if [ -f "${DIST_DIR}/${output_name}" ]; then
        echo -e "   ${GREEN}✅ Built: ${output_name}${NC}"
        return 0
    else
        echo -e "   ${RED}❌ Failed to build${NC}"
        return 1
    fi
}

create_archive() {
    local file="$1"
    
    cd "$DIST_DIR"
    
    if [[ "$file" == *.exe ]]; then
        local zip_name="${file%.exe}.zip"
        zip -q "$zip_name" "$file"
        rm "$file"
        echo -e "   ${GREEN}✅ Created: ${zip_name}${NC}"
    else
        local tar_name="${file}.tar.gz"
        tar -czf "$tar_name" "$file"
        rm "$file"
        echo -e "   ${GREEN}✅ Created: ${tar_name}${NC}"
    fi
    
    cd "$PROJECT_ROOT"
}

NO_ARCHIVE=false
CLEAN=false
PLATFORMS=()

while [[ $# -gt 0 ]]; do
    case "$1" in
        --no-archive)
            NO_ARCHIVE=true
            shift
            ;;
        --clean)
            CLEAN=true
            shift
            ;;
        -h|--help|help)
            show_help
            exit 0
            ;;
        linux-amd64)
            PLATFORMS+=("linux:amd64")
            shift
            ;;
        linux-arm64)
            PLATFORMS+=("linux:arm64")
            shift
            ;;
        linux)
            PLATFORMS+=("linux:amd64" "linux:arm64")
            shift
            ;;
        mac-amd64|darwin-amd64)
            PLATFORMS+=("darwin:amd64")
            shift
            ;;
        mac-arm64|darwin-arm64)
            PLATFORMS+=("darwin:arm64")
            shift
            ;;
        mac|darwin)
            PLATFORMS+=("darwin:amd64" "darwin:arm64")
            shift
            ;;
        windows|win)
            PLATFORMS+=("windows:amd64")
            shift
            ;;
        all)
            PLATFORMS+=("linux:amd64" "linux:arm64" "darwin:amd64" "darwin:arm64" "windows:amd64")
            shift
            ;;
        *)
            echo -e "${RED}Unknown option: $1${NC}"
            show_help
            exit 1
            ;;
    esac
done

if [ ${#PLATFORMS[@]} -eq 0 ]; then
    show_help
    exit 0
fi

check_deps

echo "🚀 Building Declutter v${VERSION}"
echo "================================"

if [ "$CLEAN" = true ] || [ ! -d "$DIST_DIR" ]; then
    rm -rf "$DIST_DIR"
fi
mkdir -p "$DIST_DIR"

BUILT_FILES=()

for platform in "${PLATFORMS[@]}"; do
    os="${platform%:*}"
    arch="${platform#*:}"
    
    if build_binary "$os" "$arch"; then
        if [ "$os" = "windows" ]; then
            BUILT_FILES+=("${APP_NAME}-${VERSION}-${os}-${arch}.exe")
        else
            BUILT_FILES+=("${APP_NAME}-${VERSION}-${os}-${arch}")
        fi
    fi
done

unset GOOS GOARCH CGO_ENABLED CC

if [ "$NO_ARCHIVE" = false ] && [ ${#BUILT_FILES[@]} -gt 0 ]; then
    echo ""
    echo "📁 Creating archives..."
    
    for file in "${BUILT_FILES[@]}"; do
        if [ -f "${DIST_DIR}/${file}" ]; then
            create_archive "$file"
        fi
    done
fi

echo ""
echo -e "${GREEN}🎉 Build complete!${NC}"
echo ""
echo "Files in ${DIST_DIR}/:"
ls -lh "$DIST_DIR" 2>/dev/null || echo "  (empty)"
