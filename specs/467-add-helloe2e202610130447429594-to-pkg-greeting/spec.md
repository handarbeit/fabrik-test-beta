# Feature Specification: Add HelloE2E202610130447429594 sentinel to pkg/greeting

**Feature Branch**: `fabrik/issue-467`
**Created**: 2026-10-13
**Status**: Specified
**Input**: User description: "Goal: Add one exported function and its test to `pkg/greeting`. This is the beta half of a cross-repo e2e regression test (handarbeit/fabrik#803 on-demand spawn-target repo init, handarbeit/fabrik#1308 per-run-unique fixture token). The alpha repo (`handarbeit/fabrik-test-alpha`) imports this function and cannot build until it exists on beta `main`, so this change must be merged first. Requirements: In `pkg/greeting/greeting.go`, add an exported function with exactly this name and signature: `HelloE2E202610130447429594() string`. It returns exactly the literal string `\"e2e-cross-repo-spawn-20261013-044742-9594\"`. Use the function name and string literal verbatim; do not invent your own. A per-run-unique token is already baked into both, and the scenario's final assertion depends on this exact text landing in this repo. Follow the existing convention in the file: one exported function per e2e run, each with a doc comment such as `// HelloE2E202610130447429594 returns the sentinel string for the cross-repo spawn regression test.` In `pkg/greeting/greeting_test.go`, add `TestHelloE2E202610130447429594`, following the shape of the existing per-function tests: a `const want = \"e2e-cross-repo-spawn-20261013-044742-9594\"` and a check that fails with `t.Errorf` if `HelloE2E202610130447429594() != want`. Do not modify or remove any existing function or test. Scope: Only `pkg/greeting/greeting.go` and `pkg/greeting/greeting_test.go` in this repo. No `go.mod` changes are needed (the module has no dependencies). Do not change anything in the alpha repo; that is handled by its own issue after this one lands. Acceptance: `go build ./...` and `go test ./...` pass. `grep -c HelloE2E202610130447429594 pkg/greeting/greeting.go` is at least 1, and the returned literal is exactly `e2e-cross-repo-spawn-20261013-044742-9594`. The change is merged to beta `main`, so alpha can pin a beta commit that contains it."

## Background

This issue is the beta-repo half of a cross-repo end-to-end regression test. It exercises on-demand initialisation of a spawn-target repository (handarbeit/fabrik#803) and a per-run-unique fixture token (handarbeit/fabrik#1308). The issue was spawned from parent issue handarbeit/fabrik-test-alpha#7684 as part of a multi-issue decomposition.

The alpha repo (`handarbeit/fabrik-test-alpha`) imports the new function. Alpha cannot build until the function exists on beta `main`, so this change must land first. The scenario's final assertion checks that the exact sentinel text appears in this repo.

`pkg/greeting` already holds one sentinel function per earlier e2e run. Each has a doc comment and its own test, and each returns a run-unique string of the form `e2e-cross-repo-spawn-<date>-<time>-<suffix>`. This change adds the next one.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Expose the run's sentinel function (Priority: P1)

The e2e scenario needs `pkg/greeting` to export a function named for this run that returns this run's unique string. The alpha repo can then import it, and the scenario can confirm that the text landed in beta.

**Why this priority**: It is the only deliverable. Without it, alpha cannot build and the cross-repo scenario cannot complete.

**Independent Test**: Call `HelloE2E202610130447429594()` and compare the result with `e2e-cross-repo-spawn-20261013-044742-9594`. Run `go build ./...` and `go test ./...` to confirm nothing else is affected.

**Acceptance Scenarios**:

1. **Given** the change is merged to beta `main`, **When** a consumer imports `pkg/greeting` and calls `HelloE2E202610130447429594()`, **Then** it receives exactly `e2e-cross-repo-spawn-20261013-044742-9594`.
2. **Given** the change is applied, **When** `go build ./...` and `go test ./...` run, **Then** both succeed.
3. **Given** the change is applied, **When** `grep -c HelloE2E202610130447429594 pkg/greeting/greeting.go` runs, **Then** it reports at least 1.

---

### User Story 2 - Test the sentinel function (Priority: P1)

A test protects the sentinel's name and value, in the same shape as the existing per-function tests.

**Why this priority**: The acceptance criteria require `go test ./...` to pass, and the test guards the exact string the cross-repo assertion depends on.

**Independent Test**: Run `go test ./pkg/greeting/...`. `TestHelloE2E202610130447429594` passes, and it fails if the function's return value changes.

**Acceptance Scenarios**:

1. **Given** the function returns the specified literal, **When** `TestHelloE2E202610130447429594` runs, **Then** it passes.
2. **Given** the function returned any other string, **When** the test runs, **Then** it fails via `t.Errorf` and reports both the actual and the expected value.

---

### Edge Cases

- The function name and the literal must be used verbatim. A transposed digit in the name or the string would break the alpha import or the final assertion.
- Existing sentinel functions and tests must remain unchanged. Alpha or earlier scenarios may still depend on them.
- Alpha must not be changed. Its half is a separate issue that follows this one.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `pkg/greeting` MUST export a function `HelloE2E202610130447429594() string`, with exactly this name and signature.
- **FR-002**: `HelloE2E202610130447429594()` MUST return exactly the string `e2e-cross-repo-spawn-20261013-044742-9594`.
- **FR-003**: The function MUST carry a doc comment in the established form: `// HelloE2E202610130447429594 returns the sentinel string for the cross-repo spawn regression test.`
- **FR-004**: A test named `TestHelloE2E202610130447429594` MUST exist alongside the other per-function tests. It declares `const want = "e2e-cross-repo-spawn-20261013-044742-9594"` and fails via `t.Errorf` if the function's result differs from `want`.
- **FR-005**: No existing function or test in `pkg/greeting` may be modified or removed.
- **FR-006**: The change MUST be limited to `pkg/greeting/greeting.go` and `pkg/greeting/greeting_test.go`. It MUST NOT touch `go.mod` or anything in the alpha repo.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `go build ./...` and `go test ./...` both pass on the change.
- **SC-002**: `grep -c HelloE2E202610130447429594 pkg/greeting/greeting.go` reports at least 1, and the returned literal is exactly `e2e-cross-repo-spawn-20261013-044742-9594`.
- **SC-003**: The change is merged to beta `main`, so alpha can pin a beta commit that contains the function.
- **SC-004**: The diff touches only the two in-scope files and removes no existing lines.

## Assumptions

- The module has no dependencies, so no `go.mod` change is needed (as stated in the issue).
- The function and test are appended after the existing sentinel entries, following the convention of one exported function per e2e run.
- Merging is handled by the normal pipeline; the `fabrik:yolo` label is present on this issue.

## Out of Scope

- Any change to the alpha repo (`handarbeit/fabrik-test-alpha`). Its half of the scenario is a separate issue that follows this one.
- Changes to or removal of earlier sentinel functions and tests.
- `go.mod` changes.

## Source References

- handarbeit/fabrik#803: on-demand spawn-target repo init
- handarbeit/fabrik#1308: per-run-unique fixture token
- handarbeit/fabrik-test-alpha#7684: parent issue for this decomposition
