package clipaths

import (
	"fmt"
	"hw/internal/ciphers/simplesubstit"
	"hw/pkg/tools"
)

func SimpleSubstitPath() {
	fmt.Println("Введите алфавит(каждый символ через пробел):")
	alphabet := tools.ReadString()

	fmt.Println("Введите инверт алфавит(каждый символ через пробел):")
	cryptoAlphabet := tools.ReadString()

	var k [2][]byte
	for id := range alphabet {
		if id%2 == 0 {
			k[0] = append(k[0], byte(alphabet[id]))
			k[1] = append(k[1], byte(cryptoAlphabet[id]))
		}
	}

	crypt := simplesubstit.NewSimpleSubstitutionCipher(len(k[0]), k)

	DoAction(crypt.Encrypt, crypt.Decrypt)
}
