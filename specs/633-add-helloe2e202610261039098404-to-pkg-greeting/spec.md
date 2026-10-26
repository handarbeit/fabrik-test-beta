# Feature Specification: Add HelloE2E202610261039098404 to pkg/greeting (e2e cross-repo spawn 20261026-103909-8404)

**Feature Branch**: `fabrik/issue-633`
**Created**: 2026-10-26
**Status**: Specified
**Input**: User description:

> ## Goal
>
> Add an exported greeting function and a test for it to this repository. This is the beta half of an end-to-end cross-repo spawn scenario. The parent issue is in `handarbeit/fabrik-test-alpha` (#9901) and cannot build until this lands, because alpha's `main.go` will call this function.
>
> ## Requirements
>
> 1. In `pkg/greeting/greeting.go`, append an exported function with no arguments that returns a `string`:
>
>    ```go
>    // HelloE2E202610261039098404 returns the e2e cross-repo spawn sentinel for run 20261026-103909-8404.
>    func HelloE2E202610261039098404() string {
>    	return "e2e-cross-repo-spawn-20261026-103909-8404"
>    }
>    ```
>
>    Append it at the end of the file, after the existing `HelloE2E202610260610197278` entry. Follow the existing one-function-per-sentinel pattern: a doc comment, then the function.
>
> 2. In `pkg/greeting/greeting_test.go`, append a test at the end of the file, after the existing `HelloE2E202610260610197278` test. It uses the existing shape:
>
>    ```go
>    func TestHelloE2E202610261039098404(t *testing.T) {
>    	const want = "e2e-cross-repo-spawn-20261026-103909-8404"
>    	if got := HelloE2E202610261039098404(); got != want {
>    		t.Errorf("HelloE2E202610261039098404() = %q, want %q", got, want)
>    	}
>    }
>    ```
>
>    Match the exact style of the neighbouring tests if it differs in detail.
>
> ## Constraints
>
> - Use the function name `HelloE2E202610261039098404` and the literal `e2e-cross-repo-spawn-20261026-103909-8404` exactly as written above, character for character. This token is fixed for this run: do not invent, regenerate or alter it. A scenario assertion checks for this exact text in this repository.
> - Do not change any other function, test, or file. Add no dependencies.
> - The existing `go test ./...` must still pass, including the new test.
> - No documentation changes are needed.
>
> ## Acceptance
>
> - `pkg/greeting` exports `HelloE2E202610261039098404()` returning exactly `e2e-cross-repo-spawn-20261026-103909-8404`.
> - A passing test for it exists in `pkg/greeting/greeting_test.go`.
> - The PR is merged to `main`, so that `handarbeit/fabrik-test-alpha` can advance its module pin to a commit containing the function.
>
> *Spawned by Fabrik from parent issue handarbeit/fabrik-test-alpha#9901 as a multi-issue decomposition. The parent's plan is at the link above.*

## Background

This repository is the beta half of an end-to-end cross-repo spawn scenario. The parent issue is `handarbeit/fabrik-test-alpha#9901`. Alpha's `main.go` will call a new sentinel function from this repository's `pkg/greeting` package, so alpha cannot build until the function exists in a commit on this repository's `main`. Alpha then advances its module pin to that commit.

`pkg/greeting` already follows a one-function-per-sentinel pattern. Each run adds one exported function that returns a fixed string, plus a matching test. A scenario assertion checks that this run's exact function name and string appear in this repository.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Sentinel function and test available on main (Priority: P1)

A consumer of `pkg/greeting` (alpha's `main.go`) can call `HelloE2E202610261039098404()` and get the fixed sentinel string for this run. A test in this repository proves it.

**Why this priority**: It is the only deliverable. Alpha's build is blocked until it is merged.

**Independent Test**: Run `go test ./...` in this repository. The new test passes, and every existing test still passes.

**Acceptance Scenarios**:

1. **Given** the change is merged, **When** a caller invokes `greeting.HelloE2E202610261039098404()`, **Then** it returns exactly `e2e-cross-repo-spawn-20261026-103909-8404`.
2. **Given** the change is merged, **When** `go test ./...` runs, **Then** `TestHelloE2E202610261039098404` passes and no other test fails.
3. **Given** the change is on `main`, **When** alpha advances its module pin to a commit containing it, **Then** alpha can resolve `HelloE2E202610261039098404`.

---

### Edge Cases

- A typo or alteration in the function name or the string literal makes the scenario assertion fail. Both must match the issue text character for character.
- Placing the additions anywhere other than the end of each file would break the append-only convention and risk merge conflicts with sibling sentinel issues.
- Existing sentinels must not change. Their order and content are untouched.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `pkg/greeting` exports a no-argument function named exactly `HelloE2E202610261039098404` that returns a `string`.
- **FR-002**: That function returns exactly `e2e-cross-repo-spawn-20261026-103909-8404`.
- **FR-003**: The function has a doc comment and is appended at the end of `pkg/greeting/greeting.go`, after the `HelloE2E202610260610197278` entry.
- **FR-004**: `pkg/greeting/greeting_test.go` contains `TestHelloE2E202610261039098404`, appended at the end of the file after the `HelloE2E202610260610197278` test. It asserts the function returns the string from FR-002.
- **FR-005**: The new test matches the style of the neighbouring tests, using a `const want` and a `got != want` check.
- **FR-006**: No other function, test, file or dependency is changed or added, and there are no documentation changes.
- **FR-007**: The change is merged to `main`.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `go test ./...` passes, including the new test.
- **SC-002**: The exact text `HelloE2E202610261039098404` and `e2e-cross-repo-spawn-20261026-103909-8404` is present in `pkg/greeting` on `main`.
- **SC-003**: The change touches only the two files named above, and only by appending.
- **SC-004**: Alpha can advance its module pin to a commit containing the function.

## Assumptions

- Where the issue's example test differs in small details from the neighbouring tests (for example `want %q` separated by `;` rather than `,` in the error message), the neighbouring style wins, as the issue instructs. The function name and string literal are never altered.
- The doc comment text from the issue is used as written, even if its wording differs from older sentinels' comments.
- Merging to `main` is handled by the normal pipeline (the issue carries `fabrik:yolo`).

## Out of Scope

- Changes in `handarbeit/fabrik-test-alpha`, including advancing its module pin.
- Modifying, removing or refactoring existing sentinel functions or tests.
- Documentation changes and new dependencies.
