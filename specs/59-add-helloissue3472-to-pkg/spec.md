# Feature Specification: Add HelloIssue3472() to pkg/greeting for cross-repo spawn regression test

**Feature Branch**: `fabrik/issue-59`
**Created**: 2026-07-14
**Status**: Draft
**Input**: User description: "Add HelloIssue3472 to pkg/greeting (cross-repo spawn for fabrik-test-alpha#3472)"

## Background

This is a sub-issue spawned by Fabrik's cross-repo decomposition feature from `handarbeit/fabrik-test-alpha` issue #3472, a regression run exercising the same e2e scenario as prior issues #8, #31, #52, and #3367. Alpha cannot proceed with its own Implement stage until a new function `HelloIssue3472()` is exported from `fabrik-test-beta` at a resolvable semver version. This sub-issue must reach Done before alpha's Implement stage proceeds.

This change is purely additive. The existing functions `GreetingFor(name string) string`, `HelloE2E() string`, `HelloIssue31() string`, `HelloIssue52() string`, and `HelloIssue3367() string` in `pkg/greeting/greeting.go` must remain completely unchanged. `HelloIssue3367` was added by an unrelated concurrent issue and is out of scope here. `go.mod` must remain pinned to `go 1.22` with no new dependencies.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Add HelloIssue3472 function (Priority: P1)

A consumer of `pkg/greeting` (the alpha repo) calls `HelloIssue3472()` with no arguments and receives the sentinel string `"e2e-cross-repo-spawn-3472"` that proves the cross-repo spawn pipeline ran end-to-end correctly for alpha issue #3472.

**Why this priority**: This is the entire scope of the issue — without this function, the alpha repo cannot resolve the dependency it needs and the E2E regression test cannot run.

**Independent Test**: Run `go test ./pkg/greeting/...` after the change; all existing tests must still pass, and at least one new test must assert `HelloIssue3472() == "e2e-cross-repo-spawn-3472"`.

**Acceptance Scenarios**:

1. **Given** `pkg/greeting` is imported, **When** `greeting.HelloIssue3472()` is called, **Then** it returns exactly `"e2e-cross-repo-spawn-3472"` (no trailing whitespace, no newline)
2. **Given** `pkg/greeting` is imported, **When** `greeting.GreetingFor("Alice")` is called, **Then** it still returns `"Hello, Alice, from fabrik-test-beta"` (existing function unchanged)
3. **Given** `pkg/greeting` is imported, **When** `greeting.HelloE2E()` is called, **Then** it still returns `"e2e-cross-repo-spawn"` (existing function unchanged)
4. **Given** `pkg/greeting` is imported, **When** `greeting.HelloIssue31()` is called, **Then** it still returns `"e2e-cross-repo-spawn-31"` (existing function unchanged)
5. **Given** `pkg/greeting` is imported, **When** `greeting.HelloIssue52()` is called, **Then** it still returns `"e2e-cross-repo-spawn-52"` (existing function unchanged)
6. **Given** `pkg/greeting` is imported, **When** `greeting.HelloIssue3367()` is called, **Then** it still returns `"e2e-cross-repo-spawn-3367"` (existing function unchanged, added by an unrelated concurrent issue)
7. **Given** the codebase, **When** `go build ./...` is run, **Then** it exits 0
8. **Given** the codebase, **When** `go test ./...` is run, **Then** it exits 0 with at least one test covering `HelloIssue3472`

---

### User Story 2 - Tag next semver version on merge (Priority: P1)

After the PR is merged to `main`, a new semver tag is created on the merge commit and pushed to `origin`, enabling the alpha repo to resolve a `fabrik-test-beta` version that includes `HelloIssue3472`. The exact version number is determined at merge time by inspecting existing tags — it is not fixed in advance, since other concurrent issues in this repo may also be cutting tags in the same window.

**Why this priority**: Without a new tag, `go get` in alpha's Implement stage will not resolve the new function. The tag is a hard dependency for the cross-repo E2E test.

**Independent Test**: Run `git tag -l | sort -V | tail -1` before and after the tagging step; the new tag must be strictly greater than the pre-existing highest tag and its commit must be reachable from `main`.

**Acceptance Scenarios**:

1. **Given** the PR has been merged to `main`, **When** the highest existing tag is determined via `git tag -l | sort -V | tail -1`, **Then** the next tag is computed as that version with the patch (or minor, per existing convention) component incremented
2. **Given** the next tag has been computed, **When** `git tag <next-version> <merge-commit> && git push origin <next-version>` is run, **Then** the new tag exists on the remote pointing to the merge commit

---

### Edge Cases

- `HelloIssue3472()` takes no parameters — callers must not pass any arguments
- The return value must be byte-exact: `"e2e-cross-repo-spawn-3472"` with no trailing whitespace or newline
- `HelloIssue3367` may already exist on `main` from a concurrent, unrelated issue by the time this branch is implemented or merged — it must be preserved untouched, and its presence must not be treated as a merge conflict to resolve by removal
- Tag numbering is not assumed in advance; the actual next tag must be computed by inspecting `git tag -l` at merge time, since other concurrent issues may tag `main` first

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `pkg/greeting/greeting.go` MUST export `HelloIssue3472() string` returning exactly `"e2e-cross-repo-spawn-3472"` — no parameters, no other return values
- **FR-002**: `HelloIssue3472` MUST be added as a new function in `greeting.go`, following the one-line doc comment style of `HelloIssue31`/`HelloIssue52` — `GreetingFor`, `HelloE2E`, `HelloIssue31`, `HelloIssue52`, and `HelloIssue3367` are left completely untouched
- **FR-003**: `pkg/greeting/greeting_test.go` MUST include a new `TestHelloIssue3472` function following the same one-test-per-function pattern as the existing tests, asserting `HelloIssue3472() == "e2e-cross-repo-spawn-3472"`
- **FR-004**: `go.mod` MUST remain unchanged (`module github.com/handarbeit/fabrik-test-beta`, `go 1.22`, no new dependencies, no toolchain directive bump)
- **FR-005**: `go build ./...` and `go test ./...` MUST both exit 0 before the PR is marked ready
- **FR-006**: After the PR is merged to `main`, determine the next semver tag by inspecting existing tags at merge time (do not hardcode a version number) and create + push that tag on the merge commit — this MUST occur before the corresponding pipeline stage signals completion

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `greeting.HelloIssue3472()` returns exactly `"e2e-cross-repo-spawn-3472"`
- **SC-002**: `greeting.GreetingFor("Alice")` still returns `"Hello, Alice, from fabrik-test-beta"` (no regression)
- **SC-003**: `greeting.HelloE2E()` still returns `"e2e-cross-repo-spawn"` (no regression)
- **SC-004**: `greeting.HelloIssue31()` still returns `"e2e-cross-repo-spawn-31"` (no regression)
- **SC-005**: `greeting.HelloIssue52()` still returns `"e2e-cross-repo-spawn-52"` (no regression)
- **SC-006**: `greeting.HelloIssue3367()` still returns `"e2e-cross-repo-spawn-3367"` (no regression)
- **SC-007**: `go build ./...` exits 0
- **SC-008**: `go test ./...` exits 0 with at least one test covering `HelloIssue3472`
- **SC-009**: `git ls-remote --tags origin` shows a new tag, strictly higher than the pre-merge highest tag, pointing at the merge commit
- **SC-010**: `go.mod` is unchanged (still `go 1.22`, no new dependencies)

## Assumptions

- `GreetingFor`, `HelloE2E`, `HelloIssue31`, `HelloIssue52`, and `HelloIssue3367` have external consumers (or are treated as if they do) and must not be modified
- No `go.sum` file is needed (no external dependencies added)
- CI (`.github/workflows/ci.yml`) runs `go build ./...` and `go test ./...` — no CI changes are needed
- The tagging step (FR-006) is performed after the PR is merged, not as part of the PR branch itself
- At spec time, the highest existing tag on `main` is `v0.5.0`; the actual next tag used at merge time may differ if other concurrent issues tag first, per the issue's explicit instruction

## Out of Scope

- Any changes to `GreetingFor`, `HelloE2E`, `HelloIssue31`, `HelloIssue52`, or `HelloIssue3367`
- Bumping `go.mod` version or adding dependencies
- Any changes to CI configuration
- Any changes to `fabrik-test-alpha`

## Source References

- `pkg/greeting/greeting.go` — file to add `HelloIssue3472` to
- `pkg/greeting/greeting_test.go` — file to add test coverage to
- `go.mod` — module declaration (no changes needed)
- Parent issue: `handarbeit/fabrik-test-alpha#3472`
- Prior art: `specs/37-add-helloissue3367-to-pkg/spec.md` (same pattern; note that spec hardcoded `v0.5.0` since it was the only concurrent tagger — this issue explicitly avoids hardcoding since concurrency is expected)
