# Feature Specification: Add HelloE2E202611072104036462 greeting function and test (e2e cross-repo spawn 20261107-210403-6462)

**Feature Branch**: `fabrik/issue-718`
**Created**: 2026-11-07
**Status**: Specified
**Input**: User description: "Add a run-unique exported greeting function to the `pkg/greeting` package, with a test. This is the beta-side half of a cross-repo e2e regression scenario. The alpha repo (`handarbeit/fabrik-test-alpha`) will call this function once it lands, so this change must be merged to the default branch first."

## Background

This is the beta-side half of a cross-repo e2e regression scenario. It was spawned from parent issue handarbeit/fabrik-test-alpha#10885 as a multi-issue decomposition. The alpha repo (`handarbeit/fabrik-test-alpha`) will call the new function once it lands, so this change must be merged to this repo's default branch first.

The `pkg/greeting` package already holds many run-unique sentinel functions named `HelloE2E<token>`, each returning `e2e-cross-repo-spawn-<token>`, with a matching test for each. This issue adds one more, using the per-run-unique token `20261107-210403-6462`. The token is already fixed, and the scenario's final assertion depends on this exact text landing on the default branch. The function name and literal must therefore be used exactly as written, with nothing altered or generated.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo exposes the run-unique sentinel (Priority: P1)

A maintainer running the cross-repo spawn e2e scenario needs the beta repo's `pkg/greeting` package to expose the new sentinel function, with a passing test. The alpha repo can then call it once this change is merged.

**Why this priority**: It is the only deliverable of this issue. The scenario's cross-repo assertion cannot pass without it.

**Independent Test**: Run `go test ./...` in the beta repo and confirm the new test passes. Then confirm the default branch contains the function name and the literal.

**Acceptance Scenarios**:

1. **Given** the `pkg/greeting` package, **When** `HelloE2E202611072104036462()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261107-210403-6462`.
2. **Given** the new test in `pkg/greeting/greeting_test.go`, **When** `go test ./...` runs, **Then** the test passes and asserts the exact literal.
3. **Given** the change is merged, **When** the default branch is inspected, **Then** it contains both the identifier `HelloE2E202611072104036462` and the literal `e2e-cross-repo-spawn-20261107-210403-6462`.

---

### Edge Cases

- The function name and literal must not be altered, abbreviated or regenerated. A single differing character breaks the scenario's final assertion.
- The new function must not disturb any existing `HelloE2E…` function or test in the package.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The `pkg/greeting` package MUST export a function `HelloE2E202611072104036462() string` that returns exactly the literal `e2e-cross-repo-spawn-20261107-210403-6462`.
- **FR-002**: `pkg/greeting/greeting_test.go` MUST contain a test for `HelloE2E202611072104036462` that asserts the exact returned literal `e2e-cross-repo-spawn-20261107-210403-6462`.
- **FR-003**: The new function and its test MUST follow the style of the existing `HelloE2E…` functions and tests in the same files.
- **FR-004**: The function name and string literal MUST be used exactly as written in FR-001. They MUST NOT be altered or replaced with generated values.
- **FR-005**: The change MUST be limited to the single new function and its test, with no other change to the greeting package or the rest of the repo.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `HelloE2E202611072104036462()` returns the exact string `e2e-cross-repo-spawn-20261107-210403-6462`.
- **SC-002**: `go test ./...` passes, including the new test.
- **SC-003**: After merge, the default branch contains both `HelloE2E202611072104036462` and `e2e-cross-repo-spawn-20261107-210403-6462`.
- **SC-004**: The diff touches only `pkg/greeting/greeting.go` and `pkg/greeting/greeting_test.go`, adding only the new function and its test.

## Assumptions

- The function and test are appended after the most recent existing `HelloE2E…` entries, matching the existing ordering convention.
- The new function's doc comment follows the existing pattern ("returns the sentinel string for the cross-repo spawn regression test").
- This change merges to the default branch (`main`) before the alpha side's change depends on it.

## Out of Scope *(optional)*

- Any change to existing greeting functions or tests.
- Any change outside the `pkg/greeting` package.
- The alpha-side call to this function, which is tracked in the parent issue handarbeit/fabrik-test-alpha#10885.
