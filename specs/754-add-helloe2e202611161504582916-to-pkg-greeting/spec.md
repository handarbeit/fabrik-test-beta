# Feature Specification: Add per-run greeting function to the beta repo

**Feature Branch**: `fabrik/issue-754`
**Created**: 2026-11-16
**Status**: Specified
**Input**: User description: "Add one exported function and one test to the greeting package of this repo (`handarbeit/fabrik-test-beta`), following the pattern of the existing `HelloE2E…` functions. Add an exported function `HelloE2E202611161504582916() string` to `pkg/greeting/greeting.go` returning exactly the literal `e2e-cross-repo-spawn-20261116-150458-2916`, plus a test `TestHelloE2E202611161504582916` in `pkg/greeting/greeting_test.go`. This issue was spawned from handarbeit/fabrik-test-alpha#11358 (an e2e regression test for cross-repo decomposition, #803 and #1308). Its work must land before the alpha repo can import the new function, so it has to be completed and merged first."

## Background

This issue was spawned from handarbeit/fabrik-test-alpha#11358, an end-to-end regression test for cross-repo decomposition (#803 and #1308). The alpha repo's work depends on a per-run sentinel function existing in this repo's default branch. The parent scenario's final assertion matches the exact function name and string literal in this repo's default branch. This change must therefore be completed and merged before the alpha-side change starts. The function name and literal both embed a per-run-unique token, so they must be used verbatim.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Per-run sentinel greeting is available in the beta repo (Priority: P1)

A maintainer (or the parent e2e scenario) needs the greeting package in this repo to expose the per-run sentinel function so that the alpha repo can later import it and the scenario can assert on its presence.

**Why this priority**: It is the only deliverable of this issue, and the parent scenario is blocked until it is merged.

**Independent Test**: Run `go test ./...` in this repo. The new test `TestHelloE2E202611161504582916` runs and passes. Inspect the default branch after merge for the function name and the literal.

**Acceptance Scenarios**:

1. **Given** the change is applied, **When** `HelloE2E202611161504582916()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261116-150458-2916`.
2. **Given** the change is applied, **When** `go test ./...` is run, **Then** it passes and `TestHelloE2E202611161504582916` is among the tests run.
3. **Given** the change is merged, **When** the default branch is inspected, **Then** it contains both `HelloE2E202611161504582916` and the literal `e2e-cross-repo-spawn-20261116-150458-2916`.

---

### Edge Cases

- Any existing `HelloE2E…` function or test must remain untouched, with its order and content unchanged.
- The function name and literal must match the spec character for character. A typo or a "corrected" token would break the parent scenario's final assertion.
- This issue must not carry a closing keyword for alpha's issue. The alpha repo is referenced by bare issue number only.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: Add an exported function `HelloE2E202611161504582916() string` to `pkg/greeting/greeting.go`.
- **FR-002**: The function MUST return exactly the literal `e2e-cross-repo-spawn-20261116-150458-2916`.
- **FR-003**: Add a test for it in `pkg/greeting/greeting_test.go`, named `TestHelloE2E202611161504582916`, following the existing one-test-per-function pattern (a `const want` compared with the returned value). It MUST pass under `go test ./...`.
- **FR-004**: The change is purely additive. Append the function at the end of `greeting.go` and the test at the end of `greeting_test.go`. Do not modify, reorder or remove any existing function or test.
- **FR-005**: Use the function name and the string literal verbatim. A per-run-unique token is already baked into them. Do not invent, alter, re-derive or "improve" either one. The parent scenario's final assertion matches this exact text in this repo's default branch.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: Calling `HelloE2E202611161504582916()` returns `e2e-cross-repo-spawn-20261116-150458-2916`.
- **SC-002**: `go test ./...` passes, and the new test is among the tests run.
- **SC-003**: After the change is merged, the default branch contains both `HelloE2E202611161504582916` and the literal `e2e-cross-repo-spawn-20261116-150458-2916`.

## Assumptions

- `pkg/greeting` already exists with `greeting.go` and `greeting_test.go`, and this repo's `go.mod` declares `go 1.22` with no dependencies.
- The alpha repo mentions this work by bare issue number only. This issue should not carry any closing keyword for alpha's issue.

## Out of Scope

- Any change to other functions, files or modules in this repo.
- Anything in the alpha repo. The alpha-side change is tracked by handarbeit/fabrik-test-alpha#11358 and starts only after this issue is closed.

## Source References

- handarbeit/fabrik-test-alpha#11358 (parent e2e cross-repo spawn scenario; mentioned for context only)

---

*Spawned by Fabrik from parent issue handarbeit/fabrik-test-alpha#11358 as a multi-issue decomposition. The parent's plan is at the link above.*
