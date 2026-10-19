# Feature Specification: Add HelloE2E202610192019235408 to pkg/greeting (e2e cross-repo spawn 20261019-201923-5408)

**Feature Branch**: `fabrik/issue-579`
**Created**: 2026-10-19
**Status**: Specified
**Input**: User description:

> ## Problem
>
> The alpha repo (`handarbeit/fabrik-test-alpha`, issue #9131) needs to call a new greeting function from this repo. Alpha cannot import it until it exists on this repo's default branch, so this change must land first. It is the beta-side half of an end-to-end regression test for cross-repo spawn (#803, #1308).
>
> ## Requirements
>
> 1. In `pkg/greeting/greeting.go`, add the exported function `HelloE2E202610192019235408() string` that returns exactly the string `e2e-cross-repo-spawn-20261019-201923-5408`. Follow the existing style: one function per run, with a doc comment.
> 2. In `pkg/greeting/greeting_test.go`, add a test for it in the existing style: `const want = "e2e-cross-repo-spawn-20261019-201923-5408"`, call `HelloE2E202610192019235408()`, and use `t.Errorf` if the result differs. Use only the stdlib `testing` package.
> 3. Use the function name and string literal exactly as written above. Do not rename, normalise or "improve" them. The per-run token `20261019-201923-5408` is part of the literal, and `20261019201923` plus `5408` is part of the function name. A later check verifies that this exact text appears in this repo's default branch.
> 4. Work from a fresh fetch of this repo's default branch. Place the new function and test after the latest existing `HelloE2E2026…` entries so nothing is lost or reordered.
>
> ## Scope
>
> - Only add the new function and its test. Do not change, rename or remove any existing greeting function.
> - Do not change any other file.
>
> ## Acceptance
>
> - `HelloE2E202610192019235408()` returns `e2e-cross-repo-spawn-20261019-201923-5408`.
> - `go test ./...` passes, including the new test.
> - The change is merged to this repo's default branch. A pull request that is only open is not enough, because the alpha issue stays blocked until this issue is closed.
> - The default branch contains the text `HelloE2E202610192019235408` and `e2e-cross-repo-spawn-20261019-201923-5408` verbatim.

## Background

This repository is the beta half of an end-to-end regression test for cross-repo spawn (#803, #1308). The alpha repo (`handarbeit/fabrik-test-alpha`, issue #9131) must call a new greeting function that lives here. Alpha cannot import the function until it exists on this repo's default branch, so this issue has to be merged first. Alpha's issue stays blocked until this one is closed. Each test run adds one uniquely named sentinel function, and a later check searches the default branch for the exact function name and string, which is why exact text matters.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Publish the per-run sentinel greeting (Priority: P1)

A maintainer running the cross-repo spawn regression test needs this repo's default branch to expose the sentinel function for the current run, so the alpha repo's dependent issue can unblock and consume it.

**Why this priority**: It is the only deliverable. Alpha is blocked until it is merged.

**Independent Test**: Run the greeting package's tests on the default branch after merge. The new test passes, and a text search of the default branch finds the function name and the string literal.

**Acceptance Scenarios**:

1. **Given** the default branch with earlier `HelloE2E2026…` entries, **When** this change is merged, **Then** `HelloE2E202610192019235408()` exists in the greeting package and returns `e2e-cross-repo-spawn-20261019-201923-5408`.
2. **Given** the merged change, **When** the greeting package tests run, **Then** the new test passes alongside every existing test.
3. **Given** the merged change, **When** the default branch is searched for `HelloE2E202610192019235408` and `e2e-cross-repo-spawn-20261019-201923-5408`, **Then** both are found verbatim.

---

### Edge Cases

- Earlier sentinel functions must stay intact and in their original order; the new entry goes after the latest existing `HelloE2E2026…` function and test.
- The branch must be based on a fresh fetch of the default branch, so concurrent sentinel additions are not lost or reordered.
- An open but unmerged PR does not satisfy the goal, since alpha stays blocked until the change reaches the default branch.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The greeting package MUST export a function named exactly `HelloE2E202610192019235408` that takes no arguments and returns a string.
- **FR-002**: `HelloE2E202610192019235408()` MUST return exactly `e2e-cross-repo-spawn-20261019-201923-5408`.
- **FR-003**: The new function MUST have a doc comment and follow the style of the existing sentinel functions (one function per run).
- **FR-004**: The greeting package's test file MUST include a test for the new function, in the existing style. The test declares `const want = "e2e-cross-repo-spawn-20261019-201923-5408"`, calls `HelloE2E202610192019235408()`, and reports a mismatch with `t.Errorf`. It uses only the stdlib `testing` package.
- **FR-005**: The function name and string literal MUST appear exactly as given, with no renaming, normalisation or "improvement".
- **FR-006**: The new function and test MUST be placed after the latest existing `HelloE2E2026…` entries, without changing, renaming, removing or reordering any existing function or test.
- **FR-007**: The change MUST be based on a fresh fetch of the default branch and MUST modify only the greeting source file and its test file.
- **FR-008**: The change MUST be merged to the default branch (`main`).

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `go test ./...` passes on the default branch after merge, including the new test.
- **SC-002**: A text search of the default branch finds both `HelloE2E202610192019235408` and `e2e-cross-repo-spawn-20261019-201923-5408` verbatim.
- **SC-003**: The diff touches exactly two files, the greeting source and its test, and consists only of additions.
- **SC-004**: Alpha issue #9131 is unblocked once this issue is closed.

## Assumptions

- The existing greeting package already has `HelloE2E2026…` sentinel functions and tests, and this change follows their pattern.
- The default branch is `main`.
- Merging is handled by the normal pipeline (the issue carries `fabrik:yolo`), so no separate merge instructions are needed.

## Out of Scope *(optional)*

- Changes to any existing greeting function or test.
- Changes to any file other than the greeting source and its test.
- Any work in the alpha repo, including the call to this function.

## Source References *(optional)*

- Parent issue: handarbeit/fabrik-test-alpha#9131 (this issue was spawned from it as part of a multi-issue decomposition)
- Cross-repo spawn issues #803 and #1308
