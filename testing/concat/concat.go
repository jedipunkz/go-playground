package concat

import "strings"

// Plus は + 演算子で文字列を連結する
func Plus(words []string) string {
	s := ""
	for _, w := range words {
		s += w
	}
	return s
}

// Builder は strings.Builder で文字列を連結する
func Builder(words []string) string {
	var sb strings.Builder
	for _, w := range words {
		sb.WriteString(w)
	}
	return sb.String()
}
