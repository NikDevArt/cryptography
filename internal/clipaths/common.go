package clipaths

import (
	"bufio"
	"fmt"
	"hw/pkg/tools"
	"os"
	"strconv"
)

type EncryptAction func(string) string
type DecryptAction func(string) string

func DoAction(encrypt EncryptAction, decrypt DecryptAction) {
	fmt.Println("Введите текст над которым требуется зашифрование/расшифрование:")
	text := tools.ReadString()

	fmt.Println("Введите действие(1 - зашифровать, 2 - расшифровать, остальное - сбросить данный шифр):")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	choice, _ := strconv.Atoi(scanner.Text())

	fmt.Println("############################")
	fmt.Println("Итог:")
	switch choice {
	case 1:
		encrypted := encrypt(text)
		fmt.Println("Результат зашифрования:", encrypted)
	case 2:
		decrypted := decrypt(text)
		fmt.Println("Результат расшифрования:", decrypted)
	default:
		fmt.Println("Выход")
	}
	fmt.Println("############################")
}
