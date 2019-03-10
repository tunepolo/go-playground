package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var ja = "Go言語のテスト"
	fmt.Println(ja, "length:", utf8.RuneCountInString(ja))
}
