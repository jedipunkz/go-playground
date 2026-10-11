package concat

import (
	"fmt"
	"testing"
)

func makeWords(n int) []string {
	words := make([]string, n)
	for i := range words {
		words[i] = "go"
	}
	return words
}

func BenchmarkPlus(b *testing.B) {
	words := makeWords(100) // ループの外で準備する
	for b.Loop() {
		Plus(words)
	}
}

func BenchmarkBuilder(b *testing.B) {
	words := makeWords(100)
	b.ReportAllocs() // -benchmem なしでもメモリ割り当てを出力する
	for b.Loop() {
		Builder(words)
	}
}

func BenchmarkConcat(b *testing.B) {
	for _, n := range []int{10, 100, 1000} {
		words := makeWords(n)
		b.Run(fmt.Sprintf("impl=Plus/n=%d", n), func(b *testing.B) {
			for b.Loop() {
				Plus(words)
			}
		})
		b.Run(fmt.Sprintf("impl=Builder/n=%d", n), func(b *testing.B) {
			for b.Loop() {
				Builder(words)
			}
		})
	}
}
