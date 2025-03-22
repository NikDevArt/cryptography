package clipaths

import (
	"fmt"
	"hw/internal/ciphers/aphine"
)

func AphinePath() {
	fmt.Println("Введите ключи для зашифрования(alpha, betta), через пробел:")
	var alpha int
	var betta int

	fmt.Scanln(&alpha, &betta)

	cipher, err := aphine.NewAphineCipher(alpha, betta)
	if err != nil {
		fmt.Println("Ошибка:", err.Error())
		return
	}

	DoAction(cipher.Encrypt, cipher.Decrypt)
}
