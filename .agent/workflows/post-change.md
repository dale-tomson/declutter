---
description: Checklist for validating changes, updating versions, changelogs, and documentation.
---
# Post-Change Workflow

Follow this checklist key changes are made to the codebase. This ensures the project history and documentation remain clean and consistent.

## 1. Source & Version Check

- [ ] **Check Scopes**: Did changes occur in `cmd/`, `internal/`, or `pkg/`?
- [ ] **Verify Version**: If source code changed, check `internal/version/version.go`.
    -   Has the version been bumped?
    -   Does the bump match the change type?
        -   **Patch**: Bug fixes, refactoring, non-breaking fixes.
        -   **Minor**: New features, backwards compatible.
        -   **Major**: Breaking API changes.
- [ ] **Action**: If mismatch, bump the version in `internal/version/version.go`.

## 2. Changelog Management

- [ ] **Individual File**: Ensure a file exists at `docs/changelogs/v<VERSION>.md` containing the specific notes for this version.
- [ ] **Main Changelog**: Update `CHANGELOG.md` in the project root.
    -   Add the new version entry at the top.
    -   **CRITICAL**: Keep **ONLY the last 5 versions** in `CHANGELOG.md`.
    -   Remove older versions (ensure they are preserved in `docs/changelogs/` first).

## 3. Documentation Sync

- [ ] **README Updates**:
    -   Update the "Latest Changes" section in `README.md` with a brief summary of this update.
    -   If new flags/features were added, ensure the Usage section is updated.
- [ ] **Verify Links**: Check `README.md` and `docs/` for broken relative links (e.g., ensure `changelogs/` points to `docs/changelogs/`).

## 4. Final Validation

- [ ] **Build**: Run `make build` to ensure the version string is compiling correctly.
- [ ] **Test**: Run `make test` to verify stability.

## 5. Submission

- [ ] **Checkout Main**: Switch to `main` branch (`git checkout main`).
- [ ] **Pull Latest**: Pull latest changes (`git pull origin main`).
- [ ] **Branch**: Checkout existing branch or create new one (`git checkout -b <branch>` or `git checkout <branch>`) following `CONTRIBUTING.md`.
    - [ ] `git rebase main` if needed.
    - [ ] Resolve any conflicts.
    - [ ] `git checkout --ours/theirs <file>` if needed.
- [ ] **Stage**: Stage changes (`git add .`).
- [ ] **Commit**: Commit with conventional message (`git commit -m "<type>: <message>"`).
- [ ] **Push**: Push to origin (`git push origin <branch>`).
