// Package greeting provides the canonical greeting string used by the
// fabrik-test-alpha test bed.
package greeting

// Greeting returns the greeting string. fabrik-test-alpha consumes this.
func Greeting() string {
    return "Hello from fabrik-test-beta"
}
