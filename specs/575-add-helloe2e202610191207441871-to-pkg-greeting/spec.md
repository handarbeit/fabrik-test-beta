# Feature Specification: Add HelloE2E202610191207441871 to pkg/greeting

**Feature Branch**: `fabrik/issue-575`
**Created**: 2026-10-19
**Status**: Specified
**Input**: User description: "

## Summary

Add one exported function and one test to the `pkg/greeting` package. This is the beta-side half of a cross-repo e2e regression scenario (parent issue: handarbeit/fabrik-test-alpha#8983). The parent repo imports this package and cannot build until this function exists, so this change must land first.

## Requirements

1. In `pkg/greeting/greeting.go`, append an exported function `HelloE2E202610191207441871() string` that returns exactly the literal string `\"e2e-cross-repo-spawn-20261019-120744-1871\"`. Append it at the end of the file, following the shape of the existing `HelloXxx` functions (a doc comment, then a one-line function returning the literal).
2. In `pkg/greeting/greeting_test.go`, append a test `TestHelloE2E202610191207441871` that asserts the function returns exactly `\"e2e-cross-repo-spawn-20261019-120744-1871\"`. Follow the existing test pattern: `const want = \"...\"`, then `if got := HelloE2E202610191207441871(); got != want { t.Errorf(\"HelloE2E202610191207441871() = %q; want %q\", got, want) }`.
3. Use the function name `HelloE2E202610191207441871` and the string literal `e2e-cross-repo-spawn-20261019-120744-1871` **verbatim**. A per-run-unique token is already baked into both. Do not alter them or invent your own, because the scenario's final assertion depends on this exact text landing in this repo.

## Scope

- Touch only `pkg/greeting/greeting.go` and `pkg/greeting/greeting_test.go`.
- Do not modify any existing function, test or other file. Do not add dependencies.

## Acceptance

- `go build ./...` and `go test ./...` pass in this repo.
- Calling `HelloE2E202610191207441871()` returns exactly `e2e-cross-repo-spawn-20261019-120744-1871`.
- The new test passes and would fail if the returned literal changed.
- The change is merged to the default branch, so the parent repo can reference it by commit.

Mention handarbeit/fabrik-test-alpha#8983 for context only. Do NOT use a closing keyword (`Closes`/`Fixes`/`Resolves`) on that reference.
---

*Spawned by Fabrik from parent issue handarbeit/fabrik-test-alpha#8983 as a multi-issue decomposition. The parent's plan is at the link above.*"

## Background

This issue is the beta-side half of a cross-repo end-to-end regression scenario that exercises Fabrik's cross-repo child-spawn path. The parent issue (handarbeit/fabrik-test-alpha#8983) lives in the alpha repo, which imports the `pkg/greeting` package from this repo. The alpha repo cannot build until this package exposes the new sentinel function, so this change must be merged before the parent's work can proceed.

The function name and returned literal each carry a per-run-unique token (`20261019-120744-1871`). The scenario's final assertion checks that this exact text landed in this repo, so the token must be used exactly as given and never altered or regenerated.

The package already contains similar sentinel functions (`HelloXxx`), each with a doc comment and a one-line body returning a literal, and a matching test for each. The new function follows that established shape.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta exposes the run-unique sentinel (Priority: P1)

A consumer of `pkg/greeting` (the alpha repo in the e2e scenario) calls `HelloE2E202610191207441871()` and receives the run-unique literal, proving this change has landed in beta's default branch.

**Why this priority**: It is the only deliverable of this issue; the parent repo is blocked until it exists on the default branch.

**Independent Test**: Run `go test ./pkg/greeting/...` in this repo. The new test passes, and changing the returned literal makes it fail.

**Acceptance Scenarios**:

1. **Given** the change is merged to the default branch, **When** a consumer calls `HelloE2E202610191207441871()`, **Then** it returns exactly `e2e-cross-repo-spawn-20261019-120744-1871`.
2. **Given** the change is applied, **When** `go build ./...` and `go test ./...` run in this repo, **Then** both pass.
3. **Given** the returned literal were altered in any way, **When** `TestHelloE2E202610191207441871` runs, **Then** it fails.

---

### Edge Cases

- The function name or literal is mistyped or regenerated with a different token: the parent repo's final assertion fails, so both must match the issue text character for character.
- Existing `HelloXxx` functions and tests must be unaffected; the change is purely additive.
- The reference to the parent issue must not use a closing keyword (`Closes`/`Fixes`/`Resolves`) in the PR or elsewhere, so the parent is not closed by merging this change.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The `pkg/greeting` package exports a function `HelloE2E202610191207441871() string` that returns exactly the string `e2e-cross-repo-spawn-20261019-120744-1871`.
- **FR-002**: The function is appended at the end of `pkg/greeting/greeting.go`, in the same shape as the existing `HelloXxx` functions (a doc comment, then a one-line function returning the literal).
- **FR-003**: `pkg/greeting/greeting_test.go` contains a test `TestHelloE2E202610191207441871` that asserts the function's return value equals the literal, following the existing test pattern (`const want`, then a comparison with `t.Errorf("HelloE2E202610191207441871() = %q; want %q", got, want)` on mismatch), appended at the end of the file.
- **FR-004**: The function name and the string literal are used verbatim as given in this issue, with no alteration or substitution.
- **FR-005**: Only `pkg/greeting/greeting.go` and `pkg/greeting/greeting_test.go` are modified; no existing function, test or other file is changed, and no dependencies are added.
- **FR-006**: Any reference to handarbeit/fabrik-test-alpha#8983 in the PR is for context only and does not use a closing keyword.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `go build ./...` and `go test ./...` pass in this repo with the change applied.
- **SC-002**: Calling `HelloE2E202610191207441871()` returns a string equal to `e2e-cross-repo-spawn-20261019-120744-1871` with zero differences.
- **SC-003**: The new test passes with the correct literal and fails if the literal is changed.
- **SC-004**: The change is merged to the default branch, so the parent repo can reference it by commit.
- **SC-005**: The diff touches exactly two files, and the parent issue handarbeit/fabrik-test-alpha#8983 remains open after merge.

## Assumptions

- The existing `HelloXxx` functions and tests in `pkg/greeting` provide the shape to follow; matching that shape is sufficient for consistency.
- The token `20261019-120744-1871` in the issue is final and correct; no validation or regeneration is wanted.
- The parent repo will consume this change by referencing a commit on the default branch (`main`).

## Out of Scope

- Any change to the alpha repo or the parent issue (handarbeit/fabrik-test-alpha#8983).
- Modifying, renaming or removing any existing function or test in `pkg/greeting`.
- Adding dependencies, other files, documentation or refactors.

## Source References

- Parent issue: handarbeit/fabrik-test-alpha#8983 (mentioned for context only)
