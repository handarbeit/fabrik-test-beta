# Feature Specification: Add HelloE2E202610141912195737 greeting function

**Feature Branch**: `fabrik/issue-507`
**Created**: 2026-10-14
**Status**: Specified
**Input**: User description:

> ## Goal
>
> Add a new exported greeting function that `handarbeit/fabrik-test-alpha` will import. Beta must land first because alpha cannot import a function that does not yet exist. This is the beta half of a cross-repo e2e regression scenario for on-demand spawn-target repo init (#803) and the per-run-unique fixture token (#1308).
>
> ## What needs to change
>
> In `pkg/greeting/greeting.go`:
> - Add an exported function `HelloE2E202610141912195737() string` that returns the literal string `"e2e-cross-repo-spawn-20261014-191219-5737"`.
> - Append it after the existing `HelloE2E…` functions, in the same style as its predecessors.
>
> In `pkg/greeting/greeting_test.go`:
> - Add a test for `HelloE2E202610141912195737` asserting that it returns exactly `"e2e-cross-repo-spawn-20261014-191219-5737"`.
> - Follow the style of the existing `HelloE2E…` tests.
>
> ## Constraints
>
> - Use the function name `HelloE2E202610141912195737` and the string literal `e2e-cross-repo-spawn-20261014-191219-5737` verbatim. Do not rename, regenerate or "improve" them. A per-run-unique token is baked into both, and this scenario's final assertion depends on that exact text landing in this repo.
> - Do not change any existing function or test.
> - The PR must merge to this repo's default branch before the alpha dependent issue can proceed.
> - The PR body must include `Closes #<this issue>`.
>
> ## Acceptance
>
> - `go build ./...` and `go test ./...` pass.
> - `HelloE2E202610141912195737()` returns `"e2e-cross-repo-spawn-20261014-191219-5737"`.
> - The change is merged to the default branch.
>
> *Spawned by Fabrik from parent issue handarbeit/fabrik-test-alpha#8162 as a multi-issue decomposition. The parent's plan is at the link above.*

## Background

This is the beta half of a cross-repo e2e regression scenario. It exercises on-demand initialisation of a spawn-target repository (#803) and the per-run-unique fixture token (#1308). The parent issue, `handarbeit/fabrik-test-alpha#8162`, spawned this issue as part of a multi-issue decomposition.

The alpha repository needs to import a new greeting function from this repo. That function does not exist yet, so this change has to merge first. The scenario's final assertion checks that the exact token text lands in this repo, so the identifiers and string literal must be preserved exactly.

The greeting package already holds a series of earlier sentinel functions (`HelloE2E…`), each paired with a test. This change adds one more in the same pattern.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Provide the new sentinel greeting function (Priority: P1)

A consumer, namely the alpha repository's dependent issue, can import `HelloE2E202610141912195737` from the greeting package and receive the run-unique sentinel string.

**Why this priority**: It is the only deliverable of this issue, and the alpha dependent issue is blocked until it lands.

**Independent Test**: Call `HelloE2E202610141912195737()` and compare the result with `"e2e-cross-repo-spawn-20261014-191219-5737"`. A dedicated unit test does this.

**Acceptance Scenarios**:

1. **Given** the greeting package, **When** `HelloE2E202610141912195737()` is called, **Then** it returns exactly `"e2e-cross-repo-spawn-20261014-191219-5737"`.
2. **Given** the change is applied, **When** `go build ./...` and `go test ./...` run, **Then** both succeed.
3. **Given** the change is applied, **When** the existing `HelloE2E…` functions and tests are inspected, **Then** they are unchanged.

---

### Edge Cases

- The function name or string literal is altered, for example by regeneration or by "fixing" a perceived typo. The scenario's final assertion would then fail, so both must be used verbatim.
- The new function is placed before existing ones, or an existing one is modified. The new function must be appended after the existing `HelloE2E…` functions, and nothing existing may change.
- The PR body omits `Closes #507`. Fabrik's PR discovery would then not link the PR to this issue.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The greeting package exports a function `HelloE2E202610141912195737() string` that returns exactly `"e2e-cross-repo-spawn-20261014-191219-5737"`.
- **FR-002**: The new function is appended after the existing `HelloE2E…` functions and follows their style, including a doc comment.
- **FR-003**: The greeting package's tests include a test for `HelloE2E202610141912195737` that asserts the exact return value, following the style of the existing `HelloE2E…` tests and appended after them.
- **FR-004**: No existing function or test in the greeting package is modified or removed.
- **FR-005**: The function name and string literal are used verbatim, with no renaming, regeneration or alteration.
- **FR-006**: The PR body includes `Closes #507` and merges to the repo's default branch.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `go build ./...` and `go test ./...` both pass on the change.
- **SC-002**: `HelloE2E202610141912195737()` returns exactly `"e2e-cross-repo-spawn-20261014-191219-5737"`, as confirmed by a passing test.
- **SC-003**: The diff touches only the two greeting files and adds only the new function and its test, with no changes to existing lines.
- **SC-004**: The change is merged to the default branch, so the alpha dependent issue can proceed.

## Assumptions

- The existing `pkg/greeting` package and its `HelloE2E…` functions and tests are the style reference, and the newest one is the model to follow.
- The token in the name and literal is intentionally per-run-unique, so nothing is deduplicated against earlier tokens.
- No documentation changes are needed, since this is a test-fixture sentinel function.

## Out of Scope *(optional)*

- Any change to the alpha repository or its dependent issue.
- Refactoring, consolidating or removing the existing `HelloE2E…` functions.
- Any change to engine behaviour related to #803 or #1308.

## Source References *(optional)*

- Parent issue: `handarbeit/fabrik-test-alpha#8162`
- Related engine work: #803 (on-demand spawn-target repo init), #1308 (per-run-unique fixture token)
