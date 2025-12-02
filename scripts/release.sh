#!/bin/bash

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

VERSION_FILE="$PROJECT_ROOT/internal/version/version.go"
VERSION=$(grep 'const Version' "$VERSION_FILE" | sed 's/.*"\(.*\)"/\1/')
DIST_DIR="$PROJECT_ROOT/dist"

show_help() {
    echo "Declutter Release Script"
    echo ""
    echo "Usage: $0 [options]"
    echo ""
    echo "Options:"
    echo "  --version <ver>  Override version (default: from version.go)"
    echo "  --skip-build     Skip building, use existing dist files"
    echo "  --skip-tag       Skip git tagging"
    echo "  --draft          Create draft release on GitHub"
    echo "  --dry-run        Show what would be done without executing"
    echo ""
    echo "Current version: ${VERSION}"
}

SKIP_BUILD=false
SKIP_TAG=false
DRAFT=false
DRY_RUN=false

while [[ $# -gt 0 ]]; do
    case "$1" in
        --version)
            VERSION="$2"
            shift 2
            ;;
        --skip-build)
            SKIP_BUILD=true
            shift
            ;;
        --skip-tag)
            SKIP_TAG=true
            shift
            ;;
        --draft)
            DRAFT=true
            shift
            ;;
        --dry-run)
            DRY_RUN=true
            shift
            ;;
        -h|--help|help)
            show_help
            exit 0
            ;;
        *)
            echo -e "${RED}Unknown option: $1${NC}"
            show_help
            exit 1
            ;;
    esac
done

echo -e "${BLUE}🚀 Declutter Release v${VERSION}${NC}"
echo "=============================="
echo ""

echo "📋 Pre-release checklist:"
echo "-------------------------"

CHANGELOG="$PROJECT_ROOT/CHANGELOG.md"
if grep -q "## \[${VERSION}\]" "$CHANGELOG" 2>/dev/null; then
    echo -e "${GREEN}✓${NC} CHANGELOG.md has entry for v${VERSION}"
else
    echo -e "${YELLOW}⚠${NC} CHANGELOG.md missing entry for v${VERSION}"
fi

CODE_VERSION=$(grep 'const Version' "$VERSION_FILE" | sed 's/.*"\(.*\)"/\1/')
if [ "$CODE_VERSION" = "$VERSION" ]; then
    echo -e "${GREEN}✓${NC} version.go matches release version"
else
    echo -e "${RED}✗${NC} version.go has v${CODE_VERSION}, expected v${VERSION}"
    echo "  Run: ./run.sh bump ${VERSION}"
    exit 1
fi

if git diff --quiet && git diff --staged --quiet; then
    echo -e "${GREEN}✓${NC} Working directory is clean"
else
    echo -e "${YELLOW}⚠${NC} Uncommitted changes in working directory"
fi

echo ""

if [ "$DRY_RUN" = true ]; then
    echo -e "${YELLOW}DRY RUN - No actions will be taken${NC}"
    echo ""
fi

if [ "$SKIP_BUILD" = false ]; then
    echo "📦 Step 1: Check dependencies"
    "$SCRIPT_DIR/buildDependencyChecker.sh" check
    echo ""
    
    echo "📦 Step 2: Building for all platforms"
    if [ "$DRY_RUN" = true ]; then
        echo "  Would run: ./scripts/build.sh all --clean"
    else
        "$SCRIPT_DIR/build.sh" all --clean
    fi
    echo ""
else
    echo -e "${YELLOW}Skipping build (--skip-build)${NC}"
    echo ""
fi

if [ "$SKIP_TAG" = false ]; then
    echo "🏷️  Step 3: Git tagging"
    
    if git rev-parse "v${VERSION}" &>/dev/null; then
        echo -e "${YELLOW}⚠${NC} Tag v${VERSION} already exists"
    else
        if [ "$DRY_RUN" = true ]; then
            echo "  Would run: git tag -a v${VERSION} -m \"Release v${VERSION}\""
            echo "  Would run: git push origin v${VERSION}"
        else
            echo "Creating tag v${VERSION}..."
            git tag -a "v${VERSION}" -m "Release v${VERSION}"
            echo -e "${GREEN}✓${NC} Created tag v${VERSION}"
            
            read -p "Push tag to origin? [Y/n] " -n 1 -r
            echo
            if [[ ! $REPLY =~ ^[Nn]$ ]]; then
                git push origin "v${VERSION}"
                echo -e "${GREEN}✓${NC} Pushed tag to origin"
            fi
        fi
    fi
    echo ""
else
    echo -e "${YELLOW}Skipping git tagging (--skip-tag)${NC}"
    echo ""
fi

echo "📤 Step 4: GitHub Release"

if command -v gh &> /dev/null; then
    RELEASE_FLAGS=""
    if [ "$DRAFT" = true ]; then
        RELEASE_FLAGS="--draft"
    fi
    
    RELEASE_NOTES=""
    if [ -f "$CHANGELOG" ]; then
        RELEASE_NOTES=$(awk "/## \[${VERSION}\]/,/## \[/" "$CHANGELOG" | head -n -1 | tail -n +2)
    fi
    
    if [ "$DRY_RUN" = true ]; then
        echo "  Would run: gh release create v${VERSION} ${DIST_DIR}/* --title \"v${VERSION}\" ${RELEASE_FLAGS}"
        if [ -n "$RELEASE_NOTES" ]; then
            echo "  With release notes from CHANGELOG.md"
        fi
    else
        echo "Creating GitHub release..."
        
        if [ -n "$RELEASE_NOTES" ]; then
            gh release create "v${VERSION}" "${DIST_DIR}"/* \
                --title "v${VERSION}" \
                --notes "$RELEASE_NOTES" \
                $RELEASE_FLAGS
        else
            gh release create "v${VERSION}" "${DIST_DIR}"/* \
                --title "v${VERSION}" \
                --generate-notes \
                $RELEASE_FLAGS
        fi
        
        echo -e "${GREEN}✓${NC} GitHub release created"
    fi
else
    echo -e "${YELLOW}⚠${NC} GitHub CLI (gh) not installed"
    echo ""
    echo "Manual release steps:"
    echo "  1. Go to: https://github.com/dale-tomson/declutter/releases/new"
    echo "  2. Select tag: v${VERSION}"
    echo "  3. Upload files from: ${DIST_DIR}/"
    echo ""
    echo "Or install GitHub CLI:"
    echo "  Linux: sudo apt install gh"
    echo "  macOS: brew install gh"
fi

echo ""
echo -e "${GREEN}🎉 Release process complete!${NC}"
echo ""
echo "Release files:"
ls -lh "$DIST_DIR" 2>/dev/null || echo "  (none)"
