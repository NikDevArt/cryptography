package main

import (
	"fmt"
	"hw/internal/ciphers/simplesubstit"
	"hw/pkg/tools"
)

func main() {
	var k [2][]byte

	// init alphabet
	k[0] = []byte{
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
		'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

	// init crypto alphabet
	k[1] = []byte{
		'z', 'x', 'p', 'h', 'd', 'n', 'o', 'a', 'l', 'm', 'g', 't', 'c',
		'i', 'b', 'u', 'j', 'f', 'k', 'r', 'v', 'y', 'e', 'q', 's', 'w'}

	textPath := "../../../files/result.txt"
	text := tools.ReadFile(textPath)

	cipher := simplesubstit.NewSimpleSubstitutionCipher(k)
	encrypted := cipher.Encrypt(text)

	textFrequency := tools.GetFrequencyInfo(text)
	encryptedFrequency := tools.GetFrequencyInfo(encrypted)

	for i := 0; i < len(textFrequency); i++ {
		fmt.Println(string(encryptedFrequency[i].Symbol), "-", encryptedFrequency[i].Frequency)
	}
}
