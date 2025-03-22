package alphabet

import "strings"

type Alphabet struct {
	symbolId map[byte]int
	idSymbol map[int]byte
	size     int
}

func NewAlphabet() *Alphabet {
	a := &Alphabet{}
	a.initAlphabet()
	return a
}

func (a *Alphabet) GetIdByByte(r byte) int {
	id, ok := a.symbolId[r]
	if !ok {
		return -1
	}
	return id
}

func (a *Alphabet) GetSymbolById(id int) byte {
	b, ok := a.idSymbol[id]
	if !ok {
		return byte(0)
	}
	return b
}

func (a *Alphabet) GetSize() int {
	return a.size
}

func (a *Alphabet) GetFull() string {
	var builder strings.Builder
	for i := 0; i < a.size; i++ {
		builder.WriteByte(a.idSymbol[i])
	}
	return builder.String()
}

func (a *Alphabet) initAlphabet() {
	a.symbolId = make(map[byte]int)
	a.idSymbol = make(map[int]byte)

	a.size = 0
	for i := byte('a'); i <= 'z'; i++ {
		a.symbolId[i] = a.size
		a.idSymbol[a.size] = i
		a.size++
	}

	for i := byte('A'); i <= 'Z'; i++ {
		a.symbolId[i] = a.size
		a.idSymbol[a.size] = i
		a.size++
	}

	a.symbolId['?'] = a.size
	a.idSymbol[a.size] = '?'
	a.size++

	a.symbolId['!'] = a.size
	a.idSymbol[a.size] = '!'
	a.size++
}
