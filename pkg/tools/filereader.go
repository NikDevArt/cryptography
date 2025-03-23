package tools

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ReadFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	var result strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result.WriteString(scanner.Text())
	}

	return result.String()
}
