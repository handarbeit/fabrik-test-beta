# Feature Specification: Add HelloE2E202611261503499288 greeting function (e2e cross-repo spawn)

**Feature Branch**: `fabrik/issue-786`
**Created**: 2026-11-26
**Status**: Specified
**Input**: User description: "## Goal

Add a new exported greeting function to the `pkg/greeting` package and a test for it. The alpha repo (`handarbeit/fabrik-test-alpha`) will import this function, so this change must land on `main` first.

## Requirements

- In `pkg/greeting/greeting.go`, add an exported function `HelloE2E202611261503499288() string` that returns exactly the literal string `e2e-cross-repo-spawn-20261126-150349-9288`. Give it a doc comment in the same style as the existing functions, for example: `// HelloE2E202611261503499288 returns the sentinel string for the cross-repo spawn regression test.`
- In `pkg/greeting/greeting_test.go`, add `TestHelloE2E202611261503499288`. Follow the existing pattern: declare `const want = \"e2e-cross-repo-spawn-20261126-150349-9288\"` and fail with `t.Errorf(\"HelloE2E202611261503499288() = %q; want %q\", got, want)` if the function's return value differs.
- Use the function name and the string literal exactly as written above. Do not alter, regenerate or invent a different token. A per-run-unique token is already baked into them, and the scenario's final assertion depends on this exact text appearing in this repo.
- Do not change any other existing function or test.

## Acceptance

- `go build ./...` and `go test ./...` pass in this repo.
- `HelloE2E202611261503499288()` returns exactly `e2e-cross-repo-spawn-20261126-150349-9288`.
- The change is merged to `main`, so that the alpha repo can pin a commit that contains it.

## Context

This is the beta-side half of an end-to-end regression scenario for cross-repo decomposition (#803 on-demand spawn-target repo init, #1308 per-run-unique fixture token). The parent issue is `handarbeit/fabrik-test-alpha#11867`, which calls this function from its `main.go` once this change has landed."

## Background

This issue is the beta-side half of an end-to-end regression scenario for cross-repo decomposition (on-demand spawn-target repo initialisation, #803, and per-run-unique fixture tokens, #1308). The parent issue, `handarbeit/fabrik-test-alpha#11867`, decomposed its work into this issue in the beta repo. The alpha repo will call a new greeting function from its `main.go`, so the function must exist and be merged to `main` here first, so that alpha can pin a commit containing it.

The function name and the sentinel string carry a per-run-unique token that was generated before this issue was created. The scenario's final assertion looks for that exact text in this repo, so both must be used character for character.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Expose the sentinel greeting function (Priority: P1)

A consumer of the `pkg/greeting` package (the alpha repo) can call `HelloE2E202611261503499288()` and receive the fixed sentinel string `e2e-cross-repo-spawn-20261126-150349-9288`. A test in this repo guards that behaviour.

**Why this priority**: It is the only deliverable. The alpha repo's half of the scenario is blocked until it lands on `main`.

**Independent Test**: Run `go build ./...` and `go test ./...` in this repo. The new test `TestHelloE2E202611261503499288` passes, and the existing tests are unchanged and still pass.

**Acceptance Scenarios**:

1. **Given** the `pkg/greeting` package after the change, **When** `HelloE2E202611261503499288()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261126-150349-9288`.
2. **Given** the new test follows the existing test pattern, **When** the function returns any other value, **Then** the test fails with the message `HelloE2E202611261503499288() = "<got>"; want "e2e-cross-repo-spawn-20261126-150349-9288"`.
3. **Given** the change is merged to `main`, **When** the alpha repo pins a commit of this repo, **Then** that commit contains the function and its test.

---

### Edge Cases

- A mistyped, regenerated or "tidied" token in the function name or in the string literal would break the scenario's final assertion. The text must match the issue exactly.
- The existing greeting functions and their tests must not be modified, renamed or reordered in a way that changes their behaviour.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The `pkg/greeting` package MUST export a function `HelloE2E202611261503499288() string` (defined in `pkg/greeting/greeting.go`) that returns exactly the literal string `e2e-cross-repo-spawn-20261126-150349-9288`.
- **FR-002**: The function MUST have a doc comment in the same style as the existing functions in the package, for example `// HelloE2E202611261503499288 returns the sentinel string for the cross-repo spawn regression test.`
- **FR-003**: `pkg/greeting/greeting_test.go` MUST contain `TestHelloE2E202611261503499288`, following the existing test pattern. It declares `const want = "e2e-cross-repo-spawn-20261126-150349-9288"` and fails with `t.Errorf("HelloE2E202611261503499288() = %q; want %q", got, want)` when the function's return value differs.
- **FR-004**: The function name and string literal MUST be used exactly as written in FR-001 and FR-003; the token MUST NOT be altered, regenerated or replaced.
- **FR-005**: No other existing function or test in the repository may be changed.
- **FR-006**: The change MUST be merged to `main` so that the alpha repo can pin a commit containing it.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `go build ./...` and `go test ./...` both pass in this repo with the change applied.
- **SC-002**: Calling `HelloE2E202611261503499288()` returns exactly `e2e-cross-repo-spawn-20261126-150349-9288`, verified by the new test.
- **SC-003**: The exact text `HelloE2E202611261503499288` and `e2e-cross-repo-spawn-20261126-150349-9288` appears in this repo on `main` after the merge, so the scenario's final assertion can find it.
- **SC-004**: The Go changes touch only the greeting source file and its test file, with no modifications to existing functions or tests. The only other file added is this spec under `specs/` (persisted by the Specify stage).

## Assumptions

- The existing functions in `pkg/greeting` and their tests establish a doc-comment and test style that the new code mirrors.
- Merging to `main` is handled by the normal Fabrik pipeline for this issue (it carries `fabrik:yolo`).
- The parent issue `handarbeit/fabrik-test-alpha#11867` handles the alpha-side import and call and needs no change from this issue.

## Out of Scope

- Any change to the alpha repo, including importing or calling the new function.
- Changes to, or refactoring of, any existing greeting function or test.
- Generating, validating or rotating the per-run token; it is fixed by the issue.

## Source References

- Parent issue: `handarbeit/fabrik-test-alpha#11867` (multi-issue decomposition; the parent's plan describes the cross-repo split).
- Related scenario issues, mentioned for context: #803, #1308.
