# Feature Specification: Add HelloE2E202610260610197278 to pkg/greeting

**Feature Branch**: `fabrik/issue-617`
**Created**: 2026-10-26
**Status**: Specified
**Input**: User description: "## Goal
Add one exported greeting function, and a test for it, to the `pkg/greeting` package. The alpha repo (`handarbeit/fabrik-test-alpha`) will consume it after this lands. This is a test fixture for an end-to-end cross-repo spawn regression (#803, #1308).

## Requirements
- In `pkg/greeting/greeting.go`, add the exported function `HelloE2E202610260610197278() string`. It must return exactly the literal string `e2e-cross-repo-spawn-20261026-061019-7278`. Follow the shape of the existing `HelloE2E…` functions in that file.
- In `pkg/greeting/greeting_test.go`, add a test that calls `HelloE2E202610260610197278()` and asserts the result equals `e2e-cross-repo-spawn-20261026-061019-7278`. The test must pass.

## Constraints
- Use the function name and the string literal exactly as written above. Do not change the suffix or regenerate the token. A final assertion depends on this exact text landing in this repo.
- Do not change any other behavior in this repo. The only changes are the one new function and its test.
- No `go.mod` changes should be needed.

## Acceptance
1. `HelloE2E202610260610197278()` returns exactly `e2e-cross-repo-spawn-20261026-061019-7278`.
2. `go test ./pkg/greeting/...` passes and includes the new test.
3. The change is merged to the default branch, so the alpha repo can pin to the merged commit."

## Background

This issue is a test fixture for an end-to-end regression covering cross-repo issue spawning (#803, #1308). It was spawned by Fabrik from parent issue `handarbeit/fabrik-test-alpha#9739` as part of a multi-issue decomposition. The alpha repo (`handarbeit/fabrik-test-alpha`) depends on this repo exposing a specific sentinel greeting function. After this change merges to the default branch, the alpha repo pins to the merged commit and a final assertion checks that the exact sentinel text landed here.

The `pkg/greeting` package already contains many sentinel functions of the same `HelloE2E<token>` shape. Each returns a fixed literal string for a prior run of this regression. This issue adds one more, with a new token, and does not touch the existing ones.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Expose the new sentinel function (Priority: P1)

A consumer in the alpha repo needs to call `HelloE2E202610260610197278()` from `pkg/greeting` and get back the exact sentinel string for this run.

**Why this priority**: This is the whole change. The cross-repo regression cannot complete without it.

**Independent Test**: Call `HelloE2E202610260610197278()` and compare the result with `e2e-cross-repo-spawn-20261026-061019-7278`.

**Acceptance Scenarios**:

1. **Given** the `pkg/greeting` package, **When** `HelloE2E202610260610197278()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261026-061019-7278`.
2. **Given** the existing `HelloE2E…` functions, **When** the new function is added, **Then** each existing function returns the same value as before.

---

### User Story 2 - Cover the function with a passing test (Priority: P1)

A maintainer needs an automated test that pins the new function's return value, so a later change cannot silently alter the sentinel.

**Why this priority**: The acceptance criteria require a test that passes in the package's test run.

**Independent Test**: Run the package tests and confirm the new test appears and passes.

**Acceptance Scenarios**:

1. **Given** the new test in `pkg/greeting/greeting_test.go`, **When** `go test ./pkg/greeting/...` is run, **Then** the test runs and passes.
2. **Given** the function returned a different string, **When** the test runs, **Then** it fails.

---

### Edge Cases

- The function name and string literal must match the issue text character for character. A near-miss in either one breaks the downstream assertion in the alpha repo.
- The token is fixed by the issue and must not be regenerated or altered, even though it follows the pattern of earlier tokens.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The `pkg/greeting` package MUST export a function named `HelloE2E202610260610197278` that takes no arguments and returns a `string`.
- **FR-002**: `HelloE2E202610260610197278()` MUST return exactly `e2e-cross-repo-spawn-20261026-061019-7278`.
- **FR-003**: The new function MUST follow the shape of the existing `HelloE2E…` functions in the package: a doc comment describing it as a cross-repo spawn regression sentinel, and a body that returns the literal.
- **FR-004**: The package's test file MUST contain a test that calls `HelloE2E202610260610197278()` and asserts the result equals `e2e-cross-repo-spawn-20261026-061019-7278`.
- **FR-005**: The new test MUST pass when `go test ./pkg/greeting/...` is run.
- **FR-006**: The change MUST NOT alter any existing function, test, or other repository behavior, and MUST NOT require `go.mod` changes.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: A call to `HelloE2E202610260610197278()` returns a string equal to `e2e-cross-repo-spawn-20261026-061019-7278`.
- **SC-002**: `go test ./pkg/greeting/...` passes, and its output includes the new test.
- **SC-003**: The diff for this issue adds only the one function and its test, and changes no existing function, test, or `go.mod`.
- **SC-004**: The change is merged to the default branch (`main`), so the alpha repo can pin to the merged commit.

## Assumptions

- The existing `HelloE2E…` functions in `pkg/greeting/greeting.go` are the reference for shape and doc-comment wording. The newest ones are one doc comment plus a one-line body returning a literal.
- The existing test file already holds a test per sentinel function. The new test follows the style of those tests.
- No function named `HelloE2E202610260610197278` exists in the package yet.
- The issue carries `fabrik:yolo`, so the pipeline may advance without a separate human approval of this spec.

## Out of Scope

- Any change to the alpha repo (`handarbeit/fabrik-test-alpha`), including consuming the function.
- Changes to, removal of, or consolidation of the existing `HelloE2E…` functions and tests.
- Changes to `go.mod` or any other repository configuration.
- Any behavior change in this repo beyond the one function and its test.

## Source References

- Parent issue: `handarbeit/fabrik-test-alpha#9739`
- Related regression issues: #803, #1308
