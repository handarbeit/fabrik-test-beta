# Feature Specification: Add HelloE2E202610261612416366 greeting function

**Feature Branch**: `fabrik/issue-649`
**Created**: 2026-10-26
**Status**: Specified
**Input**: User description:

> This is the beta-repo half of a cross-repo e2e regression scenario. It was spawned from `handarbeit/fabrik-test-alpha#10082`, which regression-tests on-demand spawn-target repo initialization (#803) and the per-run-unique fixture token (#1308). The alpha repo's `main.go` will call a new function from this repo's `pkg/greeting` package. The alpha repo cannot compile until this function exists on this repo's default branch, so this change must land first.
>
> - **FR-001**: `pkg/greeting/greeting.go` MUST export a function `HelloE2E202610261612416366() string`.
> - **FR-002**: That function MUST return exactly the string literal `"e2e-cross-repo-spawn-20261026-161241-6366"`.
> - **FR-003**: `pkg/greeting/greeting_test.go` MUST contain a test for that function, and it MUST pass. Follow the style of the existing `HelloE2E…` tests in that file.
> - **FR-004**: The function name and string literal MUST be used exactly as given above. They MUST NOT be altered, regenerated or replaced with other values. A per-run-unique token is already baked into both, and the scenario's final assertion depends on this exact text landing on this repo's default branch.
>
> Scope: Only the single new function and its test. No other changes to the greeting package or elsewhere in the repo. Follow the pattern of the existing `HelloE2E…` functions in `pkg/greeting/greeting.go`.
>
> Acceptance: `pkg/greeting/greeting.go` exports `HelloE2E202610261612416366() string` returning `"e2e-cross-repo-spawn-20261026-161241-6366"`. A test in `pkg/greeting/greeting_test.go` covers it and passes (`go test ./pkg/greeting/...`). After merge, this repo's default branch contains the text `e2e-cross-repo-spawn-20261026-161241-6366` and the function name `HelloE2E202610261612416366`.
>
> *Spawned by Fabrik from parent issue handarbeit/fabrik-test-alpha#10082 as a multi-issue decomposition.*

## Background

This issue is the beta-repo half of a cross-repo e2e regression scenario spawned from `handarbeit/fabrik-test-alpha#10082`. That scenario regression-tests on-demand initialization of a spawn-target repository (#803) and the per-run-unique fixture token (#1308). The alpha repo's `main.go` will call a new function in this repo's `pkg/greeting` package, so alpha cannot compile until the function is on this repo's default branch. This change must therefore land first. The package already holds several sentinel functions of the same shape, each with a matching test, from earlier scenario runs.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Sentinel function available to the alpha repo (Priority: P1)

The cross-repo scenario's alpha half needs a greeting function carrying this run's unique token. It must exist, unchanged, on this repo's default branch.

**Why this priority**: It is the only deliverable. The alpha repo cannot compile and the scenario's final assertion cannot pass without it.

**Independent Test**: Run `go test ./pkg/greeting/...`. The new test passes, and the default branch contains the function name and the string literal.

**Acceptance Scenarios**:

1. **Given** the change is merged, **When** `HelloE2E202610261612416366()` is called, **Then** it returns exactly `"e2e-cross-repo-spawn-20261026-161241-6366"`.
2. **Given** the change is merged, **When** `go test ./pkg/greeting/...` runs, **Then** all tests pass, including the new one for this function.
3. **Given** the change is merged, **When** the default branch is searched, **Then** it contains both `HelloE2E202610261612416366` and `e2e-cross-repo-spawn-20261026-161241-6366`.

---

### Edge Cases

- The name and literal carry a per-run-unique token. A typo, regeneration or normalisation (for example changing the hyphenation of the literal) breaks the scenario's final assertion.
- Existing sentinel functions and tests must remain untouched, so the diff adds only the new function and its test.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `pkg/greeting/greeting.go` MUST export a function `HelloE2E202610261612416366() string`.
- **FR-002**: That function MUST return exactly the string literal `"e2e-cross-repo-spawn-20261026-161241-6366"`.
- **FR-003**: `pkg/greeting/greeting_test.go` MUST contain a passing test for that function, in the style of the existing `HelloE2E…` tests.
- **FR-004**: The function name and string literal MUST be used exactly as given, and MUST NOT be altered, regenerated or replaced.
- **FR-005**: The change MUST add only this function and its test, with no other changes to the greeting package or the rest of the repo.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `go test ./pkg/greeting/...` passes, including a test that asserts the exact return value.
- **SC-002**: After merge, the default branch contains the function name `HelloE2E202610261612416366` and the text `e2e-cross-repo-spawn-20261026-161241-6366`.
- **SC-003**: The diff touches only `pkg/greeting/greeting.go` and `pkg/greeting/greeting_test.go`, and adds no lines other than the new function and test, apart from this spec's own file.

## Assumptions

- The new function and test are appended after the most recent existing `HelloE2E…` entries, following their doc-comment and test style.
- The default branch is `main`.

## Out of Scope *(optional)*

- Any change to existing greeting functions or tests.
- Any change to the alpha repo, which is handled by its own issue.
- Any change outside `pkg/greeting`.
