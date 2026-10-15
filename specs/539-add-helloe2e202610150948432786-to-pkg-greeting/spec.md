# Feature Specification: Add HelloE2E202610150948432786 to pkg/greeting

**Feature Branch**: `fabrik/issue-539`
**Created**: 2026-10-15
**Status**: Specified
**Input**: User description: "The alpha repo (`handarbeit/fabrik-test-alpha`) needs to call a new greeting function from this repo's `pkg/greeting` package. That function does not exist yet. This change must land on beta `main` first, because alpha cannot build against a function that is missing from the beta version it pins. Requirements: In `pkg/greeting/greeting.go`, append an exported function `HelloE2E202610150948432786() string` that returns exactly the string `\"e2e-cross-repo-spawn-20261015-094843-2786\"`. Follow the existing pattern in the file: a doc comment, then a single return. In `pkg/greeting/greeting_test.go`, append a test `TestHelloE2E202610150948432786` that asserts the function returns exactly `e2e-cross-repo-spawn-20261015-094843-2786`. Follow the existing test pattern: a `const want` string, a `got :=` call, and `t.Errorf(\"HelloE2E202610150948432786() = %q; want %q\", got, want)`. Verbatim constraint: use the function name `HelloE2E202610150948432786` and the string literal `e2e-cross-repo-spawn-20261015-094843-2786` exactly as written. Do not invent a different token or reword anything. The identifier has the digits contiguous (`20261015094843` followed by `2786`), while the literal is hyphen-segmented (`20261015-094843-2786`). A per-run-unique token is baked into both, and the scenario's final assertion depends on this exact text landing in this repo. Scope: only the one function and its one test; no other changes in this repo. Do not add a closing keyword for any other issue; the PR for this issue should carry `Closes` for this issue only. Acceptance: `go test ./pkg/greeting/...` passes, including the new test; `HelloE2E202610150948432786()` returns exactly `e2e-cross-repo-spawn-20261015-094843-2786`; the change is merged to beta `main` so the alpha repo can pin to the merge commit."

## Background

This issue was spawned from a parent issue in the alpha repo (`handarbeit/fabrik-test-alpha#8510`) as part of an end-to-end cross-repo spawn scenario. The alpha repo needs to call a new greeting function from this repo's `pkg/greeting` package, but that function does not exist yet. Because alpha pins a specific beta version, the function must land on beta `main` first; alpha cannot build against a function missing from the beta version it pins.

The `pkg/greeting` package already contains several sentinel functions of the same kind from earlier runs of this scenario (for example `HelloE2E202610150633121095`), each with a doc comment and a single return of a hyphen-segmented literal, and each with a matching test using a `const want` / `got :=` / `t.Errorf` pattern. This change adds one more in the same shape, carrying this run's unique token. The scenario's final assertion depends on this exact text appearing in this repo.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Expose the run-unique sentinel from beta (Priority: P1)

A maintainer of the alpha repo needs a function in beta's `pkg/greeting` that returns this run's sentinel string, so alpha can pin to the beta merge commit and call it.

**Why this priority**: It is the only deliverable. Alpha's work is blocked until it is merged to beta `main`.

**Independent Test**: Run `go test ./pkg/greeting/...` in this repo and confirm the new test passes and asserts the exact literal.

**Acceptance Scenarios**:

1. **Given** beta `main` with the change merged, **When** a caller invokes `HelloE2E202610150948432786()`, **Then** it returns exactly `e2e-cross-repo-spawn-20261015-094843-2786`.
2. **Given** the change on the issue branch, **When** `go test ./pkg/greeting/...` is run, **Then** all tests pass, including `TestHelloE2E202610150948432786`.
3. **Given** the change is merged, **When** the alpha repo pins to the beta merge commit, **Then** it can build against the new function.

---

### Edge Cases

- The identifier has the digits contiguous (`20261015094843` followed by `2786`) while the string literal is hyphen-segmented (`20261015-094843-2786`); the two must not be conflated or "normalised" to match each other.
- Existing sentinel functions and tests must remain unchanged; the new function and test are appended after them.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `pkg/greeting/greeting.go` MUST export a function `HelloE2E202610150948432786() string` that returns exactly `"e2e-cross-repo-spawn-20261015-094843-2786"`, appended after the existing functions and following the file's existing pattern (a doc comment, then a single return).
- **FR-002**: `pkg/greeting/greeting_test.go` MUST contain a test `TestHelloE2E202610150948432786` that asserts the function returns exactly `e2e-cross-repo-spawn-20261015-094843-2786`, following the existing test pattern (a `const want` string, a `got :=` call, and `t.Errorf("HelloE2E202610150948432786() = %q; want %q", got, want)`).
- **FR-003**: The function name and string literal MUST be used verbatim as given, with no rewording or substitution of the token.
- **FR-004**: The change MUST be limited to the one function and its one test; no other files or behavior in this repo change.
- **FR-005**: The pull request for this issue MUST carry `Closes` for this issue only, and no closing keyword for any other issue.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `go test ./pkg/greeting/...` passes, including the new test.
- **SC-002**: `HelloE2E202610150948432786()` returns exactly `e2e-cross-repo-spawn-20261015-094843-2786`.
- **SC-003**: The change is merged to beta `main`, so the alpha repo can pin to the merge commit.
- **SC-004**: The diff touches only `pkg/greeting/greeting.go` and `pkg/greeting/greeting_test.go` (plus the engine-managed spec file).

## Assumptions

- The new function and test are appended at the end of their respective files, after the most recent sentinel entries, matching prior runs.
- The doc comment follows the wording style of the existing sentinels (describing the sentinel string for the cross-repo spawn regression test).
- The engine-managed `specs/539-.../spec.md` file is expected in the PR diff and is not considered a violation of the scope limit.

## Out of Scope

- Any change in the alpha repo, including pinning to or calling the new function.
- Any other change to `pkg/greeting` or elsewhere in this repo, including refactoring or removing earlier sentinel functions.
- Closing keywords for any issue other than this one.

## Source References

- Parent issue: `handarbeit/fabrik-test-alpha#8510`
- Existing pattern: `pkg/greeting/greeting.go` and `pkg/greeting/greeting_test.go`
