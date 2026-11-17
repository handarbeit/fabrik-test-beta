# Feature Specification: Beta greeting function for cross-repo spawn e2e

**Feature Branch**: `fabrik/issue-758`
**Created**: 2026-11-17
**Status**: Specified
**Input**: User description: "e2e cross-repo spawn: add HelloE2E202611170521188784 to pkg/greeting (beta side). This is the beta-repo half of an end-to-end regression scenario for cross-repo decomposition, covering on-demand spawn-target repo init (#803) and the per-run-unique fixture token (#1308). A consumer repo, `handarbeit/fabrik-test-alpha`, will call a new exported function from this repo's `pkg/greeting` package. That function must exist here, merged to the default branch, before the alpha change can build. `pkg/greeting/greeting.go` MUST export `HelloE2E202611170521188784() string`, returning exactly `\"e2e-cross-repo-spawn-20261117-052118-8784\"`, appended at the end of the file with a doc comment, following the style of the other `HelloE2E…` functions. `pkg/greeting/greeting_test.go` MUST contain `TestHelloE2E202611170521188784`, asserting the exact returned literal, in the style of the other `TestHelloE2E…` tests (a `const want` and a `got != want` check using `t.Errorf`). Use the function name and string literal exactly as given; a per-run-unique token is baked into both, and the scenario's final assertion depends on this exact text landing in this repo."

## Background

This is the beta-repo half of an end-to-end regression scenario for cross-repo decomposition. The scenario covers on-demand spawn-target repo init (#803) and the per-run-unique fixture token (#1308). A consumer repo, `handarbeit/fabrik-test-alpha`, will call a new exported function from this repo's `pkg/greeting` package. That function must exist here, merged to the default branch, before the alpha change can build.

The parent issue lives in `handarbeit/fabrik-test-alpha` and is blocked on this one. The scenario's final assertion depends on the exact function name and string token landing in this repo, so no substitution is acceptable.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta exports the per-run sentinel greeting (Priority: P1)

The alpha repo's change needs to call a new exported function from beta's `pkg/greeting` package. Beta must provide that function, with a test, on its default branch.

**Why this priority**: It is the only deliverable of this issue; the alpha parent is blocked until it lands.

**Independent Test**: Run `go test ./pkg/greeting/...` in this repo and confirm the new test exists and passes.

**Acceptance Scenarios**:

1. **Given** the greeting package, **When** `HelloE2E202611170521188784()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261117-052118-8784`.
2. **Given** the greeting test file, **When** `go test ./pkg/greeting/...` runs, **Then** `TestHelloE2E202611170521188784` exists and passes.

---

### Edge Cases

- Other `HelloE2E…` functions and tests may be appended concurrently by other runs, producing a merge conflict at the end of the files; both sides must be kept.
- If `pkg/greeting/greeting.go` or `greeting_test.go` did not exist, they would need to be created with the package declaration and the new function/test (both currently exist).

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `pkg/greeting/greeting.go` MUST export `HelloE2E202611170521188784() string`, returning exactly `"e2e-cross-repo-spawn-20261117-052118-8784"`. The function is appended at the end of the file with a doc comment, following the style of the existing `HelloE2E…` functions.
- **FR-002**: `pkg/greeting/greeting_test.go` MUST contain `TestHelloE2E202611170521188784`, which asserts the exact returned literal using a `const want` and a `got != want` check with `t.Errorf`, following the style of the existing `TestHelloE2E…` tests.
- **FR-003**: The function name `HelloE2E202611170521188784` and the string literal `e2e-cross-repo-spawn-20261117-052118-8784` MUST be used exactly as given, with no substitution or invented alternatives.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `go test ./pkg/greeting/...` passes with the new test included.
- **SC-002**: The alpha repo's change can build against the default branch of this repo after this work merges, and the scenario's final assertion finds the exact token text in this repo.

## Assumptions

- The existing `HelloE2E…` functions and tests in `pkg/greeting` are the style reference.
- The linked PR will target the default branch `main`.

## Out of Scope *(optional)*

- Any change other than the new function and its test.
- Any change to the alpha repo (handled by the parent issue).

## Scope

- Only the function and its test. Make no other changes in this repo.
- If `pkg/greeting/greeting.go` or `greeting_test.go` do not exist, create them with the package declaration and the new function/test.
- Other `HelloE2E…` functions may be appended concurrently by other runs. On a merge conflict, keep both sides.

## Linkage

The parent issue is in `handarbeit/fabrik-test-alpha` and is blocked on this one. Mention it by bare number only in any PR text. Do NOT use a closing keyword (`Closes`/`Fixes`/`Resolves`) against it.

---

*Spawned by Fabrik from parent issue handarbeit/fabrik-test-alpha#11555 as a multi-issue decomposition. The parent's plan is at the link above.*
