# Contributing to Declutter

Thank you for your interest in contributing to Declutter! To keep our codebase clean and history readable, please follow these guidelines.

## 1. Branch Naming Convention

We use the format: `type/short-description`

*   **Prefixes**:
    *   `feat/`: A new feature
    *   `fix/`: A bug fix
    *   `docs/`: Documentation only changes
    *   `style/`: Formatting, missing semi colons, etc; no code change
    *   `refactor/`: Refactoring production code
    *   `test/`: Adding missing tests, refactoring tests
    *   `chore/`: Updating build tasks, package manager configs, etc

*   **Examples**:
    *   `feat/add-dark-mode`
    *   `fix/windows-path-issue`
    *   `docs/update-readme`

## 2. Commit Naming Convention

We follow the [Conventional Commits](https://www.conventionalcommits.org/) specification.

**Format**: `type(scope): description`

*   **Type**: Must be one of the prefixes listed above (feat, fix, docs, etc.).
*   **Scope**: (Optional) The specific component affected (e.g., `ui`, `organizer`, `ci`).
*   **Description**: Short, imperative summary (use "add" not "added").

**Examples**:
*   `feat(ui): add progress bar component`
*   `fix(organizer): resolve null pointer on empty dir`
*   `chore: update go dependencies`

## 3. Pull Request (PR) Convention

*   **Title**: Follow the Commit Naming Convention (e.g., `feat: Add dark mode`).
*   **Description**:
    *   **Summary**: What does this PR do?
    *   **Type of Change**: (New feature, Bug fix, Refactor, etc.)
    *   **Checklist**:
        *   [ ] Code builds without errors (`make build`)
        *   [ ] Tests pass (`make test`)
        *   [ ] Documentation updated (if applicable)

## 4. Workflow

1.  Create a branch from `main`.
2.  Make your changes.
3.  Commit using the naming convention.
4.  Push to origin.
5.  Open a PR against `main`.
