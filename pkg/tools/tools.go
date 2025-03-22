package tools

import (
	"bufio"
	"os"
)

func GetInvertNumberByMod(number, mod int) int {
	for i := 0; i < mod; i++ {
		if (number*i)%mod == 1 {
			return i
		}
	}
	return -1
}

func ReadString() string {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	return line[:len(line)-1]
}
