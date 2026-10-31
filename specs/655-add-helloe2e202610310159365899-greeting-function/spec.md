# Feature Specification: Add per-run-unique greeting function for the cross-repo spawn e2e scenario

**Feature Branch**: `fabrik/issue-655`
**Created**: 2026-10-30
**Status**: Specified
**Input**: User description: "Add HelloE2E202610310159365899 greeting function and test (cross-repo spawn e2e)"

## Background

This issue is the beta-repo half of a cross-repo e2e regression scenario. It was spawned from the alpha parent issue `handarbeit/fabrik-test-alpha#10328`. The alpha repo imports `github.com/handarbeit/fabrik-test-beta/pkg/greeting` and cannot call the new function until it exists on this repo's default branch. This change must land first. It guards the earlier fixes #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token).

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Alpha repo can depend on the new per-run sentinel (Priority: P1)

The cross-repo spawn scenario needs a greeting function, with a per-run-unique name and return value, to exist on this repo's default branch. The alpha repo then imports and calls it, and the scenario's final assertion checks that the exact sentinel text landed here.

**Why this priority**: It is the only deliverable of this issue. Without it the alpha half of the scenario cannot compile or assert.

**Independent Test**: Run `go test ./...` in this repo. The new test passes, and the function returns exactly `e2e-cross-repo-spawn-20261031-015936-5899`.

**Acceptance Scenarios**:

1. **Given** the greeting package on the default branch, **When** `HelloE2E202610310159365899()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261031-015936-5899`.
2. **Given** the change is merged to the default branch, **When** the alpha repo imports `github.com/handarbeit/fabrik-test-beta/pkg/greeting` and references the function, **Then** it resolves.
3. **Given** the greeting source or test file does not yet exist in this repo (on-demand initialisation of a spawn target, #803), **When** the change is made, **Then** the missing file(s) are created with `package greeting` containing the new function and test.

---

### Edge Cases

- A spawn-target repo with no `pkg/greeting` files yet: the files are created from scratch rather than the change failing.
- The function name and string literal are per-run unique. Any alteration, such as shortening or regenerating them, breaks the scenario's final assertion.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `pkg/greeting/greeting.go` MUST export `HelloE2E202610310159365899() string`, returning exactly the string `e2e-cross-repo-spawn-20261031-015936-5899`.
- **FR-002**: `pkg/greeting/greeting_test.go` MUST include a test, `TestHelloE2E202610310159365899`, that asserts the function returns exactly `e2e-cross-repo-spawn-20261031-015936-5899`.
- **FR-003**: Use the function name and the string literal exactly as written above. Do not regenerate, shorten, or alter them. The scenario's final assertion depends on this exact text landing in this repo.
- **FR-004**: Follow the existing file conventions. Append the new function at the end of `greeting.go`, with a doc comment in the same style as the earlier `HelloE2E<timestamp>` functions. Append the new test at the end of `greeting_test.go`, in the same per-function shape as the existing tests (a `const want` and a `got != want` check).
- **FR-005**: If `pkg/greeting/greeting.go` or `pkg/greeting/greeting_test.go` do not exist yet in this repo (on-demand initialisation of a spawn target, #803), they MUST be created with `package greeting` and the new function and test.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `HelloE2E202610310159365899()` returns exactly `e2e-cross-repo-spawn-20261031-015936-5899`.
- **SC-002**: `go test ./...` passes in this repo, including `TestHelloE2E202610310159365899`.
- **SC-003**: The change is merged to this repo's default branch, so the alpha repo can depend on it.

## Assumptions

- The earlier `HelloE2E<timestamp>` functions and their tests establish the doc-comment and test-shape conventions that FR-004 refers to.
- The default branch is `main`.
- The PR for this issue carries `Closes #655` and mentions the parent `handarbeit/fabrik-test-alpha#10328` by bare reference only.

## Out of Scope *(optional)*

- Any change to existing greeting functions or their behaviour.
- Any file other than `greeting.go` and `greeting_test.go`.
- Any change to the Fabrik engine.

## Scope

- In scope: the new function and its test only.
- Out of scope: any other change to greeting behaviour, any other file, and any change to the Fabrik engine.

## Acceptance

1. `HelloE2E202610310159365899()` returns exactly `e2e-cross-repo-spawn-20261031-015936-5899`.
2. `go test ./...` passes, including the new test.
3. The change is merged to this repo's default branch so the alpha repo can depend on it.

## Notes

- Mention the parent `handarbeit/fabrik-test-alpha#10328` for context in the PR body by bare reference only. Do NOT use a closing keyword (`Closes`/`Fixes`/`Resolves`) for it.
---

*Spawned by Fabrik from parent issue handarbeit/fabrik-test-alpha#10328 as a multi-issue decomposition. The parent's plan is at the link above.*
