package calc

import (
	"errors"
	"testing"
)

func TestDivParallel(t *testing.T) {
	tests := []struct {
		name    string
		a, b    int
		want    int
		wantErr error
	}{
		{name: "割り切れる", a: 10, b: 2, want: 5},
		{name: "切り捨て", a: 7, b: 2, want: 3},
		{name: "負の数", a: -9, b: 3, want: -3},
		{name: "ゼロ除算", a: 1, b: 0, wantErr: ErrDivideByZero},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := Div(tt.a, tt.b)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("Div(%d, %d) error = %v, want %v", tt.a, tt.b, err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("Div(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
