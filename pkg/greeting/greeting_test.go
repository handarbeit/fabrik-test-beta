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
