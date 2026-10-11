package strutil

import "testing"

func TestReverseLoop(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"abc", "cba"},
		{"あいう", "ういあ"},
		{"", ""},
		{"héllo", "olléh"},
	}

	for _, tt := range tests {
		got := Reverse(tt.in)
		if got != tt.want {
			t.Fatalf("Reverse(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
