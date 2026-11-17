# Feature Specification: Add HelloE2E202611170345251556 to pkg/greeting (e2e cross-repo spawn)

**Feature Branch**: `fabrik/issue-756`
**Created**: 2026-11-16
**Status**: Specified
**Input**: User description: the original issue body, reproduced verbatim below.

> ## Problem
>
> The alpha repo (`handarbeit/fabrik-test-alpha`) must call a new greeting function from this repo, but the function does not exist yet. This repo must gain it first, because alpha cannot import a symbol that is not present. This is an end-to-end regression scenario for cross-repo decomposition (#803 on-demand spawn-target repo init, #1308 per-run-unique fixture token).
>
> ## Requirements
>
> - **FR-001**: `pkg/greeting/greeting.go` MUST export a function `HelloE2E202611170345251556() string` that returns exactly the string literal `"e2e-cross-repo-spawn-20261117-034525-1556"`. Follow the style of the existing `HelloE2E…` functions in that file.
> - **FR-002**: `pkg/greeting/greeting_test.go` MUST contain a test that calls `HelloE2E202611170345251556()` and asserts it returns exactly `"e2e-cross-repo-spawn-20261117-034525-1556"`. Follow the style of the existing tests in that file.
> - **FR-007**: Use the function name and string literal above verbatim. Do not invent, rename, or "improve" them. The token `20261117-034525-1556` is a per-run-unique value, and the scenario's final assertion checks that this exact text appears in `pkg/greeting/greeting.go` in this repo.
>
> ## Scope
>
> - Only the two files named above. Make no other behaviour changes to the greeting package.
> - The change must be merged into this repo's default branch, because the alpha repo pins to the resulting commit.
>
> ## Acceptance
>
> - `HelloE2E202611170345251556()` returns `"e2e-cross-repo-spawn-20261117-034525-1556"`.
> - `go test ./...` passes in this repo, including the new test.
> - The text `e2e-cross-repo-spawn-20261117-034525-1556` is present in `pkg/greeting/greeting.go`.
>
> *Spawned by Fabrik from parent issue handarbeit/fabrik-test-alpha#11468 as a multi-issue decomposition. The parent's plan is at the link above.*

## Background

The alpha repo (`handarbeit/fabrik-test-alpha`) must call a new greeting function that lives in this repo (`handarbeit/fabrik-test-beta`). Alpha cannot import a symbol that does not exist yet, so this repo has to gain the function first, and the change has to be merged into this repo's default branch because alpha pins to the resulting commit.

This issue is one half of an end-to-end regression scenario for cross-repo decomposition. The scenario covers on-demand initialisation of the spawn-target repo (#803) and a per-run-unique fixture token (#1308). The token `20261117-034525-1556` is unique to this run. The scenario's final assertion checks that the exact text `e2e-cross-repo-spawn-20261117-034525-1556` appears in `pkg/greeting/greeting.go` in this repo, so the function name and literal must be used exactly as given.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo exposes the per-run sentinel greeting (Priority: P1)

A downstream consumer, the alpha repo, needs to call `HelloE2E202611170345251556()` from this repo's `pkg/greeting` package and receive the run-unique string. A maintainer or the e2e scenario needs to confirm the function exists and returns the expected value.

**Why this priority**: It is the only deliverable. Without it the alpha repo cannot compile against this repo and the cross-repo scenario cannot complete.

**Independent Test**: Run `go test ./...` in this repo. The new test passes, and `pkg/greeting/greeting.go` contains the text `e2e-cross-repo-spawn-20261117-034525-1556`.

**Acceptance Scenarios**:

1. **Given** the change is merged to the default branch, **When** a caller invokes `HelloE2E202611170345251556()`, **Then** it returns exactly `"e2e-cross-repo-spawn-20261117-034525-1556"`.
2. **Given** the change is merged to the default branch, **When** the file `pkg/greeting/greeting.go` is searched for `e2e-cross-repo-spawn-20261117-034525-1556`, **Then** the text is present.
3. **Given** the change is applied, **When** `go test ./...` runs in this repo, **Then** all tests pass, including the new test for `HelloE2E202611170345251556()`.

---

### Edge Cases

- The function name and literal contain a per-run-unique token. Any typo, reformatting or "improvement" (a different case, a different separator, extra whitespace) breaks the scenario's final text assertion.
- Existing `HelloE2E…` functions and their tests must be left unchanged, so the diff is purely additive.
- The change must reach the default branch (`main`). A change left only on the issue branch does not satisfy the alpha repo's pin.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `pkg/greeting/greeting.go` MUST export a function `HelloE2E202611170345251556() string` that returns exactly the string literal `"e2e-cross-repo-spawn-20261117-034525-1556"`, following the style of the existing `HelloE2E…` functions in that file.
- **FR-002**: `pkg/greeting/greeting_test.go` MUST contain a test that calls `HelloE2E202611170345251556()` and asserts it returns exactly `"e2e-cross-repo-spawn-20261117-034525-1556"`, following the style of the existing tests in that file.
- **FR-007**: The function name and string literal in FR-001 and FR-002 MUST be used verbatim, with no renaming, reformatting or alteration.

### Key Entities *(if applicable)*

- **Sentinel greeting function**: A per-run-unique exported function in `pkg/greeting` that returns a fixed string. It exists so the e2e scenario can verify the cross-repo dependency was satisfied.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: Calling `HelloE2E202611170345251556()` returns exactly `"e2e-cross-repo-spawn-20261117-034525-1556"`.
- **SC-002**: `go test ./...` passes in this repo, including the new test.
- **SC-003**: The text `e2e-cross-repo-spawn-20261117-034525-1556` is present in `pkg/greeting/greeting.go` on this repo's default branch after merge.

## Assumptions

- Requirement identifiers FR-001, FR-002 and FR-007 are kept exactly as in the original issue, because the parent scenario may reference them. The gap at FR-003 to FR-006 is intentional.
- Existing `HelloE2E…` functions and tests in `pkg/greeting` define the style to follow, and the new function and test are additive.
- Merging to the default branch happens through the normal Fabrik pipeline (PR with `Closes #756`), with no extra manual step.

## Out of Scope

- Any file other than `pkg/greeting/greeting.go` and `pkg/greeting/greeting_test.go`.
- Any behaviour change to existing functions in the greeting package.
- The alpha repo's call into this function, which is handled by the parent issue (handarbeit/fabrik-test-alpha#11468) and its other spawned work.
- Any rename, refactor or "improvement" of the function name or literal.
