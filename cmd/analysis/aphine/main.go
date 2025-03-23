package main

import (
	"fmt"
	"hw/internal/ciphers/aphine"
	"hw/pkg/tools"
)

func broke(encrypted string, alphabetSize int) {
	waysCount := 0
	for alpha := 0; alpha < alphabetSize; alpha++ {
		if tools.Gcd(alpha, alphabetSize) == 1 {
			for betta := 0; betta < alphabetSize; betta++ {
				cipher, _ := aphine.NewAphineCipher(alpha, betta, alphabetSize)
				decrypted := cipher.Decrypt(encrypted)
				fmt.Println(decrypted)
				waysCount++
			}
		}
	}
	fmt.Println("Кол-во вариантов =", waysCount)
}

func main() {
	alpha := 3
	betta := 8
	alphabetSize := 26

	// the text
	text := "homework"

	cipher, _ := aphine.NewAphineCipher(alpha, betta, alphabetSize)

	// encrypted text
	encrypted := cipher.Encrypt(text)

	// run iterating over values
	broke(encrypted, alphabetSize)
}
