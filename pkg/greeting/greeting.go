// Package greeting provides personalized greeting strings for the fabrik-test-alpha test bed.
package greeting

// GreetingFor returns a personalized greeting for the given name.
func GreetingFor(name string) string {
	return "Hello, " + name + ", from fabrik-test-beta"
}

// HelloE2E returns the sentinel string for the cross-repo spawn regression test.
func HelloE2E() string {
	return "e2e-cross-repo-spawn"
}

// HelloIssue31 returns the sentinel string for the cross-repo spawn regression test for alpha issue #31.
func HelloIssue31() string {
	return "e2e-cross-repo-spawn-31"
}

// HelloIssue52 returns the sentinel string for the cross-repo spawn regression test for alpha issue #52.
func HelloIssue52() string {
	return "e2e-cross-repo-spawn-52"
}

// HelloIssue3367 returns the sentinel string for the cross-repo spawn regression test for alpha issue #3367.
func HelloIssue3367() string {
	return "e2e-cross-repo-spawn-3367"
}

// HelloIssue3472 returns the sentinel string for the cross-repo spawn regression test for alpha issue #3472.
func HelloIssue3472() string {
	return "e2e-cross-repo-spawn-3472"
}

// HelloIssue3496 returns the sentinel string for the cross-repo spawn regression test for alpha issue #3496.
func HelloIssue3496() string {
	return "e2e-cross-repo-spawn-3496"
}
