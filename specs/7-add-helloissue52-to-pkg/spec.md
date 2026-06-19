# Feature Specification: Add HelloIssue52() to pkg/greeting for cross-repo spawn regression test

**Feature Branch**: `fabrik/issue-7`
**Created**: 2026-06-19
**Status**: Draft
**Input**: User description: "Add HelloIssue52() to pkg/greeting (cross-repo spawn for fabrik-test-alpha issue #52)"

## Background

This is a sub-issue spawned by `handarbeit/fabrik-test-alpha` issue #52, which exercises Fabrik's cross-repo decomposition path (regression test for Fabrik issue #803). Alpha cannot compile or pass tests until a new function `HelloIssue52()` is exported from `fabrik-test-beta` at a pinned semver version `v0.4.0`. This sub-issue must reach Done before alpha's Implement stage proceeds.

This change is purely additive. The existing functions `GreetingFor(name string) string`, `HelloE2E() string`, and `HelloIssue31() string` in `pkg/greeting/greeting.go` must remain completely unchanged. `go.mod` must remain pinned to `go 1.22` with no new dependencies.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Add HelloIssue52 function (Priority: P1)

A consumer of `pkg/greeting` (the alpha repo) calls `HelloIssue52()` with no arguments and receives the sentinel string `"e2e-cross-repo-spawn-52"` that proves the cross-repo spawn pipeline ran end-to-end correctly for alpha issue #52.

**Why this priority**: This is the entire scope of the issue — without this function, the alpha repo cannot compile and the E2E regression test cannot run.

**Independent Test**: Run `go test ./pkg/greeting/...` after the change; all existing tests must still pass, and at least one new test must assert `HelloIssue52() == "e2e-cross-repo-spawn-52"`.

**Acceptance Scenarios**:

1. **Given** `pkg/greeting` is imported, **When** `greeting.HelloIssue52()` is called, **Then** it returns exactly `"e2e-cross-repo-spawn-52"` (no trailing whitespace, no newline)
2. **Given** `pkg/greeting` is imported, **When** `greeting.GreetingFor("Alice")` is called, **Then** it still returns `"Hello, Alice, from fabrik-test-beta"` (existing function unchanged)
3. **Given** `pkg/greeting` is imported, **When** `greeting.HelloE2E()` is called, **Then** it still returns `"e2e-cross-repo-spawn"` (existing function unchanged)
4. **Given** `pkg/greeting` is imported, **When** `greeting.HelloIssue31()` is called, **Then** it still returns `"e2e-cross-repo-spawn-31"` (existing function unchanged)
5. **Given** the codebase, **When** `go build ./...` is run, **Then** it exits 0
6. **Given** the codebase, **When** `go test ./...` is run, **Then** it exits 0 with at least one test covering `HelloIssue52`

---

### User Story 2 - Tag v0.4.0 on merge commit (Priority: P1)

After the PR is merged to `main`, a semver tag `v0.4.0` is created on the merge commit and pushed to `origin`, enabling the alpha repo to resolve `github.com/handarbeit/fabrik-test-beta@latest` to a version that includes `HelloIssue52`.

**Why this priority**: Without the tag, `go get github.com/handarbeit/fabrik-test-beta@latest` in alpha's Implement stage will not resolve the new function. The tag is a hard dependency for the cross-repo E2E test.

**Independent Test**: Run `git ls-remote --tags origin` after merge; `refs/tags/v0.4.0` must appear.

**Acceptance Scenarios**:

1. **Given** the PR has been merged to `main`, **When** `git tag v0.4.0 <merge-commit> && git push origin v0.4.0` is run, **Then** `refs/tags/v0.4.0` exists on the remote pointing to the merge commit

---

### Edge Cases

- `HelloIssue52()` takes no parameters — callers must not pass any arguments
- The return value must be byte-exact: `"e2e-cross-repo-spawn-52"` with no trailing whitespace or newline
- Tags on this repo are sequential: `v0.1.0` (GreetingFor), `v0.2.0` (HelloE2E), `v0.3.0` (HelloIssue31) — the next tag MUST be `v0.4.0`, not any other version

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `pkg/greeting/greeting.go` MUST export `HelloIssue52() string` returning exactly `"e2e-cross-repo-spawn-52"` — no parameters, no other return values
- **FR-002**: `HelloIssue52` MUST be added as a new function in `greeting.go` — `GreetingFor`, `HelloE2E`, and `HelloIssue31` are left completely untouched
- **FR-003**: `pkg/greeting/greeting_test.go` MUST include a new `TestHelloIssue52` function following the same pattern as `TestHelloIssue31`, asserting `HelloIssue52() == "e2e-cross-repo-spawn-52"`
- **FR-004**: `go.mod` MUST remain unchanged (`module github.com/handarbeit/fabrik-test-beta`, `go 1.22`, no new dependencies)
- **FR-005**: `go build ./...` and `go test ./...` MUST both exit 0 before the PR is marked ready
- **FR-006**: After the PR is merged to `main`, create git tag `v0.4.0` on the merge commit and push it to `origin` — this MUST occur before `FABRIK_STAGE_COMPLETE` is emitted in the Implement stage

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `greeting.HelloIssue52()` returns exactly `"e2e-cross-repo-spawn-52"`
- **SC-002**: `greeting.GreetingFor("Alice")` still returns `"Hello, Alice, from fabrik-test-beta"` (no regression)
- **SC-003**: `greeting.HelloE2E()` still returns `"e2e-cross-repo-spawn"` (no regression)
- **SC-004**: `greeting.HelloIssue31()` still returns `"e2e-cross-repo-spawn-31"` (no regression)
- **SC-005**: `go build ./...` exits 0
- **SC-006**: `go test ./...` exits 0 with at least one test covering `HelloIssue52`
- **SC-007**: `git ls-remote --tags origin` shows `refs/tags/v0.4.0` after the PR merges
- **SC-008**: `go.mod` is unchanged (still `go 1.22`, no new dependencies)

## Assumptions

- `GreetingFor(name string) string`, `HelloE2E() string`, and `HelloIssue31() string` have external consumers and must not be modified
- No `go.sum` file is needed (no external dependencies added)
- CI (`.github/workflows/ci.yml`) runs `go build ./...` and `go test ./...` — no CI changes are needed
- The tagging step (FR-006) is performed by the Implement stage after the PR is merged, not as part of the PR branch itself

## Out of Scope

- Any changes to `GreetingFor`, `HelloE2E`, or `HelloIssue31` or any other existing function
- Bumping `go.mod` version or adding dependencies
- Any changes to CI configuration
- Any changes to `fabrik-test-alpha`

## Source References

- `pkg/greeting/greeting.go` — file to add `HelloIssue52` to
- `pkg/greeting/greeting_test.go` — file to add test coverage to
- `go.mod` — module declaration (no changes needed)
- Parent issue: `handarbeit/fabrik-test-alpha#52`
- Validates: Fabrik issue #803 (cross-repo decomposition)
- Prior art: `specs/5-add-helloissue31-to-pkg/spec.md` (same pattern, `v0.3.0` tag)
