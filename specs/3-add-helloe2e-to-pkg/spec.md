# Feature Specification: Add HelloE2E() to pkg/greeting for cross-repo spawn regression test

**Feature Branch**: `fabrik/issue-3`
**Created**: 2026-05-24
**Status**: Draft
**Input**: User description: "Add HelloE2E() to pkg/greeting for cross-repo spawn regression test"

## Background

This is a sub-issue spawned by `handarbeit/fabrik-test-alpha` issue #8 (E2E cross-repo spawn). Alpha calls a new function `HelloE2E()` from `pkg/greeting`; that function must exist in a tagged release of `fabrik-test-beta` before alpha's Implement stage begins. This sub-issue must reach Done before alpha proceeds.

This change validates Fabrik issue #803 (on-demand spawn-target repo init). The existing `GreetingFor(name string) string` function must not be modified in any way. `go.mod` must remain unchanged (`go 1.22`, no external dependencies).

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Add HelloE2E function (Priority: P1)

A consumer of `pkg/greeting` (the alpha repo) calls `HelloE2E()` with no arguments and receives the sentinel string `"e2e-cross-repo-spawn"` that proves the cross-repo spawn pipeline ran end-to-end correctly.

**Why this priority**: This is the entire scope of the issue — without this function, the alpha repo cannot compile and the E2E regression test cannot run.

**Independent Test**: Run `go test ./pkg/greeting/...` after the change; all existing tests must still pass, and at least one new test must assert `HelloE2E() == "e2e-cross-repo-spawn"`.

**Acceptance Scenarios**:

1. **Given** `pkg/greeting` is imported, **When** `greeting.HelloE2E()` is called, **Then** it returns exactly `"e2e-cross-repo-spawn"` (no trailing whitespace, no newline)
2. **Given** `pkg/greeting` is imported, **When** `greeting.GreetingFor("Alice")` is called, **Then** it still returns `"Hello, Alice, from fabrik-test-beta"` (existing function unchanged)
3. **Given** the codebase, **When** `go build ./...` is run, **Then** it exits 0
4. **Given** the codebase, **When** `go test ./...` is run, **Then** it exits 0 with at least one test covering `HelloE2E`

---

### User Story 2 - Tag v0.2.0 on merge commit (Priority: P1)

After the PR is merged to `main`, a semver tag `v0.2.0` is created on the merge commit and pushed to `origin`, enabling the alpha repo to reference `github.com/handarbeit/fabrik-test-beta v0.2.0` in its `go.mod`.

**Why this priority**: Without the tag, `go mod tidy` in alpha will fail when it tries to resolve `v0.2.0`. The tag is a hard dependency for the cross-repo E2E test.

**Independent Test**: Run `git ls-remote --tags origin` after merge; `refs/tags/v0.2.0` must appear.

**Acceptance Scenarios**:

1. **Given** the PR has been merged to `main`, **When** `git fetch origin main && git tag v0.2.0 origin/main && git push origin v0.2.0` is run, **Then** `refs/tags/v0.2.0` exists on the remote pointing to the merge commit

---

### Edge Cases

- `HelloE2E()` takes no parameters — callers must not pass any arguments
- The return value must be byte-exact: `"e2e-cross-repo-spawn"` with no trailing whitespace or newline

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `pkg/greeting/greeting.go` MUST export `HelloE2E() string` returning exactly `"e2e-cross-repo-spawn"` — no parameters, no other return values
- **FR-002**: `HelloE2E` MUST be added as a second function in `greeting.go` — `GreetingFor` is left completely untouched
- **FR-003**: `pkg/greeting/greeting_test.go` MUST include at least one test asserting `greeting.HelloE2E() == "e2e-cross-repo-spawn"` (either a new table row in `TestGreetingFor` or a standalone `TestHelloE2E`)
- **FR-004**: `go.mod` MUST remain unchanged (`module github.com/handarbeit/fabrik-test-beta`, `go 1.22`, no new dependencies)
- **FR-005**: `go build ./...` and `go test ./...` MUST both exit 0 before the PR is marked ready
- **FR-006**: After the PR is merged to `main`, create git tag `v0.2.0` on the merge commit and push it to `origin`

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `greeting.HelloE2E()` returns exactly `"e2e-cross-repo-spawn"`
- **SC-002**: `greeting.GreetingFor("Alice")` still returns `"Hello, Alice, from fabrik-test-beta"` (no regression)
- **SC-003**: `go build ./...` exits 0
- **SC-004**: `go test ./...` exits 0 with at least one test covering `HelloE2E`
- **SC-005**: `git ls-remote --tags origin` shows `refs/tags/v0.2.0` after the PR merges
- **SC-006**: `go.mod` is unchanged (still `go 1.22`, no new dependencies)

## Assumptions

- `GreetingFor(name string) string` has external consumers in alpha and must not be modified
- No `go.sum` file is needed (no external dependencies added)
- CI (`.github/workflows/ci.yml`) runs `go build ./...` and `go test ./...` — no CI changes are needed; `GOPRIVATE` is not set and must not be required
- The tagging step (FR-006) is performed by the Implement or Validate stage after the PR is merged, not as part of the PR branch itself

## Out of Scope

- Any changes to `GreetingFor` or any other existing function
- Bumping `go.mod` version or adding dependencies
- Any changes to CI configuration
- Any changes to `fabrik-test-alpha`

## Source References

- `pkg/greeting/greeting.go` — file to add `HelloE2E` to
- `pkg/greeting/greeting_test.go` — file to add test coverage to
- `go.mod` — module declaration (no changes needed)
- Parent issue: `handarbeit/fabrik-test-alpha#8`
- Validates: Fabrik issue #803 (on-demand spawn-target repo init)
