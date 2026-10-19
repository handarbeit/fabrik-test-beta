# Feature Specification: add HelloE2E202610192116141137 to the greeting package

**Feature Branch**: `fabrik/issue-581`
**Created**: 2026-10-19
**Status**: Specified
**Input**: User description: "This issue was spawned from handarbeit/fabrik-test-alpha#9160 (e2e cross-repo spawn). It is the prerequisite half of a two-repo change. The alpha repo will call this function after it lands, so this must land first."

## Background

This issue was spawned from handarbeit/fabrik-test-alpha#9160 (e2e cross-repo spawn). It is the prerequisite half of a two-repo change. The alpha repo will call this function after it lands, so this must land first.

The greeting package follows a one-function-per-run convention: each e2e run adds one exported function that returns a per-run-unique string, with a matching test. A per-run-unique token is baked into both the function name and the returned string. The scenario's final assertion checks that exactly this text lands on this repo's default branch, so the name and literal must be used exactly as specified.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Add the per-run greeting function (Priority: P1)

A caller in the alpha repo needs an exported function in this repo's greeting package that returns a known per-run-unique string, so the cross-repo scenario can verify the change landed.

**Why this priority**: It is the only deliverable. The alpha repo's half of the change depends on it landing first.

**Independent Test**: Run the new unit test, then `go build ./...` and `go test ./...`. All pass, and the function returns the expected string.

**Acceptance Scenarios**:

1. **Given** the greeting package, **When** `HelloE2E202610192116141137()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261019-211614-1137`.
2. **Given** the change is merged, **When** the default branch is inspected, **Then** it contains both `HelloE2E202610192116141137` and the literal `e2e-cross-repo-spawn-20261019-211614-1137`.

---

### Edge Cases

- The function name has no separators between the date, time and counter digits, while the literal separates them with hyphens. Neither may be normalised or "corrected" to match the other.
- Existing greeting functions and tests must be unchanged, so there is no collision with or alteration of earlier runs' sentinels.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `pkg/greeting/greeting.go` MUST define an exported function `HelloE2E202610192116141137() string` that returns exactly `"e2e-cross-repo-spawn-20261019-211614-1137"`. Give it a doc comment and a one-line body, following the file's existing one-function-per-run convention (for example `func HelloIssue3737() string { return "e2e-cross-repo-spawn-3737" }`).
- **FR-002**: `pkg/greeting/greeting_test.go` MUST contain a test, `TestHelloE2E202610192116141137`, that asserts the return value. Follow the existing style: a `const want = "e2e-cross-repo-spawn-20261019-211614-1137"` and `if got := HelloE2E202610192116141137(); got != want { t.Errorf("HelloE2E202610192116141137() = %q; want %q", got, want) }`.
- **FR-003**: The function name and string literal MUST be used verbatim as written above. Do not invent a substitute name or token. A per-run-unique token is baked into both, and the scenario's final assertion checks that exactly this text lands on this repo's default branch. The function name has no separators between the date, time and counter digits, while the literal separates them with hyphens.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: Calling `HelloE2E202610192116141137()` returns exactly `e2e-cross-repo-spawn-20261019-211614-1137`.
- **SC-002**: The project builds and its full test suite passes with the change in place.
- **SC-003**: After merge, the default branch contains both the function name `HelloE2E202610192116141137` and the literal `e2e-cross-repo-spawn-20261019-211614-1137`.

## Assumptions

- The existing one-function-per-run style in the greeting package (doc comment, one-line body, `const want` test) is the convention to follow.
- No documentation changes are needed, as stated in the original request.

## Out of Scope

- Any change to existing functions or tests in the greeting package.
- Changes to `go.mod` or new dependencies.
- Documentation changes.
- The alpha repo's call to this function, which is handled in handarbeit/fabrik-test-alpha#9160.

## Scope

- Change only `pkg/greeting/greeting.go` and `pkg/greeting/greeting_test.go`.
- Do not change any existing function or test in the greeting package.
- Do not change `go.mod` and do not add dependencies. No documentation changes are needed.

## Acceptance

- Calling `HelloE2E202610192116141137()` returns exactly `e2e-cross-repo-spawn-20261019-211614-1137`.
- `go build ./...` and `go test ./...` pass.
- After merge, the default branch of this repo contains both `HelloE2E202610192116141137` and the literal `e2e-cross-repo-spawn-20261019-211614-1137`.

## Notes

- Mention handarbeit/fabrik-test-alpha#9160 in the PR body for context only.
- The PR body needs `Closes #<this issue's number>` so Fabrik can discover it.
- Do not use a closing keyword (`Closes`/`Fixes`/`Resolves`) for the alpha issue.
---

*Spawned by Fabrik from parent issue handarbeit/fabrik-test-alpha#9160 as a multi-issue decomposition. The parent's plan is at the link above.*
