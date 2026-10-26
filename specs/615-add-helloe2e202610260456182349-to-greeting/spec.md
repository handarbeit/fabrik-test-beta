# Feature Specification: beta greeting function for the cross-repo spawn e2e

**Feature Branch**: `fabrik/issue-615`
**Created**: 2026-10-26
**Status**: Specified
**Input**: User description: "This is the beta-side part of a cross-repo e2e regression scenario spawned from `handarbeit/fabrik-test-alpha#9664`. The alpha repo imports this package as a Go module and cannot build until the function below exists on beta `main`. Land this change first."

## Background

The alpha repository (`handarbeit/fabrik-test-alpha`) is running a cross-repo spawn regression scenario. Its parent issue (#9664) depends on the beta repository's greeting package exposing a per-run-unique sentinel function. Alpha imports the package as a Go module, so alpha cannot build until the function exists on beta `main`. This issue supplies that function and its test, and must merge first so alpha can bump its dependency pin to the merged commit. The scenario's final assertion checks the exact function name and string in this repo, so the text must match exactly.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Sentinel function available on beta main (Priority: P1)

The e2e scenario needs the greeting package to expose a sentinel function returning a known, per-run-unique string, so that the alpha repo can build against it and the scenario can assert on its presence.

**Why this priority**: It is the whole deliverable. Alpha's parent issue is blocked until it lands.

**Independent Test**: Call the new function and compare its result with the required literal. Run the package tests and confirm the new test passes.

**Acceptance Scenarios**:

1. **Given** the greeting package on beta `main` after this change, **When** `HelloE2E202610260456182349()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261026-045618-2349`.
2. **Given** the change is applied, **When** the full test suite is run, **Then** it passes, including `TestHelloE2E202610260456182349`.
3. **Given** the change is merged to beta `main`, **When** the alpha repo bumps its beta dependency pin to the merged commit, **Then** alpha can build against the new function.

---

### Edge Cases

- The function name and string contain a per-run-unique token (`20261026-045618-2349`). Any alteration, such as shortening, re-casing or substitution, breaks the scenario's final assertion.
- Existing sentinel functions and their tests must remain unchanged, so the diff is purely additive.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The greeting package MUST export a function named exactly `HelloE2E202610260456182349` that takes no arguments and returns a `string`.
- **FR-002**: `HelloE2E202610260456182349()` MUST return exactly the literal `e2e-cross-repo-spawn-20261026-045618-2349`.
- **FR-003**: The function MUST follow the style of the existing `HelloE2E<token>()` functions, placed after the most recent one and carrying a doc comment, for example `// HelloE2E202610260456182349 returns the sentinel string for the cross-repo spawn regression test.`
- **FR-004**: The package's tests MUST include `TestHelloE2E202610260456182349`, following the existing per-function tests. It declares `const want = "e2e-cross-repo-spawn-20261026-045618-2349"` and reports an error via `t.Errorf` if the returned value differs.
- **FR-005**: No other function or test in the greeting package may be changed.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `HelloE2E202610260456182349()` returns a string byte-for-byte equal to `e2e-cross-repo-spawn-20261026-045618-2349`.
- **SC-002**: `go test ./...` passes on the change, including the new test.
- **SC-003**: The change is merged to beta `main`, and the parent issue in alpha can pin to the merged commit.
- **SC-004**: The diff contains only the one added function and the one added test, with no modifications to existing code.

## Assumptions

- The most recent existing `HelloE2E<token>()` function in the greeting package is the correct insertion point, and its style is the template to follow.
- No documentation changes are needed, since this is a fixture change only.
- The new name and string do not yet exist in the repository.

## Out of Scope *(optional)*

- Any change to the alpha repository, including bumping its dependency pin (handled by the parent issue).
- Changes to any other greeting function or test.
- Documentation updates.

## What needs to change

In this repo (`handarbeit/fabrik-test-beta`):

- In `pkg/greeting/greeting.go`, add an exported function `HelloE2E202610260456182349() string` that returns exactly the literal string `e2e-cross-repo-spawn-20261026-045618-2349`. Follow the style of the existing `HelloE2E<token>()` functions: append it after the most recent one, with a doc comment such as `// HelloE2E202610260456182349 returns the sentinel string for the cross-repo spawn regression test.`
- In `pkg/greeting/greeting_test.go`, add `TestHelloE2E202610260456182349`. Follow the existing per-function tests: declare `const want = "e2e-cross-repo-spawn-20261026-045618-2349"` and call `t.Errorf` if `got != want`.

## Constraints

- Use the function name `HelloE2E202610260456182349` and the string literal `e2e-cross-repo-spawn-20261026-045618-2349` exactly as written. Do not rename, alter, shorten or substitute them. A per-run-unique token (`20261026-045618-2349`) is already baked into both, and the scenario's final assertion checks this exact text in this repo.
- Do not change any other function or test in the greeting package.
- This is a fixture change only. No documentation changes are needed.

## Acceptance

1. `HelloE2E202610260456182349()` returns exactly `e2e-cross-repo-spawn-20261026-045618-2349`.
2. `go test ./...` passes, including the new test.
3. The change is merged to beta `main`. The parent issue (`handarbeit/fabrik-test-alpha#9664`) is blocked on this issue and will bump its beta dependency pin to the merged commit.
---

*Spawned by Fabrik from parent issue handarbeit/fabrik-test-alpha#9664 as a multi-issue decomposition. The parent's plan is at the link above.*
