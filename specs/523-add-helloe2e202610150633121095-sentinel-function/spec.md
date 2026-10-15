# Feature Specification: Add HelloE2E202610150633121095 sentinel function to pkg/greeting

**Feature Branch**: `fabrik/issue-523`
**Created**: 2026-10-15
**Status**: Specified
**Input**: User description: "This is the beta-repo half of an e2e regression test for cross-repo spawn (#803, on-demand spawn-target repo init, and #1308, per-run-unique fixture token). The alpha repo (`handarbeit/fabrik-test-alpha`, parent issue #8364) needs to import a new sentinel function from this repo's `pkg/greeting` package. This change must land on beta `main` first, because alpha cannot import a function that does not yet exist. In `pkg/greeting/greeting.go`, append an exported function named exactly `HelloE2E202610150633121095` with signature `func HelloE2E202610150633121095() string` that returns exactly the literal string `\"e2e-cross-repo-spawn-20261015-063312-1095\"`, following the existing append-only pattern, including a doc comment in the same form as earlier sentinel functions. In `pkg/greeting/greeting_test.go`, append a test `TestHelloE2E202610150633121095` in the same form as the existing sentinel tests. Use the function name and the string literal exactly as written; do not invent a different name or token. The parent's final assertion depends on this exact text landing in this repo."

## Background

This issue is the beta-repo half of an end-to-end regression test for cross-repo spawn. The test covers two behaviours: on-demand initialisation of the spawn-target repo (#803) and a per-run-unique fixture token (#1308). The parent issue is `handarbeit/fabrik-test-alpha#8364`. It needs to import a new sentinel function from this repo's `pkg/greeting` package. Alpha cannot import a function that does not yet exist, so this change must be merged to beta `main` first. The parent's final assertion depends on this exact function name and string value appearing in this repo.

The `pkg/greeting` package already holds earlier sentinel functions of the same kind, each paired with a test. Those were added by earlier e2e runs and follow an append-only pattern with a doc comment on each function. This change adds one more in the same form.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Sentinel function available on beta main (Priority: P1)

The alpha repo's e2e flow imports `HelloE2E202610150633121095` from beta's `pkg/greeting` and asserts it returns the per-run token. The function and its test must exist on beta `main`, with the exact name and value.

**Why this priority**: It is the only deliverable. The parent e2e assertion fails without it.

**Independent Test**: In this repo, run `go build ./...` and `go test ./...`. The new test passes and asserts the exact token value.

**Acceptance Scenarios**:

1. **Given** beta `main` with this change merged, **When** a caller invokes `greeting.HelloE2E202610150633121095()`, **Then** it returns exactly `e2e-cross-repo-spawn-20261015-063312-1095`.
2. **Given** the change is applied, **When** `go build ./...` and `go test ./...` run, **Then** both succeed, including `TestHelloE2E202610150633121095`.
3. **Given** the change is applied, **When** the diff is reviewed, **Then** only additions appear in `pkg/greeting/greeting.go` and `pkg/greeting/greeting_test.go`, and no existing function or test is modified or reordered.

---

### Edge Cases

- The function name or token differs by even one character from the specified text, so the parent's assertion fails. Both must match exactly.
- An existing sentinel function or test is accidentally reordered or edited. The change must be append-only.
- The change sits on an unmerged branch. Alpha can only depend on it once it is on `main`.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The `pkg/greeting` package MUST export a function named exactly `HelloE2E202610150633121095` with signature `func HelloE2E202610150633121095() string`.
- **FR-002**: The function MUST return exactly the string `e2e-cross-repo-spawn-20261015-063312-1095`.
- **FR-003**: The function MUST be appended after the existing content of `pkg/greeting/greeting.go`, with a doc comment in the same form as earlier sentinel functions (for example, `// HelloE2E202610150633121095 returns the sentinel string for the cross-repo spawn regression test.`).
- **FR-004**: A test named `TestHelloE2E202610150633121095` MUST be appended to `pkg/greeting/greeting_test.go`, in the same form as the existing sentinel tests. It MUST compare the function's return value against the expected token and report a mismatch with `t.Errorf`.
- **FR-005**: Existing functions and tests MUST NOT be modified or reordered.
- **FR-006**: The change MUST NOT add dependencies or modify `go.mod`.
- **FR-007**: Only `pkg/greeting/greeting.go` and `pkg/greeting/greeting_test.go` MAY change, apart from the spec file the engine commits.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `go build ./...` and `go test ./...` pass on the branch and on `main` after merge.
- **SC-002**: Calling `HelloE2E202610150633121095()` yields a string byte-for-byte equal to `e2e-cross-repo-spawn-20261015-063312-1095`.
- **SC-003**: The change is merged to beta `main`, so the alpha repo can depend on the resulting commit.
- **SC-004**: The diff to existing source and test files is additions only.

## Assumptions

- Earlier sentinel functions and tests exist in `pkg/greeting` and follow the pattern described in the issue. The Research and Plan stages will confirm the exact form.
- The issue's own code snippets are the authoritative form for the function and test.
- This repo's existing merge process (yolo auto-merge) delivers the change to `main`.

## Out of Scope *(optional)*

- Changes to the alpha repo or its import of the new function (tracked by the parent issue `handarbeit/fabrik-test-alpha#8364`).
- Refactoring, reordering or removing existing sentinel functions and tests.
- Any new dependency, `go.mod` change or other package change.

## Source References *(optional)*

- Parent issue: `handarbeit/fabrik-test-alpha#8364`
- Related Fabrik work: #803 (on-demand spawn-target repo init), #1308 (per-run-unique fixture token)
