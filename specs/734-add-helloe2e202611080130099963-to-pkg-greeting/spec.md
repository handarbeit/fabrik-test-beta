# Feature Specification: Add per-run-unique greeting function (beta side of e2e cross-repo spawn)

**Feature Branch**: `fabrik/issue-734`
**Created**: 2026-11-07
**Status**: Specified
**Input**: User description: "Add HelloE2E202611080130099963 to pkg/greeting (e2e cross-repo spawn). This issue was spawned from handarbeit/fabrik-test-alpha#11031, an end-to-end regression scenario for cross-repo decomposition (#803: on-demand spawn-target repo init; #1308: per-run-unique fixture token). The alpha repo cannot import this function until it exists here, so this change must land first. Add an exported function `HelloE2E202611080130099963() string` to `pkg/greeting/greeting.go` that returns exactly the literal string `e2e-cross-repo-spawn-20261108-013009-9963`, with a test in `pkg/greeting/greeting_test.go`."

## Background

This issue is the beta-repo half of an end-to-end regression scenario for cross-repo decomposition. The parent issue, handarbeit/fabrik-test-alpha#11031, was decomposed by Fabrik and this issue was spawned in the beta repo (covering #803, on-demand initialisation of the spawn-target repo, and #1308, the per-run-unique fixture token).

The alpha repo cannot import the function until it exists in this repo, so this change must land first. The function name and its return string both carry a per-run-unique token. The scenario's final assertion depends on exactly that text appearing in this repo, so both must be used verbatim.

When the PR for this issue merges, the alpha repo's `go.mod` pin will be bumped to the merge commit.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Expose the per-run-unique greeting function (Priority: P1)

A consumer, here the alpha repo's scenario, can call `HelloE2E202611080130099963()` from this repo's greeting package and receive the agreed sentinel string.

**Why this priority**: It is the only deliverable. The alpha repo is blocked until it lands, and the scenario's final assertion depends on it.

**Independent Test**: Run the greeting package's tests. The new test passes and asserts the exact return value.

**Acceptance Scenarios**:

1. **Given** the greeting package, **When** `HelloE2E202611080130099963()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261108-013009-9963`.
2. **Given** the full test suite, **When** `go test ./...` is run, **Then** it passes, including the new test.
3. **Given** the change is merged, **When** the existing greeting functions are inspected, **Then** none has been changed or removed.

---

### Edge Cases

- The function name or literal is altered, for example a typo, a different timestamp or a re-generated token. The scenario's final assertion would then fail, so both are used exactly as given.
- The change is placed in a new file instead of the existing greeting and test files. This breaks the convention the existing `HelloE2E…` functions follow, so it is not allowed.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: Add an exported function `HelloE2E202611080130099963() string` to `pkg/greeting/greeting.go`. Append it to the existing file, following the style of the existing `HelloE2E…` functions. Do not create new files for it.
- **FR-002**: The function MUST return exactly the literal string `e2e-cross-repo-spawn-20261108-013009-9963`.
- **FR-003**: Add a test for it in `pkg/greeting/greeting_test.go`, appended to the existing file and following its existing style. The test asserts that `HelloE2E202611080130099963()` returns exactly `e2e-cross-repo-spawn-20261108-013009-9963`. It MUST pass.
- **FR-004**: Use the function name and the string literal exactly as given above, verbatim. Do not invent substitute names or values. A per-run-unique token is baked into them, and the scenario's final assertion depends on this exact text landing in this repo.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: Calling `HelloE2E202611080130099963()` returns exactly `e2e-cross-repo-spawn-20261108-013009-9963`.
- **SC-002**: A test for the function exists in `pkg/greeting/greeting_test.go`, and `go test ./...` passes.
- **SC-003**: No existing function is changed or removed; the diff only adds to the two existing files.

## Assumptions

- The existing `HelloE2E…` functions and their tests in the greeting package set the style to follow, and the file layout stays as it is.
- The function name and string literal given in this issue are final and correct. No re-generation or validation of the token is needed.

## Out of Scope

- Any other greeting functions or changes to existing functions.
- Changes to the alpha repo (handled by the parent issue handarbeit/fabrik-test-alpha#11031).
- Bumping the alpha repo's `go.mod` pin; this happens after this issue's PR merges.

## Source References

- Parent issue, for context only: handarbeit/fabrik-test-alpha#11031. The PR must mention it by reference only and must NOT use a closing keyword (`Closes`/`Fixes`/`Resolves`) for it.

---

*Spawned by Fabrik from parent issue handarbeit/fabrik-test-alpha#11031 as a multi-issue decomposition. The parent's plan is at the link above.*
