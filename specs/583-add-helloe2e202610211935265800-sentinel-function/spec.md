# Feature Specification: Add HelloE2E202610211935265800 sentinel function and test to pkg/greeting

**Feature Branch**: `fabrik/issue-583`
**Created**: 2026-10-21
**Status**: Specified
**Input**: User description: "## Problem

This is the beta-repo half of a cross-repo e2e regression scenario (alpha issue `handarbeit/fabrik-test-alpha#9191`). The alpha repo's `main.go` will call a new function from this repo's `pkg/greeting` package, so the function must exist on this repo's default branch before the alpha change can land. A per-run-unique token is baked into the names below. The scenario's final assertion checks that this exact text is present in this repo, so use it verbatim.

## Requirements

1. In `pkg/greeting/greeting.go`, add an exported function with exactly this name and signature:

   ```go
   func HelloE2E202610211935265800() string
   ```

   It must return exactly this string literal:

   ```
   \"e2e-cross-repo-spawn-20261021-193526-5800\"
   ```

   Follow the existing sentinel style in the file: a doc comment of the form `// HelloE2E202610211935265800 returns the sentinel string …`, then a one-line `return` of the literal.

2. In `pkg/greeting/greeting_test.go`, add a test named `TestHelloE2E202610211935265800` that follows the existing per-sentinel pattern:

   ```go
   func TestHelloE2E202610211935265800(t *testing.T) {
       const want = \"e2e-cross-repo-spawn-20261021-193526-5800\"
       if got := HelloE2E202610211935265800(); got != want {
           t.Errorf(\"HelloE2E202610211935265800() = %q, want %q\", got, want)
       }
   }
   ```

3. Do not add any other functions, refactors, dependencies or documentation changes. `go.mod` must stay free of new dependencies.

## Token rules (do not deviate)

- The function name is `HelloE2E202610211935265800`: the digits are concatenated with **no hyphens**.
- The returned literal is `e2e-cross-repo-spawn-20261021-193526-5800`: it **has hyphens**, exactly as shown.
- Do not invent, regenerate, abbreviate or \"improve\" either string. A different name or literal fails the scenario.

## Scope notes

- Work from a fresh fetch of this repo's default branch. Any older local clone may be stale and may not show the most recent sentinel functions.
- Place the new function and test after the existing sentinels in each file (append at the end).
- If a sibling run has appended a sentinel at the same spot, keep both and rebase.

## Acceptance

- `pkg/greeting/greeting.go` on the default branch contains `HelloE2E202610211935265800() string` returning exactly `\"e2e-cross-repo-spawn-20261021-193526-5800\"`.
- `pkg/greeting/greeting_test.go` contains `TestHelloE2E202610211935265800` and `go test ./...` passes.
- `go build ./...` and `go vet ./...` pass.
- The change is merged to the default branch, so the alpha repo can pin to its merge commit."

## Background

This issue is the beta-repo half of a cross-repo e2e regression scenario, driven by alpha issue `handarbeit/fabrik-test-alpha#9191`. The alpha repo's `main.go` will call a new function from this repo's `pkg/greeting` package. That function therefore has to exist on this repo's default branch before the alpha change can land, so the alpha repo can pin to this change's merge commit.

A per-run-unique token is baked into the function name and the returned string. The scenario's final assertion checks that this exact text is present in this repo, so the names and literal must be used verbatim. The `pkg/greeting` package already holds a series of similar per-run sentinel functions and tests. This change appends one more in the same style.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Sentinel function available for the alpha repo to call (Priority: P1)

The cross-repo scenario needs a uniquely named exported function in `pkg/greeting` that returns a known string, so the alpha repo can call it and the scenario can assert the exact text exists here.

**Why this priority**: It is the whole deliverable. The alpha change cannot land until this function is on the default branch.

**Independent Test**: Build the module and call `HelloE2E202610211935265800()`. It returns `e2e-cross-repo-spawn-20261021-193526-5800`.

**Acceptance Scenarios**:

1. **Given** the default branch with this change merged, **When** the alpha repo calls `greeting.HelloE2E202610211935265800()`, **Then** it compiles and returns exactly `e2e-cross-repo-spawn-20261021-193526-5800`.
2. **Given** the merged change, **When** the scenario searches this repo for the function name and the literal, **Then** both are found verbatim.

---

### User Story 2 - Sentinel covered by a test (Priority: P1)

The sentinel has a test following the file's existing per-sentinel pattern, so regressions in the name or value are caught.

**Why this priority**: The acceptance criteria require the test to exist and `go test ./...` to pass.

**Independent Test**: Run `go test ./...`. `TestHelloE2E202610211935265800` runs and passes.

**Acceptance Scenarios**:

1. **Given** the change, **When** `go test ./...` runs, **Then** `TestHelloE2E202610211935265800` passes.
2. **Given** the function returned a different string, **When** the test runs, **Then** it fails and reports the got and want values.

---

### Edge Cases

- A sibling run may have appended another sentinel at the same spot in either file. Both are kept and the branch is rebased; no existing sentinel is altered or removed.
- A stale local clone may not show the most recent sentinels. Work starts from a fresh fetch of the default branch.
- Name and literal differ in hyphenation by design: the name has none and the literal has hyphens. Neither is to be "normalized".

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `pkg/greeting/greeting.go` MUST define the exported function `HelloE2E202610211935265800() string`.
- **FR-002**: `HelloE2E202610211935265800` MUST return exactly the literal `e2e-cross-repo-spawn-20261021-193526-5800` (with hyphens), via a one-line `return`.
- **FR-003**: The function MUST carry a doc comment of the form `// HelloE2E202610211935265800 returns the sentinel string …`, matching the existing sentinel style in the file.
- **FR-004**: `pkg/greeting/greeting_test.go` MUST contain `TestHelloE2E202610211935265800`, which asserts that the function's result equals `e2e-cross-repo-spawn-20261021-193526-5800` and reports got and want on mismatch, following the existing per-sentinel test pattern.
- **FR-005**: The new function and test MUST be appended after the existing sentinels in their respective files.
- **FR-006**: The change MUST NOT add other functions, refactors, dependencies or documentation changes, and `go.mod` MUST gain no new dependencies.
- **FR-007**: The function name and returned literal MUST be used exactly as given, with no regeneration, abbreviation or alteration.

### Key Entities

- **Sentinel function**: An exported, argument-less function in `pkg/greeting` that returns a fixed per-run-unique string, used by the e2e scenario to prove cross-repo linkage.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `pkg/greeting/greeting.go` on the default branch contains `HelloE2E202610211935265800() string` returning exactly `e2e-cross-repo-spawn-20261021-193526-5800`.
- **SC-002**: `pkg/greeting/greeting_test.go` contains `TestHelloE2E202610211935265800`, and `go test ./...` passes.
- **SC-003**: `go build ./...` and `go vet ./...` pass.
- **SC-004**: The change is merged to the default branch, so the alpha repo can pin to its merge commit.
- **SC-005**: The diff touches only `pkg/greeting/greeting.go` and `pkg/greeting/greeting_test.go` (plus the engine-generated spec file), and `go.mod` is unchanged.

## Assumptions

- The existing sentinel functions and tests in `pkg/greeting` follow a consistent style (doc comment plus one-line return; `const want` table-free test), as the issue describes.
- A sibling run appending at the same location is resolved by keeping both additions and rebasing.
- The token values in the issue are authoritative and need no validation beyond verbatim use.

## Out of Scope

- Any change to the alpha repo (`handarbeit/fabrik-test-alpha`), including the `main.go` call site.
- Refactoring or removing existing sentinels, or consolidating them.
- Documentation changes and new dependencies.
- Changes to the e2e scenario's assertions or tooling.
