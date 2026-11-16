# Feature Specification: Add HelloE2E202611161405545760 greeting function

**Feature Branch**: `fabrik/issue-752`
**Created**: 2026-11-16
**Status**: Specified
**Input**: User description: "## Problem

The alpha repo (`handarbeit/fabrik-test-alpha`) needs a new exported function from this repo's greeting package, and it cannot compile until that function exists here. This is the beta-side half of a cross-repo e2e regression scenario (on-demand spawn-target init and per-run-unique fixture token). It must land first.

## Requirements

- Add an exported function `HelloE2E202611161405545760() string` to `pkg/greeting/greeting.go`.
- It MUST return exactly the string literal `e2e-cross-repo-spawn-20261116-140554-5760`.
- Add a test for it in `pkg/greeting/greeting_test.go` asserting that exact returned literal.
- Use the function name and literal verbatim. Copy them. Do not invent, alter or regenerate any token. A final assertion checks for this exact text in this repo's default branch.
- Follow the layout and test style of the existing `HelloE2E…` functions and tests in the same files.

## Scope

- Only the one function and its test. No other changes to this repo.
- Do not change any engine behaviour.

## Acceptance

- `pkg/greeting/greeting.go` on the default branch contains `HelloE2E202611161405545760` returning `e2e-cross-repo-spawn-20261116-140554-5760`.
- `go test ./pkg/greeting/...` passes, including the new test.
- The change is merged to the default branch, so alpha can pin to the merge commit. Note the merge commit SHA when closing.

References for context only: this work is part of handarbeit/fabrik-test-alpha issue #11303. Mention it by bare number. Do not use any closing keyword."

## Background

The alpha repo (`handarbeit/fabrik-test-alpha`) depends on a new exported function in this repo's greeting package and cannot compile until it exists here. This issue is the beta-side half of a cross-repo e2e regression scenario that exercises on-demand spawn-target initialisation and a per-run-unique fixture token. It must land and be merged first so alpha can pin to the resulting merge commit.

The issue was spawned by Fabrik from parent issue handarbeit/fabrik-test-alpha#11303 as part of a multi-issue decomposition. Earlier runs of the same scenario each added one sentinel function and one test to the same two files, so this change follows an established pattern.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta greeting package exposes the new sentinel function (Priority: P1)

The alpha repo's cross-repo e2e scenario needs `HelloE2E202611161405545760` to exist in the beta greeting package, returning a fixed per-run sentinel string, so that alpha can compile against it and the final assertion can find the token in beta's default branch.

**Why this priority**: It is the only deliverable of this issue, and alpha is blocked until it is merged.

**Independent Test**: Run `go test ./pkg/greeting/...` in this repo. The new test passes. Grepping `pkg/greeting/greeting.go` on the default branch finds the function name and the literal.

**Acceptance Scenarios**:

1. **Given** the greeting package, **When** `HelloE2E202611161405545760()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261116-140554-5760`.
2. **Given** the change is merged to the default branch, **When** alpha pins to the merge commit and compiles, **Then** the reference to `HelloE2E202611161405545760` resolves.
3. **Given** the test suite, **When** `go test ./pkg/greeting/...` runs, **Then** all tests pass, including the new one.

---

### Edge Cases

- Any alteration, regeneration or retyping of the function name or literal breaks the final assertion; both must be copied verbatim from this spec.
- The change must not touch other `HelloE2E…` functions or tests, or any file outside the two named files.
- The linked alpha issue #11303 must be mentioned by bare number only, with no closing keyword, in any PR text or commit message.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `pkg/greeting/greeting.go` MUST define an exported function `HelloE2E202611161405545760() string`.
- **FR-002**: `HelloE2E202611161405545760` MUST return exactly the string literal `e2e-cross-repo-spawn-20261116-140554-5760`.
- **FR-003**: `pkg/greeting/greeting_test.go` MUST contain a test for the function that asserts the exact returned literal.
- **FR-004**: The function name and literal MUST be used verbatim as given here, with no invented, altered or regenerated token.
- **FR-005**: The function and test MUST follow the layout and style of the existing `HelloE2E…` functions and tests in the same files.
- **FR-006**: No other changes may be made to this repo: only the one function and its test, and no engine behaviour changes.
- **FR-007**: The change MUST be merged to the default branch (`main`), and the merge commit SHA noted when closing, so alpha can pin to it.
- **FR-008**: References to alpha issue #11303 MUST use the bare number only and MUST NOT use any closing keyword (`Closes`/`Fixes`/`Resolves`).

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: On the default branch, `pkg/greeting/greeting.go` contains `HelloE2E202611161405545760` returning `e2e-cross-repo-spawn-20261116-140554-5760`.
- **SC-002**: `go test ./pkg/greeting/...` passes, including the new test.
- **SC-003**: The diff touches only `pkg/greeting/greeting.go` and `pkg/greeting/greeting_test.go`, plus the engine-generated spec file for this issue.
- **SC-004**: The merge commit SHA is recorded when the issue is closed, and alpha can pin to it.

## Assumptions

- The existing `HelloE2E…` functions and tests in the two files define the layout and style to follow: a doc comment, a one-line function returning the literal, and a test with a `want` constant.
- The engine-generated `specs/<issue>-<slug>/spec.md` file is a normal part of this branch's diff and is not counted as an extra change to the repo.
- "Note the merge commit SHA when closing" is a closing-time action, not a code change.

## Out of Scope *(optional)*

- Any change to the alpha repo, including the pinning of alpha to this repo's merge commit.
- Any change to existing `HelloE2E…` functions or tests.
- Any change to engine behaviour.

## Source References *(optional)*

- Parent issue: handarbeit/fabrik-test-alpha#11303 (mentioned for context only).
- Existing pattern: the `HelloE2E…` functions and tests in `pkg/greeting/greeting.go` and `pkg/greeting/greeting_test.go`.
