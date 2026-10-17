# Feature Specification: Add HelloE2E202610171410432373 greeting function with test

**Feature Branch**: `fabrik/issue-541`
**Created**: 2026-10-17
**Status**: Specified
**Input**: User description:

> ## Problem
> Another repo (`handarbeit/fabrik-test-alpha`, issue #8514) needs to call a new exported greeting function from this repo's `pkg/greeting` package. The function must exist on this repo's default branch first. This is part of an e2e cross-repo spawn regression scenario that checks the exact text lands here.
>
> ## Requirements
> - In `pkg/greeting/greeting.go`, add an exported function `HelloE2E202610171410432373() string` that returns exactly the string `e2e-cross-repo-spawn-20261017-141043-2373`.
> - In `pkg/greeting/greeting_test.go`, add a test `TestHelloE2E202610171410432373` that asserts the function returns exactly `e2e-cross-repo-spawn-20261017-141043-2373`. Follow the existing test style (a `const want` and a `got != want` check).
> - Use the function name and the string literal verbatim. Do not rename, shorten, "correct" or otherwise alter them. A per-run-unique token is already baked into both, and the scenario's final assertion checks this exact text in this repo.
>
> ## Scope
> Only the function and its test. No other changes to this repo.
>
> ## Acceptance
> - `go test ./...` passes, including the new test.
> - The default branch contains `HelloE2E202610171410432373` returning `e2e-cross-repo-spawn-20261017-141043-2373`.
> - Alpha can then import the function and bump its `go.mod` pin to the resulting beta commit.

## Background

A cross-repo spawn regression scenario is driven from `handarbeit/fabrik-test-alpha` (parent issue #8514). The parent's plan decomposes the work so that this repo (`fabrik-test-beta`) must first expose a new exported greeting function in its `pkg/greeting` package. Alpha then imports that function and bumps its `go.mod` pin to the resulting beta commit. The scenario's final assertion checks that the exact function name and sentinel text exist in this repo, and the name and string carry a per-run-unique token, so they must be reproduced exactly. The package already holds many earlier sentinel functions of the same shape, each with a matching test. This issue adds one more.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Expose the new sentinel greeting function (Priority: P1)

As the cross-repo spawn scenario running from alpha, I need a new exported function in beta's `pkg/greeting` package that returns a known sentinel string. Alpha can then import it and the scenario can verify the text landed in this repo.

**Why this priority**: It is the whole deliverable. Alpha's dependent work cannot proceed until the function exists on beta's default branch.

**Independent Test**: Call `HelloE2E202610171410432373()` from the `greeting` package and compare the result to `e2e-cross-repo-spawn-20261017-141043-2373`.

**Acceptance Scenarios**:

1. **Given** the default branch of this repo, **When** `HelloE2E202610171410432373()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261017-141043-2373`.
2. **Given** alpha imports `pkg/greeting` at the resulting beta commit, **When** it references `HelloE2E202610171410432373`, **Then** the symbol resolves as an exported function returning a string.

---

### User Story 2 - Guard the sentinel with a test (Priority: P1)

As a maintainer, I need a test that pins the function's exact return value, so that any accidental change to the sentinel text fails the suite.

**Why this priority**: The acceptance criteria require the suite to pass with the new test, and the test is the regression guard for the exact text.

**Independent Test**: Run `go test ./...` and confirm `TestHelloE2E202610171410432373` runs and passes.

**Acceptance Scenarios**:

1. **Given** the new function and test, **When** `go test ./...` runs, **Then** all tests pass, including `TestHelloE2E202610171410432373`.
2. **Given** the function returned any other string, **When** the test runs, **Then** it fails and reports the actual and expected values.

---

### Edge Cases

- The function name and string literal must not be normalised, shortened or "corrected", even where they look unusual or resemble earlier sentinels. The embedded token is intentionally unique to this run.
- Existing sentinel functions and tests must remain unchanged. In particular, the nearest earlier one (`HelloE2E202610150948432786`) must not be altered or reordered.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The `greeting` package MUST export a function `HelloE2E202610171410432373() string` that returns exactly `e2e-cross-repo-spawn-20261017-141043-2373`.
- **FR-002**: The package's tests MUST include `TestHelloE2E202610171410432373`, which asserts the function returns exactly `e2e-cross-repo-spawn-20261017-141043-2373`.
- **FR-003**: The new test MUST follow the existing test style: a `const want` holding the expected string and a `got != want` comparison that reports a failure.
- **FR-004**: The function name and string literal MUST be used verbatim, with no renaming, shortening or alteration.
- **FR-005**: The change MUST be limited to the function and its test, with no other changes to this repo.
- **FR-006**: The function MUST be present on the repo's default branch (`main`) so that alpha can import it and pin to the resulting commit.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `go test ./...` passes with zero failures, including the new test.
- **SC-002**: The default branch contains `HelloE2E202610171410432373` returning `e2e-cross-repo-spawn-20261017-141043-2373`, and the scenario's final assertion on that exact text succeeds.
- **SC-003**: Alpha can import the function and bump its `go.mod` pin to the resulting beta commit without further changes to this repo.

## Assumptions

- The new function and test follow the same conventions as the existing sentinel functions in the package (doc comment, plain string return). The doc comment wording is not significant.
- The new function and test are added alongside the existing ones, and no existing code is modified.
- The task is a single self-contained change, so no follow-up issues in this repo are needed.

## Out of Scope *(optional)*

- Any change to existing greeting functions or tests.
- Any change outside the function and its test (docs, configuration, other packages).
- Alpha's import and `go.mod` bump, which are handled in `handarbeit/fabrik-test-alpha` under issue #8514.

## Source References *(optional)*

- Parent issue: `handarbeit/fabrik-test-alpha#8514` (the multi-issue decomposition this was spawned from)
- Existing pattern: the sentinel functions and tests in the `pkg/greeting` package
