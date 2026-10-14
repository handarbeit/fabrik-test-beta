# Feature Specification: Add HelloE2E202610140732181422 greeting sentinel (cross-repo spawn e2e)

**Feature Branch**: `fabrik/issue-503`
**Created**: 2026-10-14
**Status**: Specified
**Input**: User description: "## Problem
The sibling repo `handarbeit/fabrik-test-alpha` (issue #8135, cross-repo spawn e2e) needs to call a new exported greeting function that does not exist in this repo yet. This repo must land the function first, because alpha cannot import a function that does not exist.

## Requirements
- Add an exported function `HelloE2E202610140732181422() string` to `pkg/greeting/greeting.go`. It returns exactly the literal string `\"e2e-cross-repo-spawn-20261014-073218-1422\"`.
- Add a test for it in `pkg/greeting/greeting_test.go`. The test asserts that `HelloE2E202610140732181422()` returns exactly `\"e2e-cross-repo-spawn-20261014-073218-1422\"`.
- Follow the pattern of the existing `HelloE2E…()` functions and tests in the same files.
- Use the function name and the string literal verbatim, character for character. Do not invent your own. A per-run-unique token is already baked into both, and the scenario's final assertion depends on that exact text landing in this repo.

## Scope
Only `pkg/greeting/greeting.go` and `pkg/greeting/greeting_test.go` in this repo. No changes to any other file are needed.

## Acceptance
- `go build ./...` and `go test ./...` pass.
- The change is merged to this repo's default branch, so that alpha can pin the resulting commit with `go get`."

## Background

This issue was spawned from parent issue `handarbeit/fabrik-test-alpha#8135`, a cross-repo spawn end-to-end scenario. The sibling repo `handarbeit/fabrik-test-alpha` needs to call a new exported greeting function that does not yet exist in this repo, and cannot import a function that does not exist. This repo therefore has to land the function first, so that alpha can later pin the resulting commit with `go get`.

The greeting package already contains a series of sentinel functions of the same kind (`HelloE2E` and several timestamped `HelloE2E<token>` variants), each added for an earlier cross-repo spawn run. This change adds one more in the same style. The function name and its returned string carry a per-run-unique token, and the scenario's final assertion depends on that exact text landing in this repo, so both must be used exactly as given.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Sentinel function available for the sibling repo to import (Priority: P1)

A maintainer running the cross-repo spawn scenario needs the greeting package in this repo to export a sentinel function with a fixed name that returns a fixed string, so that the sibling repo can import it and the scenario's final assertion can find that exact text in this repo.

**Why this priority**: It is the only deliverable of the issue. The sibling repo's work is blocked until the function exists on this repo's default branch.

**Independent Test**: Run the greeting package's tests and confirm the new test passes. Calling the new function returns exactly `e2e-cross-repo-spawn-20261014-073218-1422`.

**Acceptance Scenarios**:

1. **Given** the greeting package, **When** `HelloE2E202610140732181422()` is called, **Then** it returns exactly the string `e2e-cross-repo-spawn-20261014-073218-1422`.
2. **Given** the change is merged to the default branch, **When** the sibling repo resolves this module at the resulting commit with `go get`, **Then** the function can be imported and called.
3. **Given** the change, **When** the repository's build and test commands are run, **Then** both pass.

---

### Edge Cases

- The function name or the returned string differs from the given text by even one character: the scenario's final assertion fails. Both must match character for character.
- The new function must not alter the behaviour of any existing `HelloE2E…` function or its tests.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The greeting package exports a function named exactly `HelloE2E202610140732181422` that takes no arguments and returns a string.
- **FR-002**: `HelloE2E202610140732181422()` returns exactly the literal string `e2e-cross-repo-spawn-20261014-073218-1422`.
- **FR-003**: The greeting package's tests include a test that asserts `HelloE2E202610140732181422()` returns exactly `e2e-cross-repo-spawn-20261014-073218-1422`.
- **FR-004**: The new function and its test follow the conventions of the existing `HelloE2E…` functions and tests in the same package (naming, doc comment style, test style).
- **FR-005**: The function name and the string literal are used verbatim as given, with no alteration or substitution.
- **FR-006**: Only the greeting package's source file and its test file are changed; no other file in the repository is modified.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The repository's build and full test run both complete successfully with the change applied.
- **SC-002**: The new test passes and fails if the returned string is changed in any way.
- **SC-003**: After merge to the default branch, the sibling repo can resolve this module at the resulting commit and import the new function.
- **SC-004**: The diff touches exactly two files: the greeting source file and the greeting test file.

## Assumptions

- The existing `HelloE2E…` functions and tests in the greeting package define the convention to follow; the new function is a plain addition with no shared helpers or refactoring.
- The doc comment may reference this issue and the sibling repo's issue `#8135`, in the style of the existing sentinels, though this is not required.
- Merging to the default branch follows the repository's normal pipeline; pinning the commit in the sibling repo is that repo's concern.

## Out of Scope

- Any change to files other than the greeting package's source and test files.
- Changes to the sibling repo `handarbeit/fabrik-test-alpha`, including pinning the resulting commit with `go get`.
- Refactoring or removing existing `HelloE2E…` sentinels.

## Source References

- Parent issue: `handarbeit/fabrik-test-alpha#8135`
