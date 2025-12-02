#!/bin/bash

set -e

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"

check_command() {
    if command -v "$1" &> /dev/null; then
        echo -e "${GREEN}✓${NC} $1 found: $(command -v "$1")"
        return 0
    else
        echo -e "${RED}✗${NC} $1 not found"
        return 1
    fi
}

install_go() {
    echo -e "${YELLOW}Installing Go...${NC}"
    if [[ "$OSTYPE" == "linux-gnu"* ]]; then
        if command -v apt-get &> /dev/null; then
            sudo apt-get update && sudo apt-get install -y golang-go
        elif command -v dnf &> /dev/null; then
            sudo dnf install -y golang
        elif command -v pacman &> /dev/null; then
            sudo pacman -S --noconfirm go
        else
            echo -e "${RED}Please install Go manually: https://go.dev/dl/${NC}"
            return 1
        fi
    elif [[ "$OSTYPE" == "darwin"* ]]; then
        if command -v brew &> /dev/null; then
            brew install go
        else
            echo -e "${RED}Please install Homebrew first or install Go manually${NC}"
            return 1
        fi
    fi
}

install_fyne_cli() {
    echo -e "${YELLOW}Installing Fyne CLI...${NC}"
    go install fyne.io/fyne/v2/cmd/fyne@latest
}

install_mingw() {
    echo -e "${YELLOW}Installing MinGW-w64 for Windows cross-compilation...${NC}"
    if [[ "$OSTYPE" == "linux-gnu"* ]]; then
        if command -v apt-get &> /dev/null; then
            sudo apt-get update && sudo apt-get install -y mingw-w64
        elif command -v dnf &> /dev/null; then
            sudo dnf install -y mingw64-gcc
        elif command -v pacman &> /dev/null; then
            sudo pacman -S --noconfirm mingw-w64-gcc
        else
            echo -e "${RED}Please install mingw-w64 manually${NC}"
            return 1
        fi
    elif [[ "$OSTYPE" == "darwin"* ]]; then
        if command -v brew &> /dev/null; then
            brew install mingw-w64
        else
            echo -e "${RED}Please install Homebrew first${NC}"
            return 1
        fi
    fi
}

install_linux_deps() {
    echo -e "${YELLOW}Installing Linux build dependencies...${NC}"
    if command -v apt-get &> /dev/null; then
        sudo apt-get update && sudo apt-get install -y \
            libgl1-mesa-dev \
            xorg-dev \
            libxcursor-dev \
            libxrandr-dev \
            libxinerama-dev \
            libxi-dev \
            libxxf86vm-dev
    elif command -v dnf &> /dev/null; then
        sudo dnf install -y \
            mesa-libGL-devel \
            libX11-devel \
            libXcursor-devel \
            libXrandr-devel \
            libXinerama-devel \
            libXi-devel
    elif command -v pacman &> /dev/null; then
        sudo pacman -S --noconfirm \
            mesa \
            libx11 \
            libxcursor \
            libxrandr \
            libxinerama \
            libxi
    fi
}

check_all() {
    echo "🔍 Checking build dependencies..."
    echo "================================"
    echo ""
    
    local missing=()
    
    if ! check_command "go"; then
        missing+=("go")
    else
        echo "   Version: $(go version)"
    fi
    
    if ! check_command "fyne"; then
        if command -v "$(go env GOPATH)/bin/fyne" &> /dev/null; then
            echo -e "${GREEN}✓${NC} fyne found: $(go env GOPATH)/bin/fyne"
        else
            missing+=("fyne")
        fi
    fi
    
    echo ""
    echo "Cross-compilation tools:"
    
    if ! check_command "x86_64-w64-mingw32-gcc"; then
        echo -e "${YELLOW}  (Optional: needed for Windows builds on Linux/macOS)${NC}"
    fi
    
    if [[ "$OSTYPE" == "linux-gnu"* ]]; then
        echo ""
        echo "Linux build dependencies:"
        if pkg-config --exists gl 2>/dev/null; then
            echo -e "${GREEN}✓${NC} OpenGL development libraries"
        else
            echo -e "${YELLOW}⚠${NC} OpenGL development libraries may be missing"
            missing+=("linux-deps")
        fi
    fi
    
    echo ""
    
    if [ ${#missing[@]} -eq 0 ]; then
        echo -e "${GREEN}✅ All required dependencies are installed!${NC}"
        return 0
    else
        echo -e "${YELLOW}Missing dependencies: ${missing[*]}${NC}"
        return 1
    fi
}

install_all() {
    echo "📦 Installing all build dependencies..."
    echo "======================================="
    echo ""
    
    if ! command -v go &> /dev/null; then
        install_go
    fi
    
    if [[ "$OSTYPE" == "linux-gnu"* ]]; then
        install_linux_deps
    fi
    
    install_fyne_cli
    
    echo ""
    read -p "Install MinGW-w64 for Windows cross-compilation? [y/N] " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        install_mingw
    fi
    
    echo ""
    echo -e "${GREEN}✅ Dependencies installed!${NC}"
}

case "${1:-check}" in
    check)
        check_all
        ;;
    install)
        install_all
        ;;
    go)
        install_go
        ;;
    fyne)
        install_fyne_cli
        ;;
    mingw)
        install_mingw
        ;;
    linux)
        install_linux_deps
        ;;
    *)
        echo "Declutter Build Dependency Checker"
        echo ""
        echo "Usage: $0 [command]"
        echo ""
        echo "Commands:"
        echo "  check    Check if all dependencies are installed (default)"
        echo "  install  Install all missing dependencies"
        echo "  go       Install Go"
        echo "  fyne     Install Fyne CLI"
        echo "  mingw    Install MinGW-w64 (Windows cross-compilation)"
        echo "  linux    Install Linux build dependencies"
        ;;
esac
