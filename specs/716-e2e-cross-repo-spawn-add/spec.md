# Feature Specification: Add HelloE2E202611071936261190 to pkg/greeting

**Feature Branch**: `fabrik/issue-716`
**Created**: 2026-11-07
**Status**: Specified
**Input**: User description: "
## Goal

Add one exported greeting function and a test for it to this repo. This is the beta-side half of a cross-repo change. The alpha repo (`handarbeit/fabrik-test-alpha`) will import and call this function once it has landed on `main`, so this issue must be merged first.

## What needs to change

- In `pkg/greeting/greeting.go`, add an exported function `HelloE2E202611071936261190() string` that returns exactly the literal string `e2e-cross-repo-spawn-20261107-193626-1190`.
- In `pkg/greeting/greeting_test.go`, add a test for `HelloE2E202611071936261190` that asserts it returns exactly `e2e-cross-repo-spawn-20261107-193626-1190`.
- Follow the pattern of the existing `HelloE2E…` functions and tests already in those files.
- If `pkg/greeting/greeting.go` or `pkg/greeting/greeting_test.go` does not exist yet, create the package so that the function and test above exist.

## Constraints

- Use the exact function name `HelloE2E202611071936261190` and the exact string literal `e2e-cross-repo-spawn-20261107-193626-1190`. Do not invent your own name or token. A per-run-unique token is already baked into both, and the scenario's final assertion depends on that exact text landing in this repo.
- Do not change any existing greeting functions or their tests.

## Acceptance

- `HelloE2E202611071936261190()` returns exactly `e2e-cross-repo-spawn-20261107-193626-1190`.
- The new test exists in `pkg/greeting/greeting_test.go`, and `go test ./...` passes.
- The change is merged to the default branch so that alpha can resolve it as a module dependency.

This issue is the beta-side half of handarbeit/fabrik-test-alpha issue #10794. It is an e2e regression scenario for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This issue is the beta-side half of a cross-repo end-to-end regression scenario. It was spawned from handarbeit/fabrik-test-alpha issue #10794, and it covers the on-demand initialisation of a spawn-target repo (#803) and the per-run-unique fixture token (#1308). The alpha repo will import and call a new sentinel function from this repo's `pkg/greeting` package, so the function has to be merged to the default branch first. The scenario's final assertion depends on the exact function name and string token landing in this repo. The package already holds several earlier sentinel functions of the same form, each with its own test.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Expose the sentinel function for alpha to consume (Priority: P1)

A maintainer or the e2e scenario needs `pkg/greeting` to export `HelloE2E202611071936261190`, returning the run-unique token, so that the alpha repo can resolve it as a module dependency once it is on `main`.

**Why this priority**: This is the whole deliverable. The alpha-side change is blocked until it lands.

**Independent Test**: Run `go test ./...` in this repo. The new test passes and asserts the exact token.

**Acceptance Scenarios**:

1. **Given** the change is merged to `main`, **When** `HelloE2E202611071936261190()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261107-193626-1190`.
2. **Given** the change is on the branch, **When** `go test ./...` runs, **Then** all tests pass, including the new `TestHelloE2E202611071936261190`.
3. **Given** the existing greeting functions and tests, **When** the change is applied, **Then** none of them are modified.

---

### Edge Cases

- If `pkg/greeting/greeting.go` or `greeting_test.go` is missing, the package, the function and the test must still be created. Both files exist today, so this is not expected.
- The function name and token carry a per-run-unique timestamp. They must be used exactly as given, with no normalisation or renaming.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `pkg/greeting` MUST export a function `HelloE2E202611071936261190() string` that returns exactly the literal `e2e-cross-repo-spawn-20261107-193626-1190`.
- **FR-002**: `pkg/greeting` MUST include a test for `HelloE2E202611071936261190` that asserts the exact return value above.
- **FR-003**: The new function and test MUST follow the conventions of the existing `HelloE2E…` functions and tests in the same files, including a doc comment and the same test shape.
- **FR-004**: Existing greeting functions and their tests MUST NOT be changed.
- **FR-005**: The PR for this issue MUST contain `Closes #716`. It MUST NOT use a closing keyword for the alpha-side parent issue.

### Key Entities *(if applicable)*

- **Sentinel function**: An exported, argument-less function in `pkg/greeting` that returns a fixed per-run token. The cross-repo e2e scenario uses it to prove that the beta-side change landed and is consumable.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: Calling `HelloE2E202611071936261190()` returns `e2e-cross-repo-spawn-20261107-193626-1190` with no character differing.
- **SC-002**: `go test ./...` passes with the new test included.
- **SC-003**: The change is merged to the default branch, so the alpha repo can resolve the function as a module dependency.
- **SC-004**: The diff touches no existing greeting function or test.

## Assumptions

- `pkg/greeting/greeting.go` and `greeting_test.go` already exist and follow the established `HelloE2E…` pattern. Each sentinel is a documented function returning a string literal, with a matching `Test…` function using a `want` constant.
- The Go module path and package layout need no change for alpha to import the function.
- The function name and token are final and were generated for this run, so they are not to be reconsidered.

## Out of Scope *(optional)*

- Any change to the alpha repo, including importing or calling the function. That is handled by the parent issue handarbeit/fabrik-test-alpha#10794.
- Modifying, renaming or removing existing greeting functions or tests.
- Any other changes to the repo, such as refactoring the greeting package, adding documentation or changing the module configuration.
