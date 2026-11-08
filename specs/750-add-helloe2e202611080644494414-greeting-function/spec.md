# Feature Specification: Add HelloE2E202611080644494414 greeting function and test

**Feature Branch**: `fabrik/issue-750`
**Created**: 2026-11-08
**Status**: Specified
**Input**: User description: "The alpha repo (`handarbeit/fabrik-test-alpha`) is about to call a new sentinel function from this repo's `pkg/greeting` package. Alpha cannot compile until the function exists on this repo's default branch, so this change must land first. It is the beta half of a cross-repo e2e regression scenario (#803 on-demand spawn-target repo init, #1308 per-run-unique fixture token). The parent issue is `handarbeit/fabrik-test-alpha#11258`. Requirements: (1) In `pkg/greeting/greeting.go`, add an exported function `HelloE2E202611080644494414() string` that returns exactly the string literal `e2e-cross-repo-spawn-20261108-064449-4414`. (2) In `pkg/greeting/greeting_test.go`, add a test `TestHelloE2E202611080644494414` that asserts the function returns exactly that literal. (3) Use the function name and the string literal EXACTLY as written. Do not rename, regenerate, shorten or alter either one. A per-run-unique token is already baked into them, and the scenario's final assertion depends on this exact text landing in this repo's default branch. The function-name digits `202611080644494414` are the literal's date, time and run number (`20261108`, `064449`, `4414`) with the dashes removed. Conventions: package `greeting`; append the new function at the END of `pkg/greeting/greeting.go`, after the last existing sentinel function (currently `HelloE2E202611080130099963`); follow the existing sentinel style with the doc comment `// HelloE2E202611080644494414 returns the sentinel string for the cross-repo spawn regression test.`; append the test at the end of `pkg/greeting/greeting_test.go` following the existing pattern (`const want = ...`, call the function, on mismatch `t.Errorf(\"HelloE2E202611080644494414() = %q; want %q\", got, want)`); add the spec directory under `specs/` that previous sentinel issues add (`specs/<issue-number>-add-helloe2e202611080644494414-greeting-function/`); leave all existing functions and tests unchanged; if a concurrent change leaves a trivial conflict at the tail of either file, keep both sides. Acceptance: the function is exported and returns exactly the literal; the test exists and `go build ./...`, `go vet ./...` and `go test ./...` pass; after merge the default branch contains both the function name and the literal. Out of scope: changes to existing greeting behaviour or other functions; releasing or tagging this repo (alpha pins this repo by pseudo-version); the alpha-side change (`main.go`, `main_test.go`, `go.mod`/`go.sum` pin bump), tracked in `handarbeit/fabrik-test-alpha#11258`."

## Background

The alpha repo (`handarbeit/fabrik-test-alpha`) is about to call a new sentinel function from this repo's `pkg/greeting` package. Alpha cannot compile until that function exists on this repo's default branch, so this change must land first.

This issue is the beta half of a cross-repo e2e regression scenario. The scenario exercises on-demand initialisation of a spawn-target repo (#803) and a per-run-unique fixture token (#1308). The parent issue is `handarbeit/fabrik-test-alpha#11258`.

The scenario's final assertion checks that this exact function name and literal appear on this repo's default branch. The identifiers were generated once for this run, so any alteration breaks the scenario.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Sentinel function available on the default branch (Priority: P1)

A maintainer of the cross-repo regression scenario needs the `pkg/greeting` package to export `HelloE2E202611080644494414`, returning the per-run-unique literal. The alpha repo can then compile against it, and the scenario's final assertion can find the text on the default branch.

**Why this priority**: It is the only deliverable. Alpha's change is blocked until it lands, and the scenario fails without it.

**Independent Test**: Run `go test ./pkg/greeting/...`. `TestHelloE2E202611080644494414` passes, and a text search of the default branch after merge finds both the function name and the literal.

**Acceptance Scenarios**:

1. **Given** the change is merged, **When** `HelloE2E202611080644494414()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261108-064449-4414`.
2. **Given** the change is merged, **When** `go build ./...`, `go vet ./...` and `go test ./...` run, **Then** all succeed.
3. **Given** the change is merged, **When** the default branch is searched, **Then** both `HelloE2E202611080644494414` and `e2e-cross-repo-spawn-20261108-064449-4414` are present.

---

### Edge Cases

- A concurrent change touching the tail of `pkg/greeting/greeting.go` or `greeting_test.go` may produce a trivial conflict. Keep both sides.
- The function-name digits must stay consistent with the literal (`20261108` + `064449` + `4414`). A transcription slip in either one is a defect, not a variation.
- Existing sentinel functions and tests must remain unchanged.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `pkg/greeting` (package `greeting`) MUST export `HelloE2E202611080644494414() string`, returning exactly the string literal `e2e-cross-repo-spawn-20261108-064449-4414`.
- **FR-002**: The function name and the literal MUST be used exactly as given, with no renaming, regeneration, shortening or alteration.
- **FR-003**: The function MUST be appended at the end of `pkg/greeting/greeting.go`, after the last existing sentinel function (currently `HelloE2E202611080130099963`). It MUST follow the existing sentinel style, with the doc comment `// HelloE2E202611080644494414 returns the sentinel string for the cross-repo spawn regression test.`
- **FR-004**: `pkg/greeting/greeting_test.go` MUST contain `TestHelloE2E202611080644494414`, appended at the end of the file. It MUST declare `const want = "e2e-cross-repo-spawn-20261108-064449-4414"`, call the function, and on mismatch fail with `t.Errorf("HelloE2E202611080644494414() = %q; want %q", got, want)`.
- **FR-005**: The change MUST add the spec directory that earlier sentinel issues in this repo add, `specs/750-add-helloe2e202611080644494414-greeting-function/`, matching the existing ones.
- **FR-006**: All existing functions and tests MUST remain unchanged. If a concurrent change leaves a trivial conflict at the tail of either file, both sides MUST be kept.

### Key Entities *(if applicable)*

- **Sentinel function**: An exported zero-argument function in `pkg/greeting` that returns a fixed, per-run-unique string. Other repos and e2e assertions use it as a marker.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After merge, `HelloE2E202611080644494414()` returns exactly `e2e-cross-repo-spawn-20261108-064449-4414`, as verified by `TestHelloE2E202611080644494414`.
- **SC-002**: `go build ./...`, `go vet ./...` and `go test ./...` all pass.
- **SC-003**: After merge, this repo's default branch contains both `HelloE2E202611080644494414` and `e2e-cross-repo-spawn-20261108-064449-4414`.
- **SC-004**: The alpha repo can compile against the merged commit, unblocking `handarbeit/fabrik-test-alpha#11258`.

## Assumptions

- The existing sentinel functions, tests and spec directories in this repo are the pattern to follow, and the `specs/` directory name follows the issue number as given in the requirements.
- The default branch is `main`.
- No behaviour of existing greeting functions needs to change.

## Out of Scope *(optional)*

- Any change to existing greeting behaviour or to other functions.
- Releasing or tagging this repo. The alpha repo pins this repo by pseudo-version.
- The alpha-side change (`main.go`, `main_test.go`, `go.mod`/`go.sum` pin bump), which is tracked in `handarbeit/fabrik-test-alpha#11258`.

## Source References *(optional)*

- Parent issue: `handarbeit/fabrik-test-alpha#11258`
- Related scenarios: #803 (on-demand spawn-target repo init), #1308 (per-run-unique fixture token)
