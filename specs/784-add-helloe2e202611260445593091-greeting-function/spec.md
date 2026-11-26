# Feature Specification: Add HelloE2E202611260445593091 greeting function

**Feature Branch**: `fabrik/issue-784`
**Created**: 2026-11-25
**Status**: Specified
**Input**: User description:

> ## Problem
>
> The alpha repo (`handarbeit/fabrik-test-alpha`) is about to call a new exported function from this repo's `pkg/greeting` package. The function must exist and be released on this repo's default branch first, because alpha cannot compile a call to a function that does not yet exist. This issue is the beta half of an end-to-end regression scenario for cross-repo decomposition (#803 on-demand spawn-target repo init, #1308 per-run-unique fixture token). The scenario's final assertion depends on the exact text below landing in this repo.
>
> ## Requirements
>
> 1. In `pkg/greeting/greeting.go`, add an exported function with exactly this name and signature, with a doc comment in the same style as the existing `HelloE2E…` functions:
>
>    `func HelloE2E202611260445593091() string`
>
>    It takes no arguments and MUST return exactly the string literal:
>
>    `e2e-cross-repo-spawn-20261126-044559-3091`
>
> 2. In `pkg/greeting/greeting_test.go`, add a test `TestHelloE2E202611260445593091` following the existing pattern for the other `HelloE2E…` tests (a `const want = "e2e-cross-repo-spawn-20261126-044559-3091"`, call the function, and `t.Errorf` if `got != want`). The test must pass.
>
> 3. Use the function name and the string literal verbatim. Do not invent a different token, rename the function or alter the literal in any way; the scenario's final assertion checks for this exact text.
>
> ## Scope
>
> - Only the new function and its test. Make no other change to this repo.
> - Do not change existing functions or tests.
>
> ## Acceptance Criteria
>
> - `HelloE2E202611260445593091()` returns exactly `e2e-cross-repo-spawn-20261126-044559-3091`.
> - `go test ./...` passes in this repo, including the new test.
> - The change is merged to the default branch, so that the alpha repo can depend on a commit containing it.
>
> ## Context
>
> Parent issue: handarbeit/fabrik-test-alpha#11794. The alpha repo will bump its pinned dependency on this module to a commit containing this function, call `greeting.HelloE2E202611260445593091()` from its `main.go` and print the result. Do not make any change on the alpha side from this issue.
>
> *Spawned by Fabrik from parent issue handarbeit/fabrik-test-alpha#11794 as a multi-issue decomposition. The parent's plan is at the link above.*

## Background

The alpha repo (`handarbeit/fabrik-test-alpha`) will call a new exported function from this repo's `pkg/greeting` package. Alpha cannot compile a call to a function that does not yet exist, so the function has to be present on this repo's default branch first, and alpha then pins a commit that contains it.

This issue is the beta half of an end-to-end regression scenario for cross-repo decomposition (#803 on-demand spawn-target repo init, #1308 per-run-unique fixture token). The scenario's final assertion checks for the exact function name and string literal below, so they must be reproduced character for character.

Parent issue: handarbeit/fabrik-test-alpha#11794. Alpha will bump its pinned dependency on this module to a commit containing the function, call `greeting.HelloE2E202611260445593091()` from its `main.go` and print the result. No alpha-side change belongs in this issue.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Consumer can call the new greeting function (Priority: P1)

A consumer of this module (the alpha repo) imports `pkg/greeting` and calls `HelloE2E202611260445593091()`. It receives the fixed token string, so the cross-repo scenario's final assertion can find that text in alpha's output.

**Why this priority**: It is the only deliverable of this issue, and the cross-repo scenario cannot complete without it.

**Independent Test**: Call `HelloE2E202611260445593091()` and compare the result with `e2e-cross-repo-spawn-20261126-044559-3091`. This needs no alpha-side code.

**Acceptance Scenarios**:

1. **Given** the change is on this repo's default branch, **When** a caller invokes `HelloE2E202611260445593091()`, **Then** it returns exactly `e2e-cross-repo-spawn-20261126-044559-3091`.
2. **Given** the change is on this repo's default branch, **When** `go test ./...` runs, **Then** all tests pass, including `TestHelloE2E202611260445593091`.
3. **Given** the change is merged, **When** the alpha repo pins a commit of this module that contains it, **Then** alpha can compile a call to the function.

---

### Edge Cases

- A single character differing from the specified function name or literal (a transposed digit, a different timestamp, added whitespace or a trailing newline) breaks the scenario's final assertion and counts as a failure.
- The existing `HelloE2E…` functions and their tests must keep working unchanged.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The `pkg/greeting` package exposes an exported function named exactly `HelloE2E202611260445593091` that takes no arguments and returns a `string`.
- **FR-002**: `HelloE2E202611260445593091()` returns exactly the literal `e2e-cross-repo-spawn-20261126-044559-3091`, with no extra characters.
- **FR-003**: The function has a doc comment in the same style as the existing `HelloE2E…` functions in the package.
- **FR-004**: A test named `TestHelloE2E202611260445593091` follows the existing `HelloE2E…` test pattern: a `const want = "e2e-cross-repo-spawn-20261126-044559-3091"`, a call to the function, and a `t.Errorf` when `got != want`. The test passes.
- **FR-005**: No existing function, test or other file in this repo is changed beyond adding the new function and its test.
- **FR-006**: The change reaches this repo's default branch (`main`), so a commit containing it exists for the alpha repo to depend on.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `HelloE2E202611260445593091()` returns a string byte-identical to `e2e-cross-repo-spawn-20261126-044559-3091`.
- **SC-002**: `go test ./...` passes in this repo, including the new test, with no previously passing test failing.
- **SC-003**: The default branch contains a commit with the function, so the alpha repo can pin that commit and the cross-repo scenario's final assertion finds the exact text.

## Assumptions

- The existing `HelloE2E…` functions and tests in `pkg/greeting` establish the doc-comment and test conventions to follow; no new convention is introduced.
- The timestamp-like token is a per-run fixture and is used exactly as given, not validated or regenerated.
- Making the change available to alpha means merging to the default branch. Cutting a tagged release is not required, since alpha pins a commit.

## Out of Scope *(optional)*

- Any change on the alpha side (bumping the dependency, editing `main.go`, printing the result).
- Changes to or removal of existing greeting functions or tests.
- Any other change in this repo, including docs, CI and unrelated refactors.
- Tagging or publishing a release of this module.
