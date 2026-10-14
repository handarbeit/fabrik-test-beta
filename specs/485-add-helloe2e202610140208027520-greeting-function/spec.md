# Feature Specification: Add HelloE2E202610140208027520 greeting function and test (e2e cross-repo spawn)

**Feature Branch**: `fabrik/issue-485`
**Created**: 2026-10-13
**Status**: Specified
**Input**: User description: "Add one exported function to the greeting package in this repo and test it. The alpha repo (`handarbeit/fabrik-test-alpha`) will import it, so this change must land on this repo's default branch first.

Requirements:
- In `pkg/greeting/greeting.go`, add an exported function `HelloE2E202610140208027520() string` that returns exactly the literal `e2e-cross-repo-spawn-20261014-020802-7520`.
- In `pkg/greeting/greeting_test.go`, add a test that calls `HelloE2E202610140208027520()` and asserts it returns exactly `e2e-cross-repo-spawn-20261014-020802-7520`.
- Follow the style of the existing `HelloE2E…` functions and their tests in the same files.

Constraints:
- Use the function name `HelloE2E202610140208027520` and the string literal `e2e-cross-repo-spawn-20261014-020802-7520` verbatim. They contain a per-run-unique token. Do not invent alternatives, re-derive, or alter any character. The scenario's final assertion checks that this exact text appears in this repo's default branch.
- Make no other change to this repo: no additional greeting functions, no generalisation of the fixture.

Acceptance:
1. `greeting.HelloE2E202610140208027520()` returns `e2e-cross-repo-spawn-20261014-020802-7520`.
2. A test in `pkg/greeting/greeting_test.go` covers the function.
3. `go build ./...` and `go test ./...` pass.
4. The change is merged to this repo's default branch, so the alpha repo can bump its module pin to a commit containing the function."

## Background

This issue is one half of a cross-repo e2e spawn scenario. A parent issue in `handarbeit/fabrik-test-alpha` (#7929) was decomposed into work in this repo (beta) and in alpha. The alpha repo will import a new sentinel function from beta's `pkg/greeting` package, so the function must exist on beta's default branch before alpha can bump its module pin to a commit that contains it.

The function name and returned string embed a per-run-unique token. The scenario's final assertion checks that this exact text appears on this repo's default branch, so any deviation from the specified identifiers breaks the scenario. The `pkg/greeting` package already holds a series of sentinel functions of the same shape, each paired with a test. This change adds one more in that established pattern.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Expose the run-unique sentinel from the greeting package (Priority: P1)

A consumer of the greeting package, namely the alpha repo's cross-repo spawn scenario, calls `greeting.HelloE2E202610140208027520()` and receives the exact run-unique sentinel string. The consumer can then pin a module version containing it.

**Why this priority**: This is the whole change. Without it the alpha repo cannot complete its side of the scenario.

**Independent Test**: Run the new unit test in `pkg/greeting`, then run `go build ./...` and `go test ./...` for the module. All must pass.

**Acceptance Scenarios**:

1. **Given** the greeting package on this repo's default branch, **When** a caller invokes `HelloE2E202610140208027520()`, **Then** it returns exactly `e2e-cross-repo-spawn-20261014-020802-7520`.
2. **Given** the new test in `pkg/greeting/greeting_test.go`, **When** `go test ./...` runs, **Then** the test passes, and it would fail if the function returned any other string.
3. **Given** the change is merged, **When** the alpha repo bumps its module pin to a commit on this repo's default branch, **Then** the function is present at that commit.

---

### Edge Cases

- The identifier or literal contains a transposed, dropped or altered character. The scenario's final assertion would then fail, so both must be copied verbatim from this spec.
- A sentinel function from an earlier run is accidentally modified or removed. Existing functions and tests must remain untouched.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The greeting package MUST export a function `HelloE2E202610140208027520() string` that returns exactly `e2e-cross-repo-spawn-20261014-020802-7520`.
- **FR-002**: The greeting package's test file MUST contain a test that calls `HelloE2E202610140208027520()` and asserts the result equals exactly `e2e-cross-repo-spawn-20261014-020802-7520`.
- **FR-003**: The new function and test MUST follow the style of the existing `HelloE2E…` functions and tests in the same files.
- **FR-004**: The function name and string literal MUST be used verbatim, with no re-derived or altered characters.
- **FR-005**: The module MUST build and its full test suite MUST pass with the change applied.
- **FR-006**: The change MUST make no other modifications to the repo: no additional greeting functions and no generalisation of the fixture.

### Key Entities *(if applicable)*

- **Sentinel function**: An exported, zero-argument greeting-package function that returns a fixed per-run-unique string, used by the cross-repo spawn scenario to prove a change landed on the default branch.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `HelloE2E202610140208027520()` returns `e2e-cross-repo-spawn-20261014-020802-7520`, confirmed by a passing automated test.
- **SC-002**: The module builds and the entire existing test suite passes, with the new test included and no existing test altered.
- **SC-003**: After merge, the exact function name and literal are present on this repo's default branch, so the alpha repo's scenario assertion succeeds.
- **SC-004**: The diff touches only the greeting package's source and test files, plus this issue's spec file.

## Assumptions

- The new function and test are appended after the most recent existing `HelloE2E…` entries, matching the existing doc-comment and test layout.
- "Default branch" is `main`, as stated for this repo.
- The engine's own persisted spec file for this issue (`specs/485-…/spec.md`) is an expected part of the change and does not count as "other change".

## Out of Scope *(optional)*

- Any additional greeting functions or helpers.
- Refactoring or generalising the sentinel-function pattern.
- Changes to the alpha repo, including bumping its module pin.
- Changes to existing sentinel functions or tests.
