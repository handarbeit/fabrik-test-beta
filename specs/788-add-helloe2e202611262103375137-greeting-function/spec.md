# Feature Specification: Add HelloE2E202611262103375137 greeting function (e2e cross-repo spawn)

**Feature Branch**: `fabrik/issue-788`
**Created**: 2026-11-26
**Status**: Specified
**Input**: User description: "## Problem

The alpha repository (`handarbeit/fabrik-test-alpha`) needs to call a new exported greeting function from this repository's `pkg/greeting` package. That function does not exist yet, so it must land here before alpha can import it. This is part of an end-to-end regression scenario for cross-repo decomposition (#803 on-demand spawn-target repo init, #1308 per-run-unique fixture token).

## Requirements

1. In `pkg/greeting/greeting.go`, add an exported function with this exact name and signature:

   `func HelloE2E202611262103375137() string`

   It must return exactly this literal string (no extra whitespace or newline):

   `e2e-cross-repo-spawn-20261126-210337-5137`

   Follow the existing style of the other `HelloE2E<timestamp><suffix>()` functions in that file, including a one-line doc comment, for example:
   `// HelloE2E202611262103375137 returns the sentinel string for the cross-repo spawn regression test.`

2. In `pkg/greeting/greeting_test.go`, add a test named `TestHelloE2E202611262103375137`. Follow the existing tests' shape: declare `const want = \"e2e-cross-repo-spawn-20261126-210337-5137\"`, call `HelloE2E202611262103375137()`, and report a mismatch with `t.Errorf`.

3. Do not change any other function, test or file. If `go mod tidy` is not needed, do not run it.

## Constraints

- Use the function name `HelloE2E202611262103375137` and the literal `e2e-cross-repo-spawn-20261126-210337-5137` exactly as written above. A per-run-unique token is already baked into both. Do not invent, regenerate or alter either one. The scenario's final assertion checks for this exact text in this repository.
- The PR for this issue must include `Closes #<this issue's number>` in its body. It must not use any closing keyword (`Closes`/`Fixes`/`Resolves`) next to any other issue number, including the alpha repository's issue.

## Acceptance

- `HelloE2E202611262103375137()` returns exactly `e2e-cross-repo-spawn-20261126-210337-5137`.
- `go test ./...` passes in this repository, including the new test.
- The text `HelloE2E202611262103375137` and `e2e-cross-repo-spawn-20261126-210337-5137` appears in `pkg/greeting/greeting.go` and `pkg/greeting/greeting_test.go`."

## Background

The alpha repository (`handarbeit/fabrik-test-alpha`) needs to call a new exported greeting function from this repository's `pkg/greeting` package. That function does not exist yet, so it must land here before alpha can import it. This work is one part of an end-to-end regression scenario for cross-repo decomposition (on-demand spawn-target repo init, and a per-run-unique fixture token). The scenario's final assertion looks for the exact function name and sentinel string in this repository, so both must be reproduced exactly as given. The package already contains several sibling `HelloE2E<timestamp><suffix>()` functions, each returning a sentinel string and carrying a one-line doc comment; this one follows the same pattern.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Expose the sentinel greeting function (Priority: P1)

A consumer in the alpha repository imports this repository's `pkg/greeting` package and calls `HelloE2E202611262103375137()` to obtain the run-unique sentinel string.

**Why this priority**: It is the only deliverable. Without it, the alpha repository cannot import the function and the cross-repo regression scenario cannot complete.

**Independent Test**: Run the package's unit tests. The new test calls the function and compares its result to the expected literal; the scenario's final check also searches the repository for the function name and the literal.

**Acceptance Scenarios**:

1. **Given** the `pkg/greeting` package, **When** `HelloE2E202611262103375137()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261126-210337-5137` with no extra whitespace or newline.
2. **Given** the repository after the change, **When** the full test suite is run, **Then** all tests pass, including `TestHelloE2E202611262103375137`.
3. **Given** the repository after the change, **When** its source is searched for `HelloE2E202611262103375137` and `e2e-cross-repo-spawn-20261126-210337-5137`, **Then** both strings appear in the greeting implementation file and in its test file.

---

### Edge Cases

- The returned string must match byte for byte: no trailing newline, no leading or trailing whitespace, and no change to any character of the token.
- Existing `HelloE2E*` functions and their tests must remain unchanged, so existing behaviour and sibling sentinels are not affected.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The `pkg/greeting` package MUST export a function named exactly `HelloE2E202611262103375137` that takes no arguments and returns a `string`.
- **FR-002**: `HelloE2E202611262103375137()` MUST return exactly the literal `e2e-cross-repo-spawn-20261126-210337-5137`.
- **FR-003**: The function MUST carry a one-line doc comment in the same style as the existing `HelloE2E*` functions, for example `HelloE2E202611262103375137 returns the sentinel string for the cross-repo spawn regression test.`
- **FR-004**: The package's tests MUST include a test named `TestHelloE2E202611262103375137` that declares the expected literal as `want`, calls the function, and reports any mismatch as a test error, following the shape of the existing tests.
- **FR-005**: No other function, test or file may be modified, and no dependency-tidying step is performed unless it proves necessary.
- **FR-006**: The pull request body MUST contain a closing reference only to this issue and MUST NOT put a closing keyword next to any other issue number, including the alpha repository's issue.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: Calling the new function returns exactly `e2e-cross-repo-spawn-20261126-210337-5137` (verified by the new test).
- **SC-002**: The full test run for the repository passes with zero failures.
- **SC-003**: Both the function name and the sentinel literal appear in the greeting implementation file and in its test file.
- **SC-004**: The change set touches only the greeting implementation file and its test file, plus the persisted spec file the engine commits for this issue.

## Assumptions

- The function name and the sentinel literal are fixed by the scenario and are used exactly as given, with no regeneration or alteration.
- The existing sibling `HelloE2E*` functions and tests define the style to follow, so the new function is added alongside them.
- No dependency changes are needed, so dependency tidying is skipped.

## Out of Scope *(optional)*

- Any change to the alpha repository, including importing or calling the new function there.
- Changes to, or removal of, any existing `HelloE2E*` function or test.
- Refactoring or consolidating the sentinel functions.
- Any other file, documentation or configuration change.
