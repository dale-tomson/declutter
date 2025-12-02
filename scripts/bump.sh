#!/bin/bash

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"
VERSION_FILE="$PROJECT_ROOT/internal/version/version.go"

CURRENT_VERSION=$(grep 'const Version' "$VERSION_FILE" | sed 's/.*"\(.*\)"/\1/')

show_help() {
    echo "Declutter Version Bump Tool"
    echo ""
    echo "Usage: $0 [major|minor|patch|<version>]"
    echo ""
    echo "Commands:"
    echo "  major      Bump major version (1.0.0 -> 2.0.0)"
    echo "  minor      Bump minor version (1.0.0 -> 1.1.0)"
    echo "  patch      Bump patch version (1.0.0 -> 1.0.1)"
    echo "  <version>  Set specific version (e.g., 2.0.0-beta.1)"
    echo ""
    echo "Current version: ${CURRENT_VERSION}"
}

bump_version() {
    local type="$1"
    local major minor patch
    
    IFS='.' read -r major minor patch <<< "$CURRENT_VERSION"
    patch="${patch%%-*}"
    
    case "$type" in
        major)
            major=$((major + 1))
            minor=0
            patch=0
            ;;
        minor)
            minor=$((minor + 1))
            patch=0
            ;;
        patch)
            patch=$((patch + 1))
            ;;
    esac
    
    echo "${major}.${minor}.${patch}"
}

if [ -z "$1" ]; then
    show_help
    exit 0
fi

case "$1" in
    -h|--help|help)
        show_help
        exit 0
        ;;
    major|minor|patch)
        NEW_VERSION=$(bump_version "$1")
        ;;
    *)
        NEW_VERSION="$1"
        ;;
esac

echo "📦 Bumping version: ${CURRENT_VERSION} → ${NEW_VERSION}"

sed -i "s/const Version = \".*\"/const Version = \"${NEW_VERSION}\"/" "$VERSION_FILE"

echo "✅ Updated $VERSION_FILE"

CHANGELOG="$PROJECT_ROOT/CHANGELOG.md"
if [ -f "$CHANGELOG" ]; then
    DATE=$(date +%Y-%m-%d)
    if ! grep -q "## \[${NEW_VERSION}\]" "$CHANGELOG"; then
        sed -i "/## \[Unreleased\]/a\\
\\
## [${NEW_VERSION}] - ${DATE}" "$CHANGELOG"
        echo "✅ Added version entry to CHANGELOG.md"
    fi
fi

echo ""
echo "🎉 Version bumped to ${NEW_VERSION}"
echo ""
echo "Next steps:"
echo "  1. Update CHANGELOG.md with release notes"
echo "  2. git add -A && git commit -m \"Bump version to ${NEW_VERSION}\""
echo "  3. ./run.sh release"
