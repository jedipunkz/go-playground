package strutil_test

import (
	"fmt"

	"github.com/jedipunkz/go-playground/testing/strutil"
)

func ExampleReverse() {
	fmt.Println(strutil.Reverse("hello"))
	fmt.Println(strutil.Reverse("あいう"))
	// Output:
	// olleh
	// ういあ
}

func ExampleReverse_unordered() {
	for _, s := range map[string]string{"a": "go", "b": "テスト"} {
		fmt.Println(strutil.Reverse(s))
	}
	// Unordered output:
	// og
	// トステ
}
