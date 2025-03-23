package main

import (
	"fmt"
	"hw/internal/ciphers/recurrentaphine"
	"hw/pkg/tools"
)

func broke(encrypted string, alphabetSize int) {
	waysCount := 0
	for alpha1 := 0; alpha1 < alphabetSize; alpha1++ {
		if tools.Gcd(alpha1, alphabetSize) == 1 {
			for alpha2 := 0; alpha2 < alphabetSize; alpha2++ {
				if tools.Gcd(alpha2, alphabetSize) == 1 {
					for betta1 := 0; betta1 < alphabetSize; betta1++ {
						for betta2 := 0; betta2 < alphabetSize; betta2++ {
							cipher, _ := recurrentaphine.NewRecurrentAphineCipher(alpha1, betta1, alpha2, betta2, alphabetSize)
							decrypted := cipher.Decrypt(encrypted)
							fmt.Println(waysCount, "-", decrypted)
							waysCount++
						}
					}
				}
			}

		}
	}
	fmt.Println("Кол-во вариантов =", waysCount)
}

func main() {
	alpha1 := 3
	betta1 := 8
	alpha2 := 5
	betta2 := 6
	alphabetSize := 26

	// the text
	text := "homework"

	cipher, _ := recurrentaphine.NewRecurrentAphineCipher(alpha1, betta1, alpha2, betta2, alphabetSize)

	// encrypted text
	encrypted := cipher.Encrypt(text)

	// run iterating over values
	broke(encrypted, alphabetSize)
}
