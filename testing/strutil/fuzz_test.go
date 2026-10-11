package strutil

import (
	"testing"
	"unicode/utf8"
)

func FuzzReverse(f *testing.F) {
	// seed corpus: 入力を生成するときの起点になる値
	for _, s := range []string{"abc", "あいう", "", "héllo"} {
		f.Add(s)
	}

	f.Fuzz(func(t *testing.T, s string) {
		rev := Reverse(s)
		// 性質 1: 2 回反転すると元に戻る
		if got := Reverse(rev); got != s {
			t.Errorf("Reverse(Reverse(%q)) = %q", s, got)
		}
		// 性質 2: 正しい UTF-8 を反転しても正しい UTF-8 のまま
		if utf8.ValidString(s) && !utf8.ValidString(rev) {
			t.Errorf("Reverse(%q) = %q is not valid UTF-8", s, rev)
		}
	})
}
