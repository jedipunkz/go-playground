package strutil

import "unicode/utf8"

// Reverse は文字列を rune 単位で反転する
// 不正な UTF-8 を含む場合は反転せずにそのまま返す
func Reverse(s string) string {
	if !utf8.ValidString(s) {
		return s
	}
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
