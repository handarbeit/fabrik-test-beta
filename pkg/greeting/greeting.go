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
