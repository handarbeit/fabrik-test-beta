package greeting

import "testing"

func TestGreeting(t *testing.T) {
    if got := Greeting(); got == "" {
        t.Fatal("Greeting() returned empty")
    }
}
