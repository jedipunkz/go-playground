package calc

import (
	"errors"
	"testing"
)

func TestDiv(t *testing.T) {
	// テストケースをテーブル (構造体のスライス) として定義する
	tests := []struct {
		a, b    int
		want    int
		wantErr error
	}{
		{a: 10, b: 2, want: 5},
		{a: 7, b: 2, want: 3},
		{a: -9, b: 3, want: -3},
		{a: 1, b: 0, wantErr: ErrDivideByZero},
	}

	// 同じ検証ロジックを全ケースに適用する
	for _, tt := range tests {
		got, err := Div(tt.a, tt.b)
		if !errors.Is(err, tt.wantErr) {
			t.Errorf("Div(%d, %d) error = %v, want %v", tt.a, tt.b, err, tt.wantErr)
		}
		if got != tt.want {
			t.Errorf("Div(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}
