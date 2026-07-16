# Feature Specification: Add HelloIssue3496 sentinel to pkg/greeting for cross-repo spawn regression test

**Feature Branch**: `fabrik/issue-83`
**Created**: 2026-07-16
**Status**: Draft
**Input**: User description: "Add HelloIssue3496 sentinel to pkg/greeting (e2e cross-repo spawn regression #3496)"

## Background

This is a sub-issue spawned by Fabrik's Plan stage for `handarbeit/fabrik-test-alpha#3496`, a regression test of Fabrik's cross-repo decomposition and on-demand spawn-target repo init (Fabrik issue #803). It follows the exact pattern of four prior runs already merged into this repo: `HelloE2E`, `HelloIssue31`, `HelloIssue52`, and `HelloIssue3472`.

Alpha's Implement stage for `#3496` is blocked on this sub-issue reaching Done. Once merged, alpha will bump its `go.mod`/`go.sum` to reference this commit and add a call to `HelloIssue3496()` in `main.go`.

This change is purely additive. The existing functions `GreetingFor(name string) string`, `HelloE2E() string`, `HelloIssue31() string`, `HelloIssue52() string`, `HelloIssue3472() string`, and `HelloIssue3367() string` in `pkg/greeting/greeting.go` must remain completely unchanged. `go.mod` must remain pinned to `go 1.22` with no new dependencies.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Add HelloIssue3496 function (Priority: P1)

A consumer of `pkg/greeting` (the alpha repo) calls `HelloIssue3496()` with no arguments and receives the sentinel string `"e2e-cross-repo-spawn-3496"` that proves the cross-repo spawn pipeline ran end-to-end correctly for alpha issue #3496.

**Why this priority**: This is the entire scope of the issue — without this function, the alpha repo cannot resolve the dependency it needs and the E2E regression test cannot run.

**Independent Test**: Run `go test ./pkg/greeting/...` after the change; all existing tests must still pass, and at least one new test must assert `HelloIssue3496() == "e2e-cross-repo-spawn-3496"`.

**Acceptance Scenarios**:

1. **Given** `pkg/greeting` is imported, **When** `greeting.HelloIssue3496()` is called, **Then** it returns exactly `"e2e-cross-repo-spawn-3496"` (no trailing whitespace, no newline)
2. **Given** `pkg/greeting` is imported, **When** `greeting.GreetingFor("Alice")` is called, **Then** it still returns `"Hello, Alice, from fabrik-test-beta"` (existing function unchanged)
3. **Given** `pkg/greeting` is imported, **When** `greeting.HelloE2E()` is called, **Then** it still returns `"e2e-cross-repo-spawn"` (existing function unchanged)
4. **Given** `pkg/greeting` is imported, **When** `greeting.HelloIssue31()` is called, **Then** it still returns `"e2e-cross-repo-spawn-31"` (existing function unchanged)
5. **Given** `pkg/greeting` is imported, **When** `greeting.HelloIssue52()` is called, **Then** it still returns `"e2e-cross-repo-spawn-52"` (existing function unchanged)
6. **Given** `pkg/greeting` is imported, **When** `greeting.HelloIssue3367()` is called, **Then** it still returns `"e2e-cross-repo-spawn-3367"` (existing function unchanged)
7. **Given** `pkg/greeting` is imported, **When** `greeting.HelloIssue3472()` is called, **Then** it still returns `"e2e-cross-repo-spawn-3472"` (existing function unchanged)
8. **Given** the codebase, **When** `go build ./...` is run, **Then** it exits 0
9. **Given** the codebase, **When** `go test ./...` is run, **Then** it exits 0 with at least one test covering `HelloIssue3496`

---

### Edge Cases

- `HelloIssue3496()` takes no parameters — callers must not pass any arguments
- The return value must be byte-exact: `"e2e-cross-repo-spawn-3496"` with no trailing whitespace or newline
- No tagging or semver release step is required for this issue — alpha references beta by commit, not by tag (unlike prior sub-issues #37 and #59, which required a pinned semver tag)

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `pkg/greeting/greeting.go` MUST export `HelloIssue3496() string` returning exactly `"e2e-cross-repo-spawn-3496"` — no parameters, no other return values
- **FR-002**: `HelloIssue3496` MUST be added as a new function in `greeting.go`, appended after the last existing `HelloIssueNN` function (`HelloIssue3472`), following the same one-line doc comment style — `GreetingFor`, `HelloE2E`, `HelloIssue31`, `HelloIssue52`, `HelloIssue3367`, and `HelloIssue3472` are left completely untouched
- **FR-003**: `pkg/greeting/greeting_test.go` MUST include a new `TestHelloIssue3496` function following the same one-test-per-function template as the existing tests (`const want = "e2e-cross-repo-spawn-3496"; if got := HelloIssue3496(); got != want { t.Errorf(...) }`)
- **FR-004**: `go.mod` MUST remain unchanged (`module github.com/handarbeit/fabrik-test-beta`, `go 1.22`, no new dependencies, no toolchain directive bump)
- **FR-005**: `go build ./...` and `go test ./...` MUST both exit 0 before the PR is marked ready
- **FR-006**: No CI workflow changes are needed — `.github/workflows/ci.yml` already runs build+test on push/PR to `main`
- **FR-007**: No documentation changes are needed

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `greeting.HelloIssue3496()` returns exactly `"e2e-cross-repo-spawn-3496"`
- **SC-002**: `greeting.GreetingFor("Alice")` still returns `"Hello, Alice, from fabrik-test-beta"` (no regression)
- **SC-003**: `greeting.HelloE2E()` still returns `"e2e-cross-repo-spawn"` (no regression)
- **SC-004**: `greeting.HelloIssue31()` still returns `"e2e-cross-repo-spawn-31"` (no regression)
- **SC-005**: `greeting.HelloIssue52()` still returns `"e2e-cross-repo-spawn-52"` (no regression)
- **SC-006**: `greeting.HelloIssue3367()` still returns `"e2e-cross-repo-spawn-3367"` (no regression)
- **SC-007**: `greeting.HelloIssue3472()` still returns `"e2e-cross-repo-spawn-3472"` (no regression)
- **SC-008**: `go build ./...` exits 0
- **SC-009**: `go test ./...` exits 0 with at least one test covering `HelloIssue3496`
- **SC-010**: `go.mod` is unchanged (still `go 1.22`, no new dependencies)

## Assumptions

- `GreetingFor`, `HelloE2E`, `HelloIssue31`, `HelloIssue52`, `HelloIssue3367`, and `HelloIssue3472` have external consumers (or are treated as if they do) and must not be modified
- No `go.sum` file changes are needed (no external dependencies added)
- CI (`.github/workflows/ci.yml`) runs `go build ./...` and `go test ./...` — no CI changes are needed
- Alpha will reference this repo by commit (not by semver tag) for issue #3496, so no tagging step is in scope for this sub-issue

## Out of Scope

- Any changes to `GreetingFor`, `HelloE2E`, `HelloIssue31`, `HelloIssue52`, `HelloIssue3367`, or `HelloIssue3472`
- Bumping `go.mod` version or adding dependencies
- Any changes to CI configuration
- Creating or pushing a semver tag
- Any changes to `fabrik-test-alpha`

## Source References

- `pkg/greeting/greeting.go` — file to add `HelloIssue3496` to
- `pkg/greeting/greeting_test.go` — file to add test coverage to
- `go.mod` — module declaration (no changes needed)
- Parent issue: `handarbeit/fabrik-test-alpha#3496`
- Prior art: `specs/59-add-helloissue3472-to-pkg/spec.md`, `specs/37-add-helloissue3367-to-pkg/spec.md`, and the earlier `HelloIssue31`/`HelloIssue52` sub-issues — same pattern, repeated for each new regression run
