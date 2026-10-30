# Feature Specification: Add HelloE2E202610301951096404 greeting function (e2e cross-repo spawn)

**Feature Branch**: `fabrik/issue-653`
**Created**: 2026-10-30
**Status**: Specified
**Input**: User description: "# Feature: add `HelloE2E202610301951096404` to the greeting package

## Context

This issue is spawned from `handarbeit/fabrik-test-alpha#10234`, an end-to-end regression scenario for cross-repo decomposition (#803 on-demand spawn-target init, #1308 per-run-unique fixture token). The alpha repo's `main.go` will call this function once it lands, so this change must be merged to the default branch first.

The function name and string literal below contain a per-run-unique token. Use them verbatim. Do not rename them, shorten them or invent other tokens, because the scenario's final assertion checks for this exact text in this repo's default branch.

## Requirements

1. In `pkg/greeting/greeting.go`, add an exported function:
   - Name: `HelloE2E202610301951096404`
   - Signature: `func HelloE2E202610301951096404() string`
   - Return value: exactly the literal `\"e2e-cross-repo-spawn-20261030-195109-6404\"`
   - Follow the existing append-only convention. Append it after the most recent `HelloE2E…` function, with a doc comment in the same style, for example `// HelloE2E202610301951096404 returns the sentinel string for the cross-repo spawn regression test.`
2. In `pkg/greeting/greeting_test.go`, append a test in the existing per-function style:
   ```go
   func TestHelloE2E202610301951096404(t *testing.T) {
       const want = \"e2e-cross-repo-spawn-20261030-195109-6404\"
       if got := HelloE2E202610301951096404(); got != want {
           t.Errorf(\"HelloE2E202610301951096404() = %q; want %q\", got, want)
       }
   }
   ```
3. Do not modify any other function, test or file. Do not refactor the greeting package. If the package or either file does not exist yet, create it with the same package name (`greeting`) and the module path `github.com/handarbeit/fabrik-test-beta`.

## Acceptance

- `pkg/greeting/greeting.go` on the default branch contains both `HelloE2E202610301951096404` and `e2e-cross-repo-spawn-20261030-195109-6404`.
- `go build ./...` and `go test ./...` pass, including the new test.
- The change is merged to the default branch. Alpha issue #10234 stays blocked until this issue is closed.

## Out of scope

Any engine behaviour change, other greeting functions, or work in other repositories."

## Background

This issue was spawned by Fabrik from parent issue `handarbeit/fabrik-test-alpha#10234`. That parent is an end-to-end regression scenario for cross-repo decomposition. It covers on-demand spawn-target initialisation (#803) and a per-run-unique fixture token (#1308).

The alpha repository's `main.go` will call the new function once it exists, so this change must be merged to this repository's default branch (`main`) first. Alpha issue #10234 stays blocked until this issue is closed.

The scenario's final assertion checks for the exact function name and sentinel string in this repository's default branch. Both values contain a per-run-unique token and must be used verbatim, with no renaming, shortening or substitution.

The greeting package already exists and follows an append-only convention. Each prior end-to-end run added one `HelloE2E<token>` function and one matching test.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Sentinel greeting function is available on the default branch (Priority: P1)

The cross-repo spawn regression scenario needs a function with a known name that returns a known sentinel string. The alpha repository can then call it, and the scenario can assert that the exact text landed on `main`.

**Why this priority**: This is the only deliverable of the issue. The parent scenario cannot complete without it.

**Independent Test**: Build and test this repository, then inspect `pkg/greeting/greeting.go` on `main` for both the function name and the sentinel literal.

**Acceptance Scenarios**:

1. **Given** the greeting package, **When** `HelloE2E202610301951096404()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261030-195109-6404`.
2. **Given** the change is merged, **When** `pkg/greeting/greeting.go` on `main` is searched, **Then** it contains both `HelloE2E202610301951096404` and `e2e-cross-repo-spawn-20261030-195109-6404`.
3. **Given** the change is applied, **When** `go build ./...` and `go test ./...` run, **Then** both pass, including `TestHelloE2E202610301951096404`.

---

### Edge Cases

- A naming or literal typo: any deviation from the specified token makes the parent scenario's final assertion fail, so both values must be copied exactly.
- Placement: the function must be appended after the most recent existing `HelloE2E…` function, and the test appended at the end of the test file. It must not be inserted mid-file or reorder existing code.
- An existing function or test with the same name: none is expected, since the token is unique per run. If one is found, that is a conflict to surface, not to overwrite.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `pkg/greeting/greeting.go` MUST define an exported function `HelloE2E202610301951096404` with signature `func HelloE2E202610301951096404() string`.
- **FR-002**: The function MUST return exactly the string `e2e-cross-repo-spawn-20261030-195109-6404`.
- **FR-003**: The function MUST be appended after the most recent existing `HelloE2E…` function, and carry a doc comment in the same style as its neighbours, for example `// HelloE2E202610301951096404 returns the sentinel string for the cross-repo spawn regression test.`
- **FR-004**: `pkg/greeting/greeting_test.go` MUST gain a test `TestHelloE2E202610301951096404`, appended in the existing per-function style. It compares the function's result with the sentinel string and fails with a message showing the got and want values.
- **FR-005**: No other function, test or file may be modified, and the greeting package must not be refactored.
- **FR-006**: The change MUST be merged to the default branch (`main`), and this issue closed, so that alpha issue #10234 is unblocked.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `pkg/greeting/greeting.go` on the default branch contains both `HelloE2E202610301951096404` and `e2e-cross-repo-spawn-20261030-195109-6404`.
- **SC-002**: `go build ./...` and `go test ./...` pass with the new test included.
- **SC-003**: The diff against `main` touches only the two greeting files, plus the engine's spec projection for this issue under `specs/`. It adds one function and one test and changes no existing code.
- **SC-004**: Alpha issue #10234 can proceed once this issue is closed.

## Assumptions

- The `pkg/greeting` package and both files already exist (verified at specification time), so nothing needs creating. The module path is `github.com/handarbeit/fabrik-test-beta` and the package name is `greeting`.
- The most recent existing `HelloE2E…` function is the last of its kind in `greeting.go`, and "append after it" means placing the new function at the end of that group, following the file's current convention.
- The issue's own `Closes #653` PR linkage applies as usual. Linkage to the parent alpha issue is a cross-repo mention only, not a closing keyword.

## Out of Scope

- Any engine behaviour change.
- Any other greeting functions or changes to existing ones.
- Work in other repositories, including the alpha repository's `main.go` call site.

## Source References

- Parent issue: `handarbeit/fabrik-test-alpha#10234`
- Related Fabrik work: #803 (on-demand spawn-target init), #1308 (per-run-unique fixture token)
