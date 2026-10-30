# Feature Specification: Add HelloE2E202610300724254399 sentinel greeting function

**Feature Branch**: `fabrik/issue-651`
**Created**: 2026-10-30
**Status**: Specified
**Input**: User description: "Add HelloE2E202610300724254399 to pkg/greeting (e2e cross-repo spawn). The alpha repo (`handarbeit/fabrik-test-alpha`) needs to call a new greeting function from this repo. Alpha's build fails until this repo exposes that function, so this change must land first. This is an e2e regression scenario for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token). The function name and string literal below contain a per-run-unique token and MUST be used exactly as written. Do not rename them, reformat them or invent a different token. The scenario's final assertion looks for this exact text in this repo. Requirements: In `pkg/greeting/greeting.go`, add an exported function named exactly `HelloE2E202610300724254399` with signature `() string`, returning exactly the literal `e2e-cross-repo-spawn-20261030-072425-4399`. Add a doc comment above it in the same style as the existing `HelloE2E…` functions, for example: `// HelloE2E202610300724254399 returns the sentinel string for the cross-repo spawn regression test.` In `pkg/greeting/greeting_test.go`, add `TestHelloE2E202610300724254399`, following the existing per-function test pattern: a `const want = "e2e-cross-repo-spawn-20261030-072425-4399"` and a `got != want` check that reports failure with `t.Errorf`. Keep all existing functions and tests unchanged. This change only adds to them. Scope: Only `pkg/greeting/greeting.go` and `pkg/greeting/greeting_test.go` change. No other greeting functionality, no dependency changes and no refactoring. Acceptance: `HelloE2E202610300724254399()` returns exactly `e2e-cross-repo-spawn-20261030-072425-4399`. `go build ./...` and `go test ./...` pass in this repo, including the new test. The change is merged to this repo's default branch, so the alpha repo can pin to a commit that contains the function."

## Background

This issue was spawned from parent issue handarbeit/fabrik-test-alpha#10162 as part of a multi-issue decomposition. The alpha repository (`handarbeit/fabrik-test-alpha`) must call a new greeting function that does not yet exist in this repository, so alpha's build fails until this repository provides it. This change must therefore land first, and alpha then pins to a commit of this repository that contains the function.

The change is a fixture for an end-to-end regression scenario covering two behaviours: on-demand initialisation of a spawn-target repository (#803) and per-run-unique fixture tokens (#1308). The function name and returned string embed a token unique to this run. The scenario's final assertion searches this repository for that exact text. Any deviation (rename, reformatting, or a different token) breaks the scenario.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Alpha repo can consume the sentinel greeting function (Priority: P1)

A developer or automated pipeline working in the alpha repository needs to call `HelloE2E202610300724254399` from this repository's greeting package. After this change merges to the default branch, alpha can pin to a commit containing the function and build successfully.

**Why this priority**: This is the entire purpose of the issue. Alpha's build is blocked until the function exists here.

**Independent Test**: Call `HelloE2E202610300724254399()` and compare the result with `e2e-cross-repo-spawn-20261030-072425-4399`. This is covered by the new unit test, with no dependency on the alpha repo.

**Acceptance Scenarios**:

1. **Given** the change is merged to this repo's default branch, **When** `HelloE2E202610300724254399()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261030-072425-4399`.
2. **Given** the change is merged, **When** `go build ./...` and `go test ./...` run in this repo, **Then** both succeed, including the new test.

---

### User Story 2 - The e2e scenario's final assertion finds the exact token (Priority: P1)

The cross-repo spawn regression scenario ends by checking this repository for the per-run-unique function name and string literal.

**Why this priority**: A wrong or reformatted token fails the regression scenario even if the code compiles.

**Independent Test**: Search the repository for the strings `HelloE2E202610300724254399` and `e2e-cross-repo-spawn-20261030-072425-4399` and confirm each is present, spelled exactly as given.

**Acceptance Scenarios**:

1. **Given** the change is merged, **When** the repository is searched for `HelloE2E202610300724254399`, **Then** the exported function and its test are found.
2. **Given** the change is merged, **When** the repository is searched for `e2e-cross-repo-spawn-20261030-072425-4399`, **Then** the literal is found in both the function and the test.

---

### Edge Cases

- Existing `HelloE2E…` functions and their tests must be unaffected, with no signature, behaviour or ordering change to them.
- A per-run-unique token that is mistyped, case-altered or split across lines would pass compilation but fail the scenario's assertion. The token must appear exactly as given.
- The new test must fail if the function returns anything other than the expected literal.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The greeting package MUST expose an exported function named exactly `HelloE2E202610300724254399` with signature `() string`.
- **FR-002**: `HelloE2E202610300724254399()` MUST return exactly the literal `e2e-cross-repo-spawn-20261030-072425-4399`.
- **FR-003**: The function MUST have a doc comment in the same style as the existing `HelloE2E…` functions, for example `// HelloE2E202610300724254399 returns the sentinel string for the cross-repo spawn regression test.`
- **FR-004**: The greeting package's tests MUST include `TestHelloE2E202610300724254399`, following the existing per-function test pattern: a `const want = "e2e-cross-repo-spawn-20261030-072425-4399"` and a `got != want` comparison that reports failure via `t.Errorf`.
- **FR-005**: All existing greeting functions and tests MUST remain unchanged. This change only adds to them.
- **FR-006**: Only `pkg/greeting/greeting.go` and `pkg/greeting/greeting_test.go` MAY change.
- **FR-007**: The function name and string literal MUST be used exactly as specified, with no renaming, reformatting or substitute token.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `HelloE2E202610300724254399()` returns a value byte-for-byte equal to `e2e-cross-repo-spawn-20261030-072425-4399`.
- **SC-002**: `go build ./...` and `go test ./...` pass in this repo, including the new test.
- **SC-003**: A repository search finds the exact function name and exact literal, matching the e2e scenario's final assertion.
- **SC-004**: The diff touches only the two named files and removes or modifies no existing lines in them.
- **SC-005**: The change is merged to the default branch, so the alpha repo can pin to a commit that contains the function.

## Assumptions

- The existing `HelloE2E…` functions and tests in the greeting package establish the style for doc comments and tests, and the new additions follow it.
- The default branch is `main`, and merging to it is the handoff point for the alpha repo.
- The token is final and supplied by the scenario. It is not regenerated or validated here.

## Out of Scope *(optional)*

- Any other greeting functionality.
- Dependency changes.
- Refactoring existing code or tests.
- Changes to the alpha repo, which consumes this function in a separate issue.
- Changes to files other than the greeting source file and its test file.

## Source References *(optional)*

- Parent issue: handarbeit/fabrik-test-alpha#10162 (multi-issue decomposition that spawned this issue)
- e2e regression scenario context: #803 (on-demand spawn-target repo init), #1308 (per-run-unique fixture token)
