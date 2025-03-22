package paths

import (
	"bufio"
	"fmt"
	"hw/internal/clipaths"
	"hw/pkg/alphabet"
	"os"
	"strconv"
)

func Run() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Выберите тип шифра(число):")
		fmt.Println("1 - Шифр простой замены")
		fmt.Println("2 - Афинный шифр")
		fmt.Println("3 - Афинный рекуррентный шифр")
		fmt.Println("4 - Вывести алфавит")
		fmt.Println(">4 - Выйти из программы")

		scanner.Scan()
		choice, _ := strconv.Atoi(scanner.Text())

		switch choice {
		case 1:
			clipaths.SimpleSubstitPath()
		case 2:
			clipaths.AphinePath()
		case 3:
			clipaths.RecurrentAphinePath()
		case 4:
			a := alphabet.NewAlphabet()
			fmt.Println(a.GetFull())
		default:
			fmt.Println("Выход из программы.")
			return
		}
	}
}
