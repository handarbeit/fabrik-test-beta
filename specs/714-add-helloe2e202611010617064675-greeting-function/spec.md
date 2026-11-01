# Feature Specification: Add HelloE2E202611010617064675 greeting function (e2e cross-repo spawn 20261101-061706-4675)

**Feature Branch**: `fabrik/issue-714`
**Created**: 2026-11-01
**Status**: Specified
**Input**: User description: "## Goal

Add one exported function and a test to the greeting package so the alpha repo (`handarbeit/fabrik-test-alpha`) can import it. This is part of an end-to-end regression scenario for cross-repo decomposition (#803, #1308). The alpha repo cannot build until this change has landed on `main`, so this work must merge first.

## Requirements

1. In `pkg/greeting/greeting.go`, add an exported function:

   ```go
   // HelloE2E202611010617064675 returns the sentinel string for the cross-repo spawn regression test.
   func HelloE2E202611010617064675() string {
       return "e2e-cross-repo-spawn-20261101-061706-4675"
   }
   ```

   Follow the style of the existing `HelloE2E…` functions in that file and add it after the most recent one (`HelloE2E202610311432580141`).

2. In `pkg/greeting/greeting_test.go`, add a test for the new function that asserts it returns exactly `\"e2e-cross-repo-spawn-20261101-061706-4675\"`. Mirror the style of the neighbouring `HelloE2E…` tests in that file (table-driven or individual, whichever is used there).

3. The function name `HelloE2E202611010617064675` and the literal `e2e-cross-repo-spawn-20261101-061706-4675` MUST be used verbatim. The token `20261101-061706-4675` is unique to this run, and the scenario's final assertion checks for this exact text in this repo. Do not invent or substitute any other token, casing or spelling.

## Scope

- Change only the single function and its test. Make no other changes to this repo.
- Do not use a closing keyword (`Closes`/`Fixes`/`Resolves`) for anything other than this issue.

## Acceptance

- `go test ./pkg/greeting/...` passes, including the new test.
- `HelloE2E202611010617064675()` returns exactly `e2e-cross-repo-spawn-20261101-061706-4675`.
- The change is merged to `main` so that the alpha repo can pin to a commit containing it."

## Background

This issue is one half of an end-to-end regression scenario for cross-repo decomposition (#803, #1308). A parent issue in the alpha repo (`handarbeit/fabrik-test-alpha#10674`) was decomposed by Fabrik into work in this repo (`fabrik-test-beta`) and work in the alpha repo. The alpha repo imports the `greeting` package from this repo, so it cannot build until a commit containing the new sentinel function exists on this repo's `main`. This work must therefore merge first.

The scenario's final assertion looks for the run-unique token `20261101-061706-4675` in this repo, so the exact spelling of the function name and the returned string matters. The `greeting` package already holds many similar sentinel functions, one per previous e2e run, and this adds the next in the series.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Expose the run-unique sentinel from the greeting package (Priority: P1)

The e2e scenario needs the greeting package to expose a sentinel function carrying this run's unique token. The alpha repo can then import it and the scenario's final assertion can find that token in this repo.

**Why this priority**: It is the only deliverable. The alpha repo's build and the scenario's final assertion both depend on it.

**Independent Test**: Run the greeting package tests. The new test passes and asserts the exact returned string.

**Acceptance Scenarios**:

1. **Given** the greeting package on `main` with this change merged, **When** `HelloE2E202611010617064675()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261101-061706-4675`.
2. **Given** the change is on the issue branch, **When** `go test ./pkg/greeting/...` is run, **Then** all tests pass, including the new one.
3. **Given** the change is merged to `main`, **When** the alpha repo pins to a commit on `main` containing it, **Then** the alpha repo can import and call the function.

---

### Edge Cases

- A misspelled name, a wrong casing or a substituted token fails the scenario's final text assertion. The name and literal must be exact.
- The new function must sit after `HelloE2E202610311432580141`, the most recent existing sentinel, so the series stays chronological.
- The new test must not duplicate or alter the tests of existing sentinels.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The greeting package MUST export a function named exactly `HelloE2E202611010617064675` that takes no arguments and returns a string.
- **FR-002**: `HelloE2E202611010617064675()` MUST return exactly `e2e-cross-repo-spawn-20261101-061706-4675`.
- **FR-003**: The function MUST carry a doc comment in the style of the existing `HelloE2E…` functions, describing it as the sentinel string for the cross-repo spawn regression test.
- **FR-004**: The function MUST be placed after the most recent existing sentinel, `HelloE2E202610311432580141`, in the same file as the other sentinels.
- **FR-005**: A test for the new function MUST assert that it returns exactly `e2e-cross-repo-spawn-20261101-061706-4675`, and MUST follow the style of the neighbouring `HelloE2E…` tests (table-driven or individual, whichever they use).
- **FR-006**: The function name and the returned literal MUST be used verbatim, with no other token, casing or spelling.
- **FR-007**: The change MUST be limited to the one function and its test. No other file or behaviour in this repo may change.
- **FR-008**: The PR for this issue MUST use a closing keyword only for this issue (#714) and for nothing else.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `go test ./pkg/greeting/...` passes with the new test included.
- **SC-002**: A search of the repo for `20261101-061706-4675` finds the new function's returned literal and the new test's expected value.
- **SC-003**: The change is merged to `main`, so the alpha repo can pin to a commit containing it and build.
- **SC-004**: The diff touches exactly two files, the greeting source and its test, and adds only the new function and its test.

## Assumptions

- The `greeting` package keeps its current location (`pkg/greeting`), so the issue's file references hold.
- The existing test file follows a consistent convention for sentinel tests. The new test follows the same convention, whichever form it is.
- The issue's code snippet is authoritative for the function's name, signature and return value. The surrounding doc-comment wording may vary slightly, as it already does between existing sentinels.

## Out of Scope

- Any change to existing sentinel functions or tests.
- Any change to the alpha repo or to the cross-repo decomposition machinery.
- Refactoring, deduplicating or tidying the growing list of sentinel functions.
- Documentation or tooling changes of any kind.

## Source References

- Parent issue: `handarbeit/fabrik-test-alpha#10674` (multi-issue decomposition that spawned this issue)
- Related e2e scenario issues: #803, #1308
