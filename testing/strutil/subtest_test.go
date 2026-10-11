package strutil

import "testing"

func TestReverseSubtest(t *testing.T) {
	tests := []struct {
		name, in, want string
	}{
		{"ascii", "abc", "cba"},
		{"japanese", "あいう", "ういあ"},
		{"empty", "", ""},
		{"accent", "héllo", "olléh"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reverse(tt.in)
			if got != tt.want {
				t.Fatalf("Reverse(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
