package main

import (
	"fmt"
	"hw/pkg/tools"
	"os"
	"strings"
)

func main() {
	// Open the file
	text := tools.ReadFile("../../files/article.txt")
	fmt.Println(text)

	var textResult strings.Builder
	for i := 0; i < len(text); i++ {
		symbol := text[i]
		if symbol >= 'A' && symbol <= 'Z' {
			symbol = byte('a' + (symbol - 'A'))
		}
		if 'a' <= symbol && symbol <= 'z' {
			textResult.WriteByte(symbol)
		}
	}

	os.WriteFile("../../files/result.txt", []byte(textResult.String()), 0644)
}
