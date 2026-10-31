# Feature Specification: Add HelloE2E202610311432580141 sentinel to the greeting package (e2e cross-repo spawn)

**Feature Branch**: `fabrik/issue-684`
**Created**: 2026-10-31
**Status**: Specified
**Input**: User description: "Add per-run greeting function for the e2e cross-repo spawn scenario" — the full original request is preserved verbatim below.

> # Add per-run greeting function for the e2e cross-repo spawn scenario
>
> This is the beta-side half of an end-to-end regression scenario for cross-repo decomposition (#803 on-demand spawn-target repo init, #1308 per-run-unique fixture token). The alpha repo (`handarbeit/fabrik-test-alpha`) will import this function afterwards, so this work must land on the default branch first.
>
> ## Requirements
>
> - In `pkg/greeting/greeting.go`, add the exported function `HelloE2E202610311432580141() string` that returns exactly the literal string `e2e-cross-repo-spawn-20261031-143258-0141`. Follow the existing pattern in the file: a doc comment `// HelloE2E202610311432580141 returns the sentinel string for the cross-repo spawn regression test.`, appended at the end of the file after the existing `HelloE2E…` functions.
> - In `pkg/greeting/greeting_test.go`, add `TestHelloE2E202610311432580141` following the existing test pattern: `const want = "e2e-cross-repo-spawn-20261031-143258-0141"` and `if got := HelloE2E202610311432580141(); got != want { t.Errorf(...) }`.
>
> ## Constraints
>
> - Use the function name `HelloE2E202610311432580141` and the literal `e2e-cross-repo-spawn-20261031-143258-0141` verbatim. A per-run-unique token is baked into them and the scenario's final assertion depends on that exact text landing in this repo. Do not invent, regenerate or paraphrase them.
> - Make no other change to the greeting package beyond this one function and its test.
>
> ## Acceptance
>
> - `pkg/greeting/greeting.go` exports `HelloE2E202610311432580141() string` returning exactly `e2e-cross-repo-spawn-20261031-143258-0141`.
> - The new test in `pkg/greeting/greeting_test.go` passes, and `go test ./...` passes in this repo.
> - The change is merged to the default branch so the alpha repo can bump its module pin to a commit containing it.
>
> *Spawned by Fabrik from parent issue handarbeit/fabrik-test-alpha#10484 as a multi-issue decomposition. The parent's plan is at the link above.*

## Background

This issue is the beta-repository half of an end-to-end regression scenario for cross-repo decomposition. The scenario covers on-demand initialisation of a spawn-target repository (#803) and a per-run-unique fixture token (#1308). A parent issue in the alpha repository (`handarbeit/fabrik-test-alpha#10484`) was decomposed into this child.

The alpha repository will import the new greeting function once it exists, and the scenario's final assertion depends on the exact per-run token text landing in this repository. This change therefore has to be merged to the default branch (`main`) first, so the alpha repository can bump its module pin to a commit that contains it.

The existing greeting package already holds several sentinel functions of the same kind (`HelloE2E…`), each with a matching test. This issue adds one more in the same style.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Expose the per-run sentinel from the greeting package (Priority: P1)

A consumer of the greeting package (in this scenario, the alpha repository) can call a new exported function and receive the per-run sentinel string for this scenario run.

**Why this priority**: It is the only deliverable. The alpha repository's half of the scenario is blocked until this lands on the default branch.

**Independent Test**: Call the new function from a test in the same package and compare the result to the expected literal.

**Acceptance Scenarios**:

1. **Given** the greeting package on the default branch, **When** a caller invokes `HelloE2E202610311432580141()`, **Then** it returns exactly `e2e-cross-repo-spawn-20261031-143258-0141`.
2. **Given** the new test `TestHelloE2E202610311432580141`, **When** the repository's full test suite runs, **Then** the test passes and the suite passes overall.
3. **Given** the change is merged to the default branch, **When** the alpha repository bumps its module pin to a commit containing the merge, **Then** it can import and call the function.

---

### Edge Cases

- The function name and returned literal must match the per-run token exactly. A paraphrased, regenerated or transposed token would break the scenario's final assertion without failing this repository's own tests.
- Existing sentinel functions and their tests must be left untouched. The new function is appended after them.
- The function takes no arguments and has no side effects. It returns a constant.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The greeting package MUST export a function named exactly `HelloE2E202610311432580141` that takes no arguments and returns a `string`.
- **FR-002**: The function MUST return exactly the literal `e2e-cross-repo-spawn-20261031-143258-0141`, with no added or altered characters.
- **FR-003**: The function MUST carry the doc comment `// HelloE2E202610311432580141 returns the sentinel string for the cross-repo spawn regression test.` and MUST be appended at the end of the file, after the existing `HelloE2E…` functions, following the existing pattern.
- **FR-004**: The greeting package's tests MUST include `TestHelloE2E202610311432580141`, following the existing test pattern. It declares `const want = "e2e-cross-repo-spawn-20261031-143258-0141"` and fails with `t.Errorf` if `HelloE2E202610311432580141()` does not equal `want`.
- **FR-005**: The change MUST NOT alter anything else in the greeting package beyond this one function and its test.
- **FR-006**: The change MUST be merged to the default branch (`main`) so a commit containing it is available for the alpha repository's module pin.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: Calling the new function returns a string byte-for-byte equal to `e2e-cross-repo-spawn-20261031-143258-0141`.
- **SC-002**: The new test passes, and the full repository test suite passes with no regressions.
- **SC-003**: The diff to the greeting package consists solely of the one new function and the one new test (no other lines changed).
- **SC-004**: A commit containing the function exists on the default branch, so the alpha repository's scenario can resolve and assert on the exact token.

## Assumptions

- The existing `HelloE2E…` functions and tests in the greeting package define the "existing pattern". The new additions mirror them in style, and the new function goes at the end of the file after them.
- The token in the function name and literal is the final per-run value supplied by the scenario. It is used as given and never regenerated.
- Landing on the default branch happens through the normal pipeline PR and merge. This issue defines no separate release step.
- The alpha repository's import and module-pin bump are handled by its own issue.

## Out of Scope *(optional)*

- Any other change to the greeting package, including edits, refactors or removal of existing sentinel functions and tests.
- The alpha repository's import of this function and its module-pin bump.
- Changes to the cross-repo spawn mechanism (#803) or the per-run fixture token generation (#1308).
- Documentation or release-note changes.
