# Feature Specification: Add HelloE2E202610250554429282 to pkg/greeting (e2e cross-repo spawn)

**Feature Branch**: `fabrik/issue-599`
**Created**: 2026-10-25
**Status**: Specified
**Input**: User description: "Add one exported greeting function and its test to this repo's `pkg/greeting` package. This is a trivial, self-contained fixture change for an end-to-end cross-repo spawn regression test (#803 on-demand spawn-target repo init, #1308 per-run-unique fixture token). The same change should flow straight through the pipeline without needing design decisions. Another repository (`handarbeit/fabrik-test-alpha`) is blocked on this issue and will import the function once it is merged, so it must land promptly on the default branch.

Requirements:
- In `pkg/greeting/greeting.go`, add an exported function `HelloE2E202610250554429282() string` that returns exactly the literal string `\"e2e-cross-repo-spawn-20261025-055442-9282\"`. Follow the existing shape of the other `HelloE2E...` functions in the file: a doc comment above the function, and a body that returns the literal.
- In `pkg/greeting/greeting_test.go`, add `TestHelloE2E202610250554429282`. Follow the existing pattern: `const want = \"e2e-cross-repo-spawn-20261025-055442-9282\"` and `if got := HelloE2E202610250554429282(); got != want { t.Errorf(...) }`.

Constraints:
- Use the function name `HelloE2E202610250554429282` and the string `e2e-cross-repo-spawn-20261025-055442-9282` exactly as written above. Do not invent a different name or literal. A per-run-unique token is baked into both, and the scenario's final assertion checks for this exact text on the default branch.
- Work from a fresh fetch of the default branch, because local clones of this repo may be stale.
- Make no other changes to the greeting package or elsewhere.

Acceptance:
- `go test ./pkg/greeting/...` passes, including the new test.
- The default branch contains both `HelloE2E202610250554429282` and `e2e-cross-repo-spawn-20261025-055442-9282`.

*Spawned by Fabrik from parent issue handarbeit/fabrik-test-alpha#9329 as a multi-issue decomposition. The parent's plan is at the link above.*"

## Background

This issue is a fixture for an end-to-end regression test of cross-repo child spawning (#803, on-demand initialisation of a spawn-target repo; #1308, per-run-unique fixture tokens). A parent issue in `handarbeit/fabrik-test-alpha` (#9329) decomposed its work and spawned this child here. The parent's repository is blocked on this issue and will import the new function once it is merged, so the change must reach the default branch (`main`) promptly. The `pkg/greeting` package already holds many functions of the same kind, each returning a per-run-unique sentinel string; this adds one more for the current run. The scenario's final assertion looks for the exact function name and string on the default branch.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Provide the per-run sentinel greeting (Priority: P1)

A consumer (the blocked repository `handarbeit/fabrik-test-alpha`, and the e2e scenario's final assertion) needs an exported function in `pkg/greeting` that returns this run's unique sentinel string, together with a test that pins its value.

**Why this priority**: It is the only deliverable; the parent repository cannot proceed until it is on the default branch.

**Independent Test**: Run the package's tests; the new test passes and the default branch contains both the function name and the literal.

**Acceptance Scenarios**:

1. **Given** the `pkg/greeting` package on the default branch after merge, **When** a caller invokes `HelloE2E202610250554429282()`, **Then** it returns exactly `e2e-cross-repo-spawn-20261025-055442-9282`.
2. **Given** the updated package, **When** the package tests are run, **Then** `TestHelloE2E202610250554429282` exists and passes along with all existing tests.
3. **Given** the merged change, **When** the default branch is searched for the function name and the literal, **Then** both are found.

---

### Edge Cases

- The local clone may be stale: the change must be based on a fresh fetch of the default branch so it does not conflict with or omit earlier sentinel functions.
- The name and literal embed a per-run token; any deviation (typo, different timestamp, different suffix) would fail the scenario's final assertion.
- Existing `HelloE2E...` functions and tests must remain untouched.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `pkg/greeting` MUST export a function named exactly `HelloE2E202610250554429282` taking no arguments and returning a `string`.
- **FR-002**: `HelloE2E202610250554429282()` MUST return exactly the string `e2e-cross-repo-spawn-20261025-055442-9282`.
- **FR-003**: The function MUST follow the shape of the existing `HelloE2E...` functions: a doc comment above it and a body that returns the literal.
- **FR-004**: A test named `TestHelloE2E202610250554429282` MUST be added to the package's test file, following the existing pattern (a `want` constant equal to the literal and a failure report when the function's result differs).
- **FR-005**: The change MUST be based on a fresh fetch of the default branch.
- **FR-006**: No other changes MAY be made to the greeting package or elsewhere in the repository.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The package's test suite passes, including the new test.
- **SC-002**: After merge, the default branch contains both `HelloE2E202610250554429282` and `e2e-cross-repo-spawn-20261025-055442-9282`.
- **SC-003**: The diff touches only the greeting source file and its test file, and only adds the new function and test.

## Assumptions

- The default branch is `main`, as stated in the issue context.
- The existing `HelloE2E...` functions and tests in `pkg/greeting` define the conventions to follow (doc comment wording, constant-and-compare test style).
- The new function is appended after the existing ones; ordering is not significant.
- The issue's requirements are fully prescriptive, so no design decisions or clarifications are needed.

## Out of Scope

- Any change to other greeting functions or to anything outside `pkg/greeting`.
- Changes in `handarbeit/fabrik-test-alpha`, including importing the new function (handled by the blocked parent issue).
- Refactoring or consolidating the existing sentinel functions.

## Source References

- Parent issue: `handarbeit/fabrik-test-alpha#9329` (cross-repo spawn decomposition)
- Existing sentinel functions in `pkg/greeting/greeting.go` and their tests in `pkg/greeting/greeting_test.go`
