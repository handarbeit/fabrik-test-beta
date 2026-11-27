# Feature Specification: Add `HelloE2E202611270131363503` greeting function

**Feature Branch**: `fabrik/issue-804`
**Created**: 2026-11-26
**Status**: Specified
**Input**: User description: "# Add `HelloE2E202611270131363503` greeting function

## Problem

The cross-repo spawn e2e regression scenario (parent: handarbeit/fabrik-test-alpha#12109) needs an exported function in this repo's `pkg/greeting` package. Alpha will import it, so this change must land first.

## Requirements

- In `pkg/greeting/greeting.go`, add an exported function named exactly `HelloE2E202611270131363503`. It takes no arguments and returns a `string`.
- It MUST return exactly the literal `e2e-cross-repo-spawn-20261127-013136-3503`.
- Give it a doc comment in the same style as the existing `HelloE2E<token>` functions (\"returns the sentinel string for the cross-repo spawn regression test\").
- In `pkg/greeting/greeting_test.go`, add `TestHelloE2E202611270131363503`, following the existing pattern: `const want = \"e2e-cross-repo-spawn-20261127-013136-3503\"`; call `HelloE2E202611270131363503()`; `t.Errorf` on mismatch.
- Append the new function and test after the existing ones (the newest existing one is `HelloE2E202611262103375137`).

## Constraints

- Use the function name and string literal verbatim. They contain a per-run-unique token, and the scenario's final assertion depends on this exact text landing in this repo. Do not invent or alter either one.
- Do not change anything else in this repo.

## Acceptance

- `go test ./pkg/greeting/...` passes, including the new test.
- `pkg/greeting/greeting.go` contains `HelloE2E202611270131363503` returning `e2e-cross-repo-spawn-20261127-013136-3503`.
- The change is merged to `main` so the parent's alpha repo can depend on that commit.

## References

Mentioned for context: handarbeit/fabrik-test-alpha#12109"

## Background

A cross-repo spawn end-to-end regression scenario, driven by parent issue handarbeit/fabrik-test-alpha#12109, needs a new exported sentinel function in this repo's `greeting` package. The alpha repo will import it, so this change must be merged to `main` first. The scenario's final assertion depends on the exact function name and string value (both contain a per-run-unique token) appearing in this repo, so neither may be invented or altered. The package already holds a series of sentinel functions of the same shape, each returning a per-run string for earlier runs of this scenario; this adds the newest one.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Provide the sentinel function for the cross-repo scenario (Priority: P1)

The cross-repo spawn regression scenario needs a new exported sentinel function in the greeting package so that the alpha repo can depend on it once this change is merged.

**Why this priority**: It is the only deliverable of this issue, and the parent scenario cannot complete without it.

**Independent Test**: Run the greeting package's tests; the new test passes and confirms the function returns the exact expected string.

**Acceptance Scenarios**:

1. **Given** the greeting package, **When** `HelloE2E202611270131363503()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261127-013136-3503`.
2. **Given** the new test `TestHelloE2E202611270131363503`, **When** the greeting package's tests run, **Then** the test passes.
3. **Given** the change is merged, **When** the alpha repo depends on the resulting `main` commit, **Then** it can import `HelloE2E202611270131363503` from the greeting package.

---

### Edge Cases

- A typo in the function name or string literal would pass the repo's own test (which restates the literal) yet fail the parent scenario's final assertion, so both must match the issue text character for character.
- The new function and test must be appended after the existing newest one (`HelloE2E202611262103375137`) so existing functions and tests are unaffected.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The greeting package exports a function named exactly `HelloE2E202611270131363503` that takes no arguments and returns a `string`.
- **FR-002**: `HelloE2E202611270131363503` returns exactly the literal `e2e-cross-repo-spawn-20261127-013136-3503`.
- **FR-003**: The function has a doc comment in the same style as the existing sentinel functions, stating that it returns the sentinel string for the cross-repo spawn regression test.
- **FR-004**: The greeting package's test file contains `TestHelloE2E202611270131363503`, following the existing pattern: it declares the expected value as the constant `e2e-cross-repo-spawn-20261127-013136-3503`, calls `HelloE2E202611270131363503()`, and reports an error on mismatch.
- **FR-005**: The new function and its test are placed after the existing ones in their files (the newest existing one is `HelloE2E202611262103375137`).
- **FR-006**: No other changes are made to this repository.

### Key Entities *(if applicable)*

- **Sentinel function**: An exported, no-argument greeting-package function returning a fixed per-run-unique string, used by the cross-repo e2e scenario to prove that a change landed in this repo.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The greeting package's tests pass, including the new test.
- **SC-002**: The greeting package source contains `HelloE2E202611270131363503` returning `e2e-cross-repo-spawn-20261127-013136-3503`, character for character.
- **SC-003**: The change is merged to `main`, so the parent scenario's alpha repo can depend on that commit.
- **SC-004**: The diff touches only the greeting package's source and test files, plus the persisted spec file for this issue.

## Assumptions

- The existing sentinel functions and tests follow one uniform pattern (doc comment, no-argument function returning a literal; constant `want`, call, `t.Errorf` on mismatch), which the new ones mirror.
- The function name and string literal given in the issue are authoritative and already correct; they are used verbatim.

## Out of Scope *(optional)*

- Any change to the alpha repo or to the parent scenario itself.
- Any change to existing sentinel functions, tests, or other files in this repo.
- Refactoring or consolidating the series of sentinel functions.

## Source References *(optional)*

- Parent issue (mentioned for context): handarbeit/fabrik-test-alpha#12109
