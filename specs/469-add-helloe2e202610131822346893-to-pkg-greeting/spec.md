# Feature Specification: Add HelloE2E202610131822346893 to pkg/greeting

**Feature Branch**: `fabrik/issue-469`
**Created**: 2026-10-13
**Status**: Specified
**Input**: User description:

~~~~
## Problem

The alpha repo (`handarbeit/fabrik-test-alpha`) needs to call a new greeting function from this repo as part of a cross-repo e2e regression test (#803 on-demand spawn-target repo init, #1308 per-run-unique fixture token). Alpha cannot import the function until it exists on this repo's `main`, so this change must land first.

## Requirements

1. In `pkg/greeting/greeting.go`, append an exported function with a doc comment, following the existing pattern of the preceding `HelloE2E…` functions:

   ```go
   // HelloE2E202610131822346893 returns the sentinel string for the cross-repo spawn regression test.
   func HelloE2E202610131822346893() string { return "e2e-cross-repo-spawn-20261013-182234-6893" }
   ```

2. In `pkg/greeting/greeting_test.go`, append a test for it, following the pattern of the existing per-function tests:

   ```go
   func TestHelloE2E202610131822346893(t *testing.T) {
   	const want = "e2e-cross-repo-spawn-20261013-182234-6893"
   	if got := HelloE2E202610131822346893(); got != want {
   		t.Errorf("HelloE2E202610131822346893() = %q, want %q", got, want)
   	}
   }
   ```

3. Use the exact function name `HelloE2E202610131822346893` and the exact string literal `e2e-cross-repo-spawn-20261013-182234-6893` verbatim. Do not invent your own: a per-run-unique token is baked into both, and the scenario's final assertion depends on that exact text landing in this repo.

## Scope

- Only the two files above change. Both files are appended to by every run, so place the new blocks at the end and keep any existing functions and tests intact. If the end of either file conflicts with a concurrent append, keep both blocks.
- No new dependencies, and no changes to any other file.

## Acceptance

- `go build ./...`, `go vet ./...` and `go test ./...` pass.
- `HelloE2E202610131822346893()` returns exactly `e2e-cross-repo-spawn-20261013-182234-6893`.
- The change is merged to this repo's default branch so the alpha repo can pin to a commit that contains it.
~~~~

*Spawned by Fabrik from parent issue handarbeit/fabrik-test-alpha#7757 as a multi-issue decomposition. The parent's plan is at the link above.*

## Background

The alpha repo (`handarbeit/fabrik-test-alpha`) runs a cross-repo e2e regression test covering on-demand spawn-target repo initialisation (#803) and per-run-unique fixture tokens (#1308). As part of that scenario, alpha must call a new greeting function that lives in this repo's `pkg/greeting` package. Alpha cannot import the function until it exists on this repo's default branch, so this change has to land first. The function name and its returned string both embed a per-run-unique token (`20261013-182234-6893`), and the scenario's final assertion depends on that exact text appearing in this repo.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Sentinel greeting function available on main (Priority: P1)

The alpha repo's cross-repo e2e scenario can pin to a commit of this repo that exports `HelloE2E202610131822346893` and call it to obtain the run's sentinel string.

**Why this priority**: It is the only deliverable. The alpha scenario cannot proceed or assert success without it.

**Independent Test**: Run the new unit test in `pkg/greeting` and confirm it passes. Then check that the function and literal exist on the default branch after merge.

**Acceptance Scenarios**:

1. **Given** the change is merged, **When** a caller invokes `HelloE2E202610131822346893()`, **Then** it returns exactly `e2e-cross-repo-spawn-20261013-182234-6893`.
2. **Given** the change is applied, **When** `go build ./...`, `go vet ./...` and `go test ./...` run, **Then** all three succeed.
3. **Given** earlier `HelloE2E…` functions and tests exist, **When** the change is applied, **Then** all of them remain present and unmodified.

---

### Edge Cases

- A concurrent run appends its own block to the same two files: both blocks must be kept, so a conflict at the end of either file is resolved by retaining both.
- The token in the function name and the string literal must match the run's token exactly; a differing or regenerated token breaks the scenario's final assertion.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `pkg/greeting/greeting.go` exports a function named exactly `HelloE2E202610131822346893` that takes no arguments and returns the string `e2e-cross-repo-spawn-20261013-182234-6893`.
- **FR-002**: The function carries a doc comment in the same style as the preceding `HelloE2E…` functions, describing it as the sentinel string for the cross-repo spawn regression test.
- **FR-003**: `pkg/greeting/greeting_test.go` contains a test named exactly `TestHelloE2E202610131822346893` that asserts the function returns exactly `e2e-cross-repo-spawn-20261013-182234-6893`, following the existing per-function test pattern.
- **FR-004**: Both additions are appended at the end of their files, and all existing functions and tests are left intact.
- **FR-005**: No files other than the two named above change, and no new dependencies are introduced.
- **FR-006**: The function name and string literal use the provided token text verbatim and are not regenerated or altered.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `go build ./...`, `go vet ./...` and `go test ./...` all pass with the change applied.
- **SC-002**: The new test passes and asserts the exact expected string.
- **SC-003**: The diff touches exactly two files and only adds lines (no removals of existing functions or tests).
- **SC-004**: The change is merged to the default branch, so the alpha repo can pin to a commit containing `HelloE2E202610131822346893`.

## Assumptions

- The preceding `HelloE2E…` functions and tests in these files provide the pattern to follow; the snippets in the request are authoritative for the exact code.
- Concurrent appends from other runs are resolved by keeping both blocks.

## Out of Scope *(optional)*

- Changes to the alpha repo or to its e2e scenario.
- Any refactoring, removal or consolidation of existing `HelloE2E…` functions and tests.
- Any files other than `pkg/greeting/greeting.go` and `pkg/greeting/greeting_test.go`.
- New dependencies.
