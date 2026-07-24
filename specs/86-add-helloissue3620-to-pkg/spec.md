# Feature Specification: Add HelloIssue3620 sentinel to pkg/greeting for cross-repo spawn regression test

**Feature Branch**: `fabrik/issue-86`
**Created**: 2026-07-24
**Status**: Draft
**Input**: User description: "e2e cross-repo spawn member for #3620: add HelloIssue3620"

## Background

This is a sub-issue spawned by Fabrik's Plan stage for `handarbeit/fabrik-test-alpha#3620`, a regression run of Fabrik's cross-repo decomposition capability. It follows the exact pattern of prior sibling functions `HelloE2E`, `HelloIssue31`, `HelloIssue52`, `HelloIssue3367`, `HelloIssue3472`, and `HelloIssue3496`, all already merged in this repo's `pkg/greeting/greeting.go`.

This change is purely additive. The existing functions `GreetingFor(name string) string`, `HelloE2E() string`, `HelloIssue31() string`, `HelloIssue52() string`, `HelloIssue3367() string`, `HelloIssue3472() string`, and `HelloIssue3496() string` in `pkg/greeting/greeting.go` must remain completely unchanged. `go.mod` must remain pinned to `go 1.22` with no new dependencies.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Add HelloIssue3620 function (Priority: P1)

A consumer of `pkg/greeting` (the alpha repo) calls `HelloIssue3620()` with no arguments and receives the sentinel string `"e2e-cross-repo-spawn-3620"` that proves the cross-repo spawn pipeline ran end-to-end correctly for alpha issue #3620.

**Why this priority**: This is the entire scope of the issue — without this function, the alpha repo cannot resolve the dependency it needs and the E2E regression test cannot run.

**Independent Test**: Run `go test ./pkg/greeting/...` after the change; all existing tests must still pass, and a new test must assert `HelloIssue3620() == "e2e-cross-repo-spawn-3620"`.

**Acceptance Scenarios**:

1. **Given** `pkg/greeting` is imported, **When** `greeting.HelloIssue3620()` is called, **Then** it returns exactly `"e2e-cross-repo-spawn-3620"` (no trailing whitespace, no newline)
2. **Given** `pkg/greeting` is imported, **When** `greeting.GreetingFor("Alice")` is called, **Then** it still returns `"Hello, Alice, from fabrik-test-beta"` (existing function unchanged)
3. **Given** `pkg/greeting` is imported, **When** `greeting.HelloE2E()` is called, **Then** it still returns `"e2e-cross-repo-spawn"` (existing function unchanged)
4. **Given** `pkg/greeting` is imported, **When** `greeting.HelloIssue31()` is called, **Then** it still returns `"e2e-cross-repo-spawn-31"` (existing function unchanged)
5. **Given** `pkg/greeting` is imported, **When** `greeting.HelloIssue52()` is called, **Then** it still returns `"e2e-cross-repo-spawn-52"` (existing function unchanged)
6. **Given** `pkg/greeting` is imported, **When** `greeting.HelloIssue3367()` is called, **Then** it still returns `"e2e-cross-repo-spawn-3367"` (existing function unchanged)
7. **Given** `pkg/greeting` is imported, **When** `greeting.HelloIssue3472()` is called, **Then** it still returns `"e2e-cross-repo-spawn-3472"` (existing function unchanged)
8. **Given** `pkg/greeting` is imported, **When** `greeting.HelloIssue3496()` is called, **Then** it still returns `"e2e-cross-repo-spawn-3496"` (existing function unchanged)
9. **Given** the codebase, **When** `go build ./...` is run, **Then** it exits 0
10. **Given** the codebase, **When** `go test ./...` is run, **Then** it exits 0 with a test covering `HelloIssue3620`

---

### Edge Cases

- `HelloIssue3620()` takes no parameters — callers must not pass any arguments
- The return value must be byte-exact: `"e2e-cross-repo-spawn-3620"` with no trailing whitespace or newline

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `pkg/greeting/greeting.go` MUST export `HelloIssue3620() string` returning exactly `"e2e-cross-repo-spawn-3620"` — no parameters, no other return values
- **FR-002**: `HelloIssue3620` MUST be added as a new function in `greeting.go`, appended after the last existing `HelloIssueNN` function (`HelloIssue3496`), following the same one-line doc comment style — `GreetingFor`, `HelloE2E`, `HelloIssue31`, `HelloIssue52`, `HelloIssue3367`, `HelloIssue3472`, and `HelloIssue3496` are left completely untouched
- **FR-003**: `pkg/greeting/greeting_test.go` MUST include a new `TestHelloIssue3620` function following the same one-test-per-function template as the existing tests (`const want = "e2e-cross-repo-spawn-3620"; if got := HelloIssue3620(); got != want { t.Errorf(...) }`)
- **FR-004**: `go.mod` MUST remain unchanged (`module github.com/handarbeit/fabrik-test-beta`, `go 1.22`, no new dependencies, no toolchain directive bump)
- **FR-005**: `go build ./...` and `go test ./...` MUST both exit 0 before the PR is marked ready
- **FR-006**: No CI workflow changes are needed — `.github/workflows/ci.yml` already runs build+test on push/PR to `main`
- **FR-007**: No documentation changes are needed

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `greeting.HelloIssue3620()` returns exactly `"e2e-cross-repo-spawn-3620"`
- **SC-002**: `greeting.GreetingFor("Alice")` still returns `"Hello, Alice, from fabrik-test-beta"` (no regression)
- **SC-003**: `greeting.HelloE2E()` still returns `"e2e-cross-repo-spawn"` (no regression)
- **SC-004**: `greeting.HelloIssue31()` still returns `"e2e-cross-repo-spawn-31"` (no regression)
- **SC-005**: `greeting.HelloIssue52()` still returns `"e2e-cross-repo-spawn-52"` (no regression)
- **SC-006**: `greeting.HelloIssue3367()` still returns `"e2e-cross-repo-spawn-3367"` (no regression)
- **SC-007**: `greeting.HelloIssue3472()` still returns `"e2e-cross-repo-spawn-3472"` (no regression)
- **SC-008**: `greeting.HelloIssue3496()` still returns `"e2e-cross-repo-spawn-3496"` (no regression)
- **SC-009**: `go build ./...` exits 0
- **SC-010**: `go test ./...` exits 0 with a test covering `HelloIssue3620`
- **SC-011**: `go.mod` is unchanged (still `go 1.22`, no new dependencies)

## Assumptions

- `GreetingFor`, `HelloE2E`, `HelloIssue31`, `HelloIssue52`, `HelloIssue3367`, `HelloIssue3472`, and `HelloIssue3496` have external consumers (or are treated as if they do) and must not be modified
- No `go.sum` file changes are needed (no external dependencies added)
- CI (`.github/workflows/ci.yml`) runs `go build ./...` and `go test ./...` — no CI changes are needed

## Out of Scope

- Any changes to `GreetingFor`, `HelloE2E`, `HelloIssue31`, `HelloIssue52`, `HelloIssue3367`, `HelloIssue3472`, or `HelloIssue3496`
- Bumping `go.mod` version or adding dependencies
- Any changes to CI configuration
- Any changes to `fabrik-test-alpha`

## Source References

- `pkg/greeting/greeting.go` — file to add `HelloIssue3620` to
- `pkg/greeting/greeting_test.go` — file to add test coverage to
- `go.mod` — module declaration (no changes needed)
- Parent issue: `handarbeit/fabrik-test-alpha#3620`
- Prior art: `specs/83-add-helloissue3496-sentinel-to/spec.md`, `specs/59-add-helloissue3472-to-pkg/spec.md`, and the earlier `HelloIssue31`/`HelloIssue52`/`HelloIssue3367` sub-issues — same pattern, repeated for each new regression run
