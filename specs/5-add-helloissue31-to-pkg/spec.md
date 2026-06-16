# Feature Specification: Add HelloIssue31() to pkg/greeting for cross-repo spawn regression test

**Feature Branch**: `fabrik/issue-5`
**Created**: 2026-06-16
**Status**: Draft
**Input**: User description: "Add HelloIssue31() to pkg/greeting (cross-repo spawn for alpha issue #31)"

## Background

This is a sub-issue spawned by `handarbeit/fabrik-test-alpha` issue #31, which exercises Fabrik's cross-repo decomposition path (regression test for Fabrik issue #803). Alpha cannot compile or pass tests until a new function `HelloIssue31()` is exported from `fabrik-test-beta` at a pinned semver version `v0.3.0`. This sub-issue must reach Done before alpha's Implement stage proceeds.

This change is purely additive. The existing functions `GreetingFor(name string) string` and `HelloE2E() string` in `pkg/greeting/greeting.go` must remain completely unchanged. `go.mod` must remain pinned to `go 1.22` with no new dependencies.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Add HelloIssue31 function (Priority: P1)

A consumer of `pkg/greeting` (the alpha repo) calls `HelloIssue31()` with no arguments and receives the sentinel string `"e2e-cross-repo-spawn-31"` that proves the cross-repo spawn pipeline ran end-to-end correctly for alpha issue #31.

**Why this priority**: This is the entire scope of the issue — without this function, the alpha repo cannot compile and the E2E regression test cannot run.

**Independent Test**: Run `go test ./pkg/greeting/...` after the change; all existing tests must still pass, and at least one new test must assert `HelloIssue31() == "e2e-cross-repo-spawn-31"`.

**Acceptance Scenarios**:

1. **Given** `pkg/greeting` is imported, **When** `greeting.HelloIssue31()` is called, **Then** it returns exactly `"e2e-cross-repo-spawn-31"` (no trailing whitespace, no newline)
2. **Given** `pkg/greeting` is imported, **When** `greeting.GreetingFor("Alice")` is called, **Then** it still returns `"Hello, Alice, from fabrik-test-beta"` (existing function unchanged)
3. **Given** `pkg/greeting` is imported, **When** `greeting.HelloE2E()` is called, **Then** it still returns `"e2e-cross-repo-spawn"` (existing function unchanged)
4. **Given** the codebase, **When** `go build ./...` is run, **Then** it exits 0
5. **Given** the codebase, **When** `go test ./...` is run, **Then** it exits 0 with at least one test covering `HelloIssue31`

---

### User Story 2 - Tag v0.3.0 on merge commit (Priority: P1)

After the PR is merged to `main`, a semver tag `v0.3.0` is created on the merge commit and pushed to `origin`, enabling the alpha repo to reference `github.com/handarbeit/fabrik-test-beta v0.3.0` in its `go.mod`.

**Why this priority**: Without the tag, `go get github.com/handarbeit/fabrik-test-beta@v0.3.0` in alpha's Implement stage will fail. The tag is a hard dependency for the cross-repo E2E test.

**Independent Test**: Run `git ls-remote --tags origin` after merge; `refs/tags/v0.3.0` must appear.

**Acceptance Scenarios**:

1. **Given** the PR has been merged to `main`, **When** `git fetch origin main && git checkout main && git pull origin main && git tag v0.3.0 && git push origin v0.3.0` is run, **Then** `refs/tags/v0.3.0` exists on the remote pointing to the merge commit

---

### Edge Cases

- `HelloIssue31()` takes no parameters — callers must not pass any arguments
- The return value must be byte-exact: `"e2e-cross-repo-spawn-31"` with no trailing whitespace or newline
- Tags on this repo are sequential: `v0.1.0` (GreetingFor), `v0.2.0` (HelloE2E) — the next tag MUST be `v0.3.0`, not any other version

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `pkg/greeting/greeting.go` MUST export `HelloIssue31() string` returning exactly `"e2e-cross-repo-spawn-31"` — no parameters, no other return values
- **FR-002**: `HelloIssue31` MUST be added as a new function in `greeting.go` — `GreetingFor` and `HelloE2E` are left completely untouched
- **FR-003**: `pkg/greeting/greeting_test.go` MUST include a new `TestHelloIssue31` function following the same pattern as `TestHelloE2E`, asserting `HelloIssue31() == "e2e-cross-repo-spawn-31"`
- **FR-004**: `go.mod` MUST remain unchanged (`module github.com/handarbeit/fabrik-test-beta`, `go 1.22`, no new dependencies)
- **FR-005**: `go build ./...` and `go test ./...` MUST both exit 0 before the PR is marked ready
- **FR-006**: After the PR is merged to `main`, create git tag `v0.3.0` on the merge commit and push it to `origin` — this MUST occur before `FABRIK_STAGE_COMPLETE` is emitted in the Implement stage

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `greeting.HelloIssue31()` returns exactly `"e2e-cross-repo-spawn-31"`
- **SC-002**: `greeting.GreetingFor("Alice")` still returns `"Hello, Alice, from fabrik-test-beta"` (no regression)
- **SC-003**: `greeting.HelloE2E()` still returns `"e2e-cross-repo-spawn"` (no regression)
- **SC-004**: `go build ./...` exits 0
- **SC-005**: `go test ./...` exits 0 with at least one test covering `HelloIssue31`
- **SC-006**: `git ls-remote --tags origin` shows `refs/tags/v0.3.0` after the PR merges
- **SC-007**: `go.mod` is unchanged (still `go 1.22`, no new dependencies)

## Assumptions

- `GreetingFor(name string) string` and `HelloE2E() string` have external consumers in alpha and must not be modified
- No `go.sum` file is needed (no external dependencies added)
- CI (`.github/workflows/ci.yml`) runs `go build ./...` and `go test ./...` — no CI changes are needed
- The tagging step (FR-006) is performed by the Implement stage after the PR is merged, not as part of the PR branch itself

## Out of Scope

- Any changes to `GreetingFor` or `HelloE2E` or any other existing function
- Bumping `go.mod` version or adding dependencies
- Any changes to CI configuration
- Any changes to `fabrik-test-alpha`

## Source References

- `pkg/greeting/greeting.go` — file to add `HelloIssue31` to
- `pkg/greeting/greeting_test.go` — file to add test coverage to
- `go.mod` — module declaration (no changes needed)
- Parent issue: `handarbeit/fabrik-test-alpha#31`
- Validates: Fabrik issue #803 (cross-repo decomposition)
- Prior art: `specs/3-add-helloe2e-to-pkg/spec.md` (same pattern, `v0.2.0` tag)
