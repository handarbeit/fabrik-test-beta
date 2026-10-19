# Feature Specification: add run-unique greeting function

**Feature Branch**: `fabrik/issue-559`
**Created**: 2026-10-19
**Status**: Specified
**Input**: User description: "This issue is the beta-repo half of a cross-repo e2e scenario. The consuming change lives in `handarbeit/fabrik-test-alpha` (issue #8795), which calls this function from its `main.go`. Alpha cannot compile until this change has landed on beta's default branch. Add an exported function `HelloE2E202610190605566209() string` to `pkg/greeting/greeting.go`. It returns exactly the literal `e2e-cross-repo-spawn-20261019-060556-6209`. Add a test for it in `pkg/greeting/greeting_test.go` asserting that exact return value. Follow the style of the existing `HelloE2E…` tests in that file."

## Background

This issue is the beta-repo half of a cross-repo end-to-end scenario. The consuming change lives in `handarbeit/fabrik-test-alpha` (issue #8795), whose `main.go` calls this function. Alpha cannot compile until this change has landed on beta's default branch, so alpha must be able to pin to a commit containing it. The function name and string literal are unique to this run, and the scenario's final assertion looks for this exact text in this repo.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Provide the run-unique greeting function (Priority: P1)

A consumer in the alpha repo needs to call `HelloE2E202610190605566209()` from beta's greeting package and receive the run-unique token string, so alpha compiles and the cross-repo scenario can complete.

**Why this priority**: It is the only deliverable of this issue; alpha is blocked without it.

**Independent Test**: Run the greeting package tests; the new test asserts the exact return value.

**Acceptance Scenarios**:

1. **Given** the greeting package on the default branch, **When** `HelloE2E202610190605566209()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261019-060556-6209`.
2. **Given** the change is merged, **When** the greeting package tests, build and vet are run, **Then** all succeed.

---

### Edge Cases

- The function name or literal is altered, even slightly (for example a different timestamp digit): the scenario's final assertion fails to find the token, so both must be used verbatim.
- Existing greeting functions or tests are modified as a side effect: out of bounds, see Constraints.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The greeting package MUST export a function `HelloE2E202610190605566209() string` in `pkg/greeting/greeting.go` that returns exactly the literal `e2e-cross-repo-spawn-20261019-060556-6209`.
- **FR-002**: `pkg/greeting/greeting_test.go` MUST contain a test asserting that exact return value, in the style of the existing `HelloE2E…` tests in that file.
- **FR-003**: The new function MUST follow the shape of the existing `HelloE2E<timestamp>` functions in `pkg/greeting/greeting.go`.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `HelloE2E202610190605566209()` returns exactly `e2e-cross-repo-spawn-20261019-060556-6209`.
- **SC-002**: `go test ./pkg/greeting/...` passes, including the new test.
- **SC-003**: `go build ./...` and `go vet ./...` succeed.
- **SC-004**: The change is merged to the default branch so that alpha can pin to a commit containing it.

## Assumptions

- The existing `HelloE2E<timestamp>` functions and tests in the greeting package provide a clear shape and style to follow.
- No documentation or module metadata updates are needed beyond the new function and test.

## Out of Scope *(optional)*

- Changes to other greeting functions, other packages, or module metadata.
- The consuming change in `handarbeit/fabrik-test-alpha` (tracked in alpha issue #8795).

## Constraints

- Use the function name `HelloE2E202610190605566209` and the string literal `e2e-cross-repo-spawn-20261019-060556-6209` verbatim. The token is unique to this run, and the scenario's final assertion looks for this exact text in this repo. Do not invent or substitute another name or value.
- Follow the shape of the existing `HelloE2E<timestamp>` functions in `pkg/greeting/greeting.go`.
- Make no other changes to this repo. Do not touch other greeting functions, other packages or module metadata beyond what the new function and test need.

## Acceptance

1. `HelloE2E202610190605566209()` returns exactly `e2e-cross-repo-spawn-20261019-060556-6209`.
2. `go test ./pkg/greeting/...` passes, including the new test.
3. `go build ./...` and `go vet ./...` succeed.
4. The change is merged to the default branch so that alpha can pin to a commit containing it.

## Source References

- Parent issue: handarbeit/fabrik-test-alpha#8795 (multi-issue decomposition; this issue was spawned by Fabrik from it).
