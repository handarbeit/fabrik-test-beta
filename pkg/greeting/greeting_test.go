package greeting

import "testing"

func TestGreetingFor(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"", "Hello, , from fabrik-test-beta"},
		{"Alice", "Hello, Alice, from fabrik-test-beta"},
	}
	for _, tt := range tests {
		if got := GreetingFor(tt.name); got != tt.want {
			t.Errorf("GreetingFor(%q) = %q, want %q", tt.name, got, tt.want)
		}
	}
}

func TestHelloE2E(t *testing.T) {
	const want = "e2e-cross-repo-spawn"
	if got := HelloE2E(); got != want {
		t.Errorf("HelloE2E() = %q, want %q", got, want)
	}
}

func TestHelloIssue31(t *testing.T) {
	const want = "e2e-cross-repo-spawn-31"
	if got := HelloIssue31(); got != want {
		t.Errorf("HelloIssue31() = %q, want %q", got, want)
	}
}

func TestHelloIssue52(t *testing.T) {
	const want = "e2e-cross-repo-spawn-52"
	if got := HelloIssue52(); got != want {
		t.Errorf("HelloIssue52() = %q; want %q", got, want)
	}
}

func TestHelloIssue3367(t *testing.T) {
	const want = "e2e-cross-repo-spawn-3367"
	if got := HelloIssue3367(); got != want {
		t.Errorf("HelloIssue3367() = %q; want %q", got, want)
	}
}

func TestHelloIssue3472(t *testing.T) {
	const want = "e2e-cross-repo-spawn-3472"
	if got := HelloIssue3472(); got != want {
		t.Errorf("HelloIssue3472() = %q; want %q", got, want)
	}
}
