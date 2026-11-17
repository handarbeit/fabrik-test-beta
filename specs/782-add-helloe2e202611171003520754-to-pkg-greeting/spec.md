# Feature Specification: Add HelloE2E202611171003520754 to pkg/greeting

**Feature Branch**: `fabrik/issue-782`
**Created**: 2026-11-17
**Status**: Specified
**Input**: User description: "Add HelloE2E202611171003520754 to pkg/greeting (e2e cross-repo spawn 20261117-100352-0754). Problem: The alpha repo (`handarbeit/fabrik-test-alpha`) needs to call a new greeting function from this repo's `pkg/greeting` package, but the function does not exist yet. Alpha cannot import it until it lands here, so this work must merge first. Requirements: Add an exported function to `pkg/greeting/greeting.go`: Name: `HelloE2E202611171003520754` (exactly, with no substitution); Signature: `func HelloE2E202611171003520754() string`; Return value: the literal string `e2e-cross-repo-spawn-20261117-100352-0754` (exactly). Add a test for it in `pkg/greeting/greeting_test.go`. The test must call `HelloE2E202611171003520754()` and assert that it returns exactly `e2e-cross-repo-spawn-20261117-100352-0754`, using that literal verbatim in the assertion. Follow the style of the existing `HelloE2E…` / `HelloIssueN` functions and tests in the same files. Scope: Change only `pkg/greeting/greeting.go` and `pkg/greeting/greeting_test.go`. Append the new function and test. Do not remove or modify any existing function or test. Other concurrent runs also append to these files, so if the rebase conflicts, keep both sides' additions. Do not touch any other file. Acceptance: `HelloE2E202611171003520754()` returns exactly `e2e-cross-repo-spawn-20261117-100352-0754`. `go test ./pkg/greeting/...` passes, including the new test. The change is merged to the beta repo's default branch, so the alpha repo can depend on it. Context: This is the beta-side half of an e2e regression scenario for cross-repo decomposition (parent: `handarbeit/fabrik-test-alpha#11705`). Do not use a different function name or string. The scenario's final assertion looks for this exact text in this repo."

## Background

The alpha repo (`handarbeit/fabrik-test-alpha`) needs to call a new greeting function from this repo's `pkg/greeting` package. Alpha cannot import the function until it exists here, so this work must merge first. This issue is the beta-side half of an end-to-end regression scenario for cross-repo decomposition (parent: `handarbeit/fabrik-test-alpha#11705`). The scenario's final assertion looks for the exact function name and string in this repo, so neither may vary.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta exposes the sentinel greeting (Priority: P1)

A consumer in the alpha repo can call `HelloE2E202611171003520754()` from `pkg/greeting` and receive the agreed sentinel string.

**Why this priority**: It is the only deliverable, and alpha's dependent work is blocked until it lands on the default branch.

**Independent Test**: Run `go test ./pkg/greeting/...`; the new test calls the function and asserts the exact returned string.

**Acceptance Scenarios**:

1. **Given** the change is merged, **When** `HelloE2E202611171003520754()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261117-100352-0754`.
2. **Given** the change is merged, **When** `go test ./pkg/greeting/...` runs, **Then** all tests pass, including the new one.
3. **Given** other concurrent runs have appended their own functions and tests, **When** this change is rebased, **Then** both sides' additions are retained.

---

### Edge Cases

- A rebase conflict in either file caused by concurrent appends: keep both sides' additions; never drop or alter another run's function or test.
- Near-identical function names or strings from sibling runs: the name and string here must match the issue exactly, with no substitution.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `pkg/greeting/greeting.go` exports `func HelloE2E202611171003520754() string`, with that exact name.
- **FR-002**: The function returns exactly the literal `e2e-cross-repo-spawn-20261117-100352-0754`.
- **FR-003**: `pkg/greeting/greeting_test.go` contains a test that calls `HelloE2E202611171003520754()` and asserts the result equals the literal `e2e-cross-repo-spawn-20261117-100352-0754`, written verbatim in the assertion.
- **FR-004**: The function and test follow the style of the existing `HelloE2E…` functions and tests in the same files, and are appended after the existing content.
- **FR-005**: No existing function or test is removed or modified.
- **FR-006**: Only `pkg/greeting/greeting.go` and `pkg/greeting/greeting_test.go` are changed.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `go test ./pkg/greeting/...` passes with the new test included.
- **SC-002**: The exact string `e2e-cross-repo-spawn-20261117-100352-0754` and the exact name `HelloE2E202611171003520754` are present in this repo on its default branch after merge.
- **SC-003**: The diff touches exactly two files, and every existing function and test is unchanged.

## Assumptions

- The existing `HelloE2E…` functions and tests in the two files define the style to follow.
- The default branch is `main`, and merging there is what makes the function available to alpha.

## Out of Scope

- Any change to the alpha repo.
- Any file other than the two named, including docs and configuration.
- Modifying, reordering or removing existing greeting functions or tests.
