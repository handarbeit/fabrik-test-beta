# Feature Specification: Add HelloE2E202610191543299409 greeting function (e2e cross-repo spawn)

**Feature Branch**: `fabrik/issue-577`
**Created**: 2026-10-19
**Status**: Specified
**Input**: User description: "Add one exported function and one test to the greeting package in this repo. This is the beta half of a cross-repo end-to-end regression (parent: handarbeit/fabrik-test-alpha#9076). The alpha repo cannot compile its change until this one has landed on the default branch, so this work comes first. In `pkg/greeting/greeting.go`, add an exported function with exactly this name and signature, returning exactly this literal: `// HelloE2E202610191543299409 returns the sentinel string for the cross-repo spawn regression test.` / `func HelloE2E202610191543299409() string { return \"e2e-cross-repo-spawn-20261019-154329-9409\" }`. Follow the style of the existing `HelloE2E...` functions in that file. In `pkg/greeting/greeting_test.go`, add `TestHelloE2E202610191543299409`, following the existing per-function tests: declare `const want = \"e2e-cross-repo-spawn-20261019-154329-9409\"` and compare it with the call result. Use the function name `HelloE2E202610191543299409` and the literal `e2e-cross-repo-spawn-20261019-154329-9409` verbatim; they carry a per-run-unique token, so do not paraphrase, regenerate, reformat or substitute them. A final assertion in the parent scenario checks for this exact text in this repo's default branch. Scope: change only the one function and its test; do not touch the existing functions or tests, and make no other changes to this repo. No new dependencies; the module is `github.com/handarbeit/fabrik-test-beta` with `go 1.22`. Acceptance: `HelloE2E202610191543299409()` returns `e2e-cross-repo-spawn-20261019-154329-9409`; `go test ./...` passes, including the new test; the change is merged to the default branch, so alpha can pin to a commit that contains it."

## Background

This issue is the beta-repo half of a cross-repo end-to-end regression for Fabrik's multi-issue spawn behavior, spawned from parent `handarbeit/fabrik-test-alpha#9076`. The alpha repo's change depends on a sentinel function that must exist in this repo's `greeting` package, and alpha cannot compile until that function has landed on this repo's default branch (`main`). This work therefore comes first.

The `greeting` package already holds many sentinel functions of the same shape (`HelloE2E…`), each returning a per-run-unique string, each with its own test. This change adds one more with a new per-run-unique token. A final assertion in the parent scenario looks for the exact function name and literal text on this repo's default branch, so the token must be reproduced exactly.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Sentinel function exists and is verified (Priority: P1)

The parent cross-repo scenario needs `HelloE2E202610191543299409` in this repo's `greeting` package, returning the per-run-unique string, with a passing test, merged to `main`.

**Why this priority**: It is the only deliverable of this issue. Alpha's compilation and the parent scenario's final assertion both depend on it.

**Independent Test**: Run `go test ./...` in this repo and confirm the new test passes; inspect `main` after merge for the exact function name and literal.

**Acceptance Scenarios**:

1. **Given** the `greeting` package, **When** `HelloE2E202610191543299409()` is called, **Then** it returns `e2e-cross-repo-spawn-20261019-154329-9409`.
2. **Given** the new test `TestHelloE2E202610191543299409` declares `const want = "e2e-cross-repo-spawn-20261019-154329-9409"`, **When** `go test ./...` runs, **Then** the test passes along with all existing tests.
3. **Given** the change is merged, **When** the default branch is inspected, **Then** it contains the exact text `HelloE2E202610191543299409` and `e2e-cross-repo-spawn-20261019-154329-9409`.

---

### Edge Cases

- A mistyped, reformatted or regenerated token would compile and pass its own test but fail the parent scenario's exact-text assertion, so the name and literal must be copied verbatim.
- The existing `HelloE2E…` functions and tests must remain unchanged; the diff must be purely additive.
- The new function must not collide with any existing identifier (none currently uses this token).

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The `greeting` package MUST export a function `HelloE2E202610191543299409` with signature `func() string` that returns exactly `e2e-cross-repo-spawn-20261019-154329-9409`.
- **FR-002**: The function MUST carry the doc comment `HelloE2E202610191543299409 returns the sentinel string for the cross-repo spawn regression test.` and follow the style of the existing `HelloE2E…` functions.
- **FR-003**: The package's tests MUST include `TestHelloE2E202610191543299409`, which declares `const want = "e2e-cross-repo-spawn-20261019-154329-9409"` and compares it with the function's return value, matching the style of the existing per-function tests.
- **FR-004**: The function name and the literal MUST be used verbatim, with no paraphrasing, regeneration, reformatting or substitution.
- **FR-005**: No existing function or test may be modified, and no other change may be made to the repository.
- **FR-006**: The change MUST add no new dependencies and keep the module as `github.com/handarbeit/fabrik-test-beta` at `go 1.22`.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `go test ./...` passes with the new test included.
- **SC-002**: The repository's default branch contains the exact strings `HelloE2E202610191543299409` and `e2e-cross-repo-spawn-20261019-154329-9409` once the change is merged.
- **SC-003**: The diff against `main` consists only of one added function and one added test.

## Assumptions

- The new function is appended after the most recent existing `HelloE2E…` function, and the new test after the most recent existing per-function test, matching existing convention.
- The existing per-function tests are the style reference for the new test; the spec does not prescribe further detail.

## Out of Scope

- Any change to existing functions, tests, documentation, build configuration or dependencies.
- The alpha repo's half of the scenario (tracked in `handarbeit/fabrik-test-alpha#9076`).

## Source References

- Parent issue: `handarbeit/fabrik-test-alpha#9076`
- Spawned by Fabrik from the parent issue as a multi-issue decomposition.
