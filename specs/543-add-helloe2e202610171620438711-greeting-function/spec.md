# Feature Specification: Add HelloE2E202610171620438711 greeting function

**Feature Branch**: `fabrik/issue-543`
**Created**: 2026-10-17
**Status**: Specified
**Input**: User description: "The alpha repo (`handarbeit/fabrik-test-alpha`) needs to call a new exported function from this repo's `pkg/greeting` package. That function does not exist yet, and alpha cannot compile a call to it until it lands here. This issue is the beta half of an end-to-end cross-repo decomposition regression test (#803 on-demand spawn-target repo init, #1308 per-run-unique fixture token). The beta change must land first. Requirements: (1) In `pkg/greeting/greeting.go`, add an exported function with this exact name and signature: `func HelloE2E202610171620438711() string`. (2) It MUST return exactly this string literal, with no other characters: `e2e-cross-repo-spawn-20261017-162043-8711`. (3) In `pkg/greeting/greeting_test.go`, add a test for it, for example `TestHelloE2E202610171620438711`. The test MUST assert that the function returns exactly `e2e-cross-repo-spawn-20261017-162043-8711`. Follow the style of the existing `TestHelloE2E…` tests in that file, and a doc comment in the style of the existing `HelloE2E…` functions. Constraints: Use the function name `HelloE2E202610171620438711` and the string literal `e2e-cross-repo-spawn-20261017-162043-8711` verbatim. A per-run-unique token is already baked into both. Do not invent, regenerate, shorten or re-derive either one. The scenario's final assertion checks that this exact text exists in this repo's default branch. Do not modify or remove any existing function or test in the package. Only add. If `pkg/greeting/greeting.go` or `greeting_test.go` is missing, create it with the correct `package greeting` declaration (on-demand initialization, #803). Acceptance: `go test ./pkg/greeting/...` passes and includes the new test. `HelloE2E202610171620438711()` returns exactly `e2e-cross-repo-spawn-20261017-162043-8711`. The change is merged to this repo's default branch, so the alpha repo can pin its `go.mod` to the resulting commit."

## Background

The alpha repo (`handarbeit/fabrik-test-alpha`) must call a new exported function from this repo's `pkg/greeting` package, and cannot compile that call until the function exists here. This issue is the beta half of an end-to-end cross-repo decomposition regression test. It exercises on-demand initialization of a spawn-target repo (#803) and a per-run-unique fixture token (#1308). The beta change must land first so the alpha repo can pin its `go.mod` to the resulting commit. The scenario's final assertion checks that the exact function name and sentinel string exist on this repo's default branch, so both must be used verbatim.

The package already holds a series of equivalent sentinel functions and matching tests. Each returns a fixed `e2e-cross-repo-spawn-…` string, and each has a doc comment and a `TestHelloE2E…` test. This change follows that established pattern.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta exposes the new sentinel function (Priority: P1)

A maintainer of the cross-repo regression scenario needs `pkg/greeting` to export `HelloE2E202610171620438711`, returning the run's sentinel string, so the alpha repo can compile a call to it and pin to the commit that adds it.

**Why this priority**: It is the only deliverable. Alpha's half of the scenario is blocked until it lands.

**Independent Test**: Run `go test ./pkg/greeting/...`. The new test passes and asserts the exact return value.

**Acceptance Scenarios**:

1. **Given** the beta default branch after this change, **When** `HelloE2E202610171620438711()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261017-162043-8711`.
2. **Given** the change is merged, **When** the scenario searches the default branch for the function name and the sentinel string, **Then** both are found verbatim.
3. **Given** the existing sentinel functions and tests, **When** this change is applied, **Then** all of them are unchanged and their tests still pass.

---

### Edge Cases

- If `pkg/greeting/greeting.go` or `greeting_test.go` is missing, it is created with a `package greeting` declaration (on-demand initialization, #803). Both files currently exist, so this should not occur here.
- The sentinel must contain no leading or trailing whitespace or other characters.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `pkg/greeting` MUST export `func HelloE2E202610171620438711() string`, named exactly as shown.
- **FR-002**: `HelloE2E202610171620438711()` MUST return exactly the string `e2e-cross-repo-spawn-20261017-162043-8711`, with no other characters.
- **FR-003**: The function MUST carry a doc comment in the style of the existing `HelloE2E…` functions.
- **FR-004**: `pkg/greeting/greeting_test.go` MUST contain a test (e.g. `TestHelloE2E202610171620438711`) in the style of the existing `TestHelloE2E…` tests, asserting the function returns exactly `e2e-cross-repo-spawn-20261017-162043-8711`.
- **FR-005**: No existing function or test in the package may be modified or removed. The change is additive only.
- **FR-006**: If `greeting.go` or `greeting_test.go` is missing, it MUST be created with a `package greeting` declaration.
- **FR-007**: The function name and string literal MUST be used verbatim and never regenerated, shortened or re-derived.

### Key Entities *(if applicable)*

- **Sentinel function**: An exported zero-argument function in `pkg/greeting` that returns a fixed per-run token string, used as a cross-repo compile and presence check.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `go test ./pkg/greeting/...` passes and includes the new test.
- **SC-002**: The exact text `HelloE2E202610171620438711` and `e2e-cross-repo-spawn-20261017-162043-8711` both appear in this repo's default branch after merge.
- **SC-003**: The alpha repo can pin its `go.mod` to the merged commit and compile a call to the function.

## Assumptions

- Both `pkg/greeting/greeting.go` and `greeting_test.go` already exist, so the on-demand creation path is not exercised.
- The existing sentinel function and test pattern is the correct style to follow.

## Out of Scope *(optional)*

- Changes to the alpha repo, including pinning its `go.mod` (handled by the sibling issue).
- Modifying, refactoring or removing existing sentinel functions or tests.
- Any change outside `pkg/greeting`.
