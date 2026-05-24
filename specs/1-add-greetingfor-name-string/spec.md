# Feature Specification: Add GreetingFor(name string) to pkg/greeting and tag v0.1.0

**Feature Branch**: `fabrik/issue-1`
**Created**: 2026-05-24
**Status**: Draft
**Input**: User description: "Add GreetingFor(name string) to pkg/greeting and tag v0.1.0"

## Background

This is a child sub-issue of `handarbeit/fabrik-test-alpha` issue #1 (Bootstrap: wire fabrik-test-alpha to fabrik-test-beta end-to-end). The parent issue implements Fabrik's cross-repo decomposition smoke test for v0.0.66. This child handles the beta side of the wiring, which must land and be tagged **before** alpha can import the new function.

`fabrik-test-beta` currently exports `Greeting() string` in `pkg/greeting/greeting.go`, returning `"Hello from fabrik-test-beta"`. This function has no external consumers (no tags exist; no other repo imports it). It must be replaced with `GreetingFor(name string) string` returning `"Hello, <name>, from fabrik-test-beta"`.

After the change is merged to `main`, the Implement stage must create a semver tag `v0.1.0` on the merge commit. Alpha's `go.mod` will reference `github.com/handarbeit/fabrik-test-beta v0.1.0` — without this tag, alpha cannot run `go get`.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Replace Greeting with GreetingFor (Priority: P1)

A consumer of `pkg/greeting` (the alpha repo, or any caller) calls `GreetingFor(name)` with a name string and receives a personalized greeting. The old `Greeting()` function is removed since it has no consumers and is being superseded.

**Why this priority**: This is the entire scope of the issue — without this change, alpha cannot import and use the new function. The tag depends on this landing first.

**Independent Test**: Run `go test ./pkg/greeting/...` after the change; all tests must pass including at least one for empty name and one for a non-empty name.

**Acceptance Scenarios**:

1. **Given** `pkg/greeting` is imported, **When** `greeting.GreetingFor("Alice")` is called, **Then** it returns `"Hello, Alice, from fabrik-test-beta"`
2. **Given** `pkg/greeting` is imported, **When** `greeting.GreetingFor("")` is called, **Then** it returns `"Hello, , from fabrik-test-beta"` (empty name passed through literally, no default substitution)
3. **Given** the codebase, **When** `go build ./...` is run, **Then** it exits 0
4. **Given** the codebase, **When** `go test ./...` is run, **Then** it exits 0 with at least two test cases passing

---

### User Story 2 - Tag v0.1.0 on merge commit (Priority: P1)

After the PR is merged to `main`, a semver tag `v0.1.0` is created on the merge commit and pushed to `origin`, enabling the alpha repo to reference `github.com/handarbeit/fabrik-test-beta v0.1.0` in its `go.mod`.

**Why this priority**: Without the tag, `go get github.com/handarbeit/fabrik-test-beta@v0.1.0` in the alpha repo will fail. The tag is a hard dependency for the cross-repo smoke test.

**Independent Test**: Run `git tag -l v0.1.0` on the remote after merge; the tag must exist and point to the merge commit SHA.

**Acceptance Scenarios**:

1. **Given** the PR has been merged to `main`, **When** `git tag v0.1.0 <merge-sha> && git push origin v0.1.0` is run, **Then** `v0.1.0` exists on the remote pointing to the merge commit

---

### Edge Cases

- `GreetingFor("")` must return `"Hello, , from fabrik-test-beta"` — the empty string is passed through literally with no substitution or default value

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `pkg/greeting/greeting.go` MUST export `GreetingFor(name string) string` returning `"Hello, " + name + ", from fabrik-test-beta"` (exact string format with commas and spaces as shown)
- **FR-002**: `Greeting() string` MUST be removed entirely — it is replaced by `GreetingFor`, and no external consumers exist
- **FR-003**: `pkg/greeting/greeting_test.go` MUST include at least two test cases: `GreetingFor("")` → `"Hello, , from fabrik-test-beta"` and `GreetingFor("Alice")` → `"Hello, Alice, from fabrik-test-beta"`
- **FR-004**: The package import path `github.com/handarbeit/fabrik-test-beta/pkg/greeting` MUST remain unchanged
- **FR-005**: `go.mod` MUST remain pinned to `go 1.22` — do not bump
- **FR-006**: `go build ./...` and `go test ./...` MUST exit 0
- **FR-007**: After the PR is merged to `main`, create git tag `v0.1.0` on the merge commit and push it (`git tag v0.1.0 <sha> && git push origin v0.1.0`)

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `greeting.GreetingFor("Alice")` returns exactly `"Hello, Alice, from fabrik-test-beta"`
- **SC-002**: `greeting.GreetingFor("")` returns exactly `"Hello, , from fabrik-test-beta"`
- **SC-003**: `go build ./...` exits 0
- **SC-004**: `go test ./...` exits 0 with at least two named test cases passing
- **SC-005**: `git tag v0.1.0` exists on the remote `origin` pointing to the merge commit

## Assumptions

- `Greeting() string` has no external consumers — confirmed by the issue (no tags exist, no other repo imports it), so removing it is safe
- The `go.mod` module path `github.com/handarbeit/fabrik-test-beta` and Go version `go 1.22` remain unchanged
- No `go.sum` file is needed (no external dependencies)
- CI (`.github/workflows/ci.yml`) runs `go build ./...` and `go test ./...` — no CI changes are needed
- The tagging step (FR-007) is performed by the Implement stage after the PR is merged, not as part of the PR branch itself

## Out of Scope

- Any changes beyond replacing `Greeting()` with `GreetingFor(name string)`
- Adding flags, middleware, or additional package files
- Bumping `go.mod` version
- Any changes to `fabrik-test-alpha`

## Source References

- `pkg/greeting/greeting.go` — current implementation to replace
- `pkg/greeting/greeting_test.go` — current tests to rewrite
- `go.mod` — module declaration (no changes needed)
- `.github/workflows/ci.yml` — CI config (no changes needed)
- Parent issue: `handarbeit/fabrik-test-alpha#1`
