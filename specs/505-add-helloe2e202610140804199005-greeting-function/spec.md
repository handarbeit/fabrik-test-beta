# Feature Specification: Add HelloE2E202610140804199005 greeting function

**Feature Branch**: `fabrik/issue-505`
**Created**: 2026-10-14
**Status**: Specified
**Input**: User description:

> ## Problem
> The alpha repo (`handarbeit/fabrik-test-alpha`) needs to call a new function from this repo's `pkg/greeting` package as part of an e2e regression test for cross-repo spawn. That function doesn't exist yet. This issue must land first, because alpha cannot import a function that does not exist.
>
> ## Requirements
> - In `pkg/greeting/greeting.go`, add an exported function `HelloE2E202610140804199005() string` that returns the literal string `"e2e-cross-repo-spawn-20261014-080419-9005"`.
> - Follow the existing pattern in the file: add a doc comment and return a literal string, and append the function after the existing `HelloE2E…` functions.
> - In `pkg/greeting/greeting_test.go`, add a test for it in the existing style: `const want = "e2e-cross-repo-spawn-20261014-080419-9005"` followed by `if got := HelloE2E202610140804199005(); got != want { t.Errorf("HelloE2E202610140804199005() = %q; want %q", got, want) }`.
> - Use the function name and string literal exactly as given. Do not invent a different name or value. A per-run-unique token is baked into them, and the e2e scenario's final assertion depends on this exact text landing in this repo.
>
> ## Scope
> Only `pkg/greeting/greeting.go` and `pkg/greeting/greeting_test.go`. No new dependencies and no other files. Do not modify or remove existing functions.
>
> ## Acceptance
> - `go build ./...` and `go test ./...` pass.
> - The function and its test are merged to the default branch, so alpha can bump its module pin to a commit that contains it.
>
> *Spawned by Fabrik from parent issue handarbeit/fabrik-test-alpha#8152 as a multi-issue decomposition. The parent's plan is at the link above.*

## Background

The alpha repository (`handarbeit/fabrik-test-alpha`) is running an e2e regression scenario for cross-repo spawn. That scenario needs to call a new sentinel greeting function exported from this repository's `pkg/greeting` package. The function does not exist yet, and alpha cannot import something that has not landed, so this issue must be merged first. Alpha will then bump its module pin to a commit containing the function. The function name and returned string embed a per-run-unique token, and the scenario's final assertion depends on that exact text appearing in this repo. The names and values below are therefore fixed and must not be altered.

This issue was spawned from parent issue handarbeit/fabrik-test-alpha#8152 as part of a multi-issue decomposition.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Expose the sentinel greeting function (Priority: P1)

A consumer of the `pkg/greeting` package (the alpha repo's e2e scenario) can call `HelloE2E202610140804199005()` and receive the agreed sentinel string, with a test in this repo guarding that behavior.

**Why this priority**: It is the only deliverable. Alpha's scenario is blocked until it is merged to the default branch.

**Independent Test**: Run `go build ./...` and `go test ./...` in this repo. The new test passes and nothing else changes.

**Acceptance Scenarios**:

1. **Given** the `pkg/greeting` package, **When** `HelloE2E202610140804199005()` is called, **Then** it returns exactly `"e2e-cross-repo-spawn-20261014-080419-9005"`.
2. **Given** the existing `HelloE2E…` functions and their tests, **When** the new function and test are added, **Then** the existing functions and tests are unchanged and still pass.
3. **Given** the change is merged to the default branch, **When** alpha bumps its module pin to a commit containing it, **Then** alpha can import and call the function.

---

### Edge Cases

- A mistyped function name or string literal (even one character) would silently break the e2e scenario's final assertion, so both must match the given text exactly.
- The new function must not collide with, shadow or alter any existing `HelloE2E…` function.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The `pkg/greeting` package MUST export a function named exactly `HelloE2E202610140804199005` that takes no arguments and returns a `string`.
- **FR-002**: `HelloE2E202610140804199005()` MUST return exactly the literal string `"e2e-cross-repo-spawn-20261014-080419-9005"`.
- **FR-003**: The function MUST follow the file's existing pattern: a doc comment, a literal-string return, and placement after the existing `HelloE2E…` functions.
- **FR-004**: The package's test file MUST include a test for the function in the existing style. It compares the result to `const want = "e2e-cross-repo-spawn-20261014-080419-9005"` and reports a mismatch with `t.Errorf("HelloE2E202610140804199005() = %q; want %q", got, want)`.
- **FR-005**: Existing functions and tests in the package MUST NOT be modified or removed.
- **FR-006**: The change MUST be limited to `pkg/greeting/greeting.go` and `pkg/greeting/greeting_test.go`, with no new dependencies and no other files.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `go build ./...` and `go test ./...` both pass on the change, with the new test among those run.
- **SC-002**: After merge, the default branch contains the function and test with the exact name and literal given, so alpha can bump its module pin to a commit that includes them.
- **SC-003**: The change touches exactly two files, and the diff contains only additions.

## Assumptions

- The existing `pkg/greeting` package already has `HelloE2E…` sentinel functions, which this change mirrors.
- The existing tests use plain `testing` with a `const want` comparison, as the issue describes.
- The function name and string literal are fixed by the e2e scenario and are not open to improvement.

## Out of Scope

- Any change to the alpha repo, including bumping its module pin and calling the new function.
- Modifying, renaming, deduplicating or removing existing `HelloE2E…` functions.
- New dependencies, new files, documentation changes or refactoring of the package.
- Any behavior beyond returning the fixed literal string.
