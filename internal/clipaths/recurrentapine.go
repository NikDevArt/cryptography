package clipaths

import (
	"fmt"
	"hw/internal/ciphers/recurrentaphine"
)

func RecurrentAphinePath() {
	fmt.Println("Введите мощность алфавита(длину префикса алфавита):")
	var alphabetSize int
	fmt.Scanln(&alphabetSize)

	fmt.Println("Введите ключи для зашифрования(alpha1, betta1, alpha2, betta2)")
	var alpha1, alpha2, betta1, betta2 int

	fmt.Scanln(&alpha1, &betta1, &alpha2, &betta2)

	cipher, err := recurrentaphine.NewRecurrentAphineCipher(alpha1, betta1, alpha2, betta2, alphabetSize)
	if err != nil {
		fmt.Println("Ошибка:", err.Error())
		return
	}

	DoAction(cipher.Encrypt, cipher.Decrypt)
}
