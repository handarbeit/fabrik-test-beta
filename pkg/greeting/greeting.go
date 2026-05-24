// Package greeting provides personalized greeting strings for the fabrik-test-alpha test bed.
package greeting

// GreetingFor returns a personalized greeting for the given name.
func GreetingFor(name string) string {
	return "Hello, " + name + ", from fabrik-test-beta"
}
