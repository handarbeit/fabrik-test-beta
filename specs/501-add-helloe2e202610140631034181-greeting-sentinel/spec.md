# Feature Specification: Add HelloE2E202610140631034181 greeting sentinel (e2e cross-repo spawn, #8119)

**Feature Branch**: `fabrik/issue-501`
**Created**: 2026-10-14
**Status**: Specified
**Input**: User description:

> ## Problem
> The alpha repo (`handarbeit/fabrik-test-alpha`, issue #8119) needs to call a new exported function from this repo's `pkg/greeting` package. That function does not exist yet, so it must land here first. This is an end-to-end regression scenario for on-demand spawn-target repo init (#803) and the per-run-unique fixture token (#1308).
>
> ## Requirements
> - Add an exported function `HelloE2E202610140631034181() string` to `pkg/greeting/greeting.go`. It must return exactly the string literal `e2e-cross-repo-spawn-20261014-063103-4181`. Give it a one-line doc comment, matching the existing sentinel functions in that file (for example `HelloIssue3737`).
> - Add a test `TestHelloE2E202610140631034181` to `pkg/greeting/greeting_test.go`. Use the same shape as the existing sentinel tests: a `const want = "e2e-cross-repo-spawn-20261014-063103-4181"` and a `got != want` check.
> - Use the function name and the string literal **verbatim**. Do not rename, normalise or regenerate them. A per-run-unique token is already embedded, and the scenario's final assertion checks this exact text in this repo's default branch.
>
> ## Scope
> - Only the two files above change. Do not add any other greeting functionality, and do not modify other functions or tests.
> - If `pkg/greeting` does not exist, create it with the function and test; the existing sentinels indicate it does.
>
> ## Acceptance
> - `go test ./pkg/greeting/...` passes, including the new test.
> - `HelloE2E202610140631034181()` returns exactly `e2e-cross-repo-spawn-20261014-063103-4181`.
> - The change is merged to this repo's default branch, so the alpha repo can pin to the merge commit.
>
> ## Context
> This issue exists to unblock handarbeit/fabrik-test-alpha#8119, which will bump its `go.mod` pin to the merge commit of this change and then call the new function from its `main.go`. Mention #8119 in the PR body by bare number only; do NOT use a closing keyword (`Closes`/`Fixes`/`Resolves`) near it, since it lives in another repo and must close on its own.
>
> *Spawned by Fabrik from parent issue handarbeit/fabrik-test-alpha#8119 as a multi-issue decomposition. The parent's plan is at the link above.*

## Background

The alpha repo (`handarbeit/fabrik-test-alpha`) issue #8119 needs to call a new exported function from this repo's `pkg/greeting` package. The function does not exist yet, so it has to land here first. Alpha will then bump its `go.mod` pin to the merge commit of this change and call the function from its `main.go`.

This is an end-to-end regression scenario for two behaviours:

- on-demand spawn-target repo init (#803)
- the per-run-unique fixture token (#1308)

The scenario's final assertion checks that this exact sentinel text exists in this repo's default branch, so the name and literal must not be altered in any way.

The `pkg/greeting` package already holds several equivalent sentinel functions (for example `HelloIssue3737`, which returns `e2e-cross-repo-spawn-3737`), each with a one-line doc comment and a matching test. This change adds one more of the same kind.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Alpha can consume the new sentinel (Priority: P1)

The alpha repo's issue #8119 pins this repo at the merge commit of this change and calls the new sentinel function. The function therefore has to exist on this repo's default branch, with the exact name and return value the scenario expects.

**Why this priority**: It is the whole purpose of the issue. Without it the cross-repo scenario cannot complete.

**Independent Test**: Run the greeting package tests and confirm that the new test exists and passes. Then confirm that the default branch contains the exact function name and literal after the merge.

**Acceptance Scenarios**:

1. **Given** the greeting package on the default branch, **When** `HelloE2E202610140631034181()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261014-063103-4181`.
2. **Given** the new test `TestHelloE2E202610140631034181`, **When** the greeting package tests run, **Then** the test passes together with all existing tests.
3. **Given** the change is merged, **When** the alpha repo pins to the merge commit, **Then** the sentinel function is available to it.

---

### Edge Cases

- The name and the literal must be reproduced character for character. Any normalisation, regeneration or reformatting of the per-run-unique token breaks the scenario's final assertion.
- If the greeting package were missing, it would have to be created along with the function and test. The existing sentinels show it is present.
- The PR body must mention #8119 by bare number only. A closing keyword near it would try to close an issue in another repo, which must close on its own.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The greeting package exposes an exported function `HelloE2E202610140631034181() string` that returns exactly `e2e-cross-repo-spawn-20261014-063103-4181`.
- **FR-002**: The new function carries a one-line doc comment, matching the style of the existing sentinel functions (for example `HelloIssue3737`).
- **FR-003**: The greeting package tests include `TestHelloE2E202610140631034181`. It uses the same shape as the existing sentinel tests: a `const want = "e2e-cross-repo-spawn-20261014-063103-4181"` and a `got != want` check.
- **FR-004**: The function name and string literal are used verbatim, with no renaming, normalisation or regeneration.
- **FR-005**: Only the greeting package's source file and its test file change. No other greeting functionality is added, and no other functions or tests are modified.
- **FR-006**: The PR body mentions #8119 by bare number only, with no closing keyword (`Closes`/`Fixes`/`Resolves`) near it.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The greeting package tests pass, including the new test.
- **SC-002**: `HelloE2E202610140631034181()` returns a string exactly equal to `e2e-cross-repo-spawn-20261014-063103-4181`.
- **SC-003**: After merge, the exact function name and literal are present on this repo's default branch, so the scenario's final assertion succeeds.
- **SC-004**: The diff touches exactly two files, and existing functions and tests are unchanged.

## Assumptions

- The greeting package already exists with the sentinel functions and tests described, so nothing needs creating.
- The function and test follow the existing sentinel pattern in that package: a one-line doc comment, a constant return, and a `got != want` test.
- The alpha repo handles its own `go.mod` bump and call site. That is separate work tracked in alpha issue #8119.

## Out of Scope *(optional)*

- Any other greeting functionality, and any change to existing functions or tests.
- The alpha repo's `go.mod` pin bump and its `main.go` call (alpha issue #8119).
- Changes to the cross-repo spawn or fixture-token mechanisms themselves (#803, #1308).

## Source References *(optional)*

- handarbeit/fabrik-test-alpha#8119 (parent issue that consumes this sentinel)
- Existing sentinel pattern: `HelloIssue3737` in the greeting package
