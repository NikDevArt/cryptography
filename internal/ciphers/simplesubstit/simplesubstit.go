package simplesubstit

import "strings"

type key struct {
	straight map[byte]byte
	invert   map[byte]byte
}

func newSimpleSubstitutionCipherKey(keyTable [2][]byte) *key {
	k := &key{
		straight: make(map[byte]byte),
		invert:   make(map[byte]byte),
	}

	for i := 0; i < len(keyTable[0]); i++ {
		k.straight[keyTable[0][i]] = keyTable[1][i]
		k.invert[keyTable[1][i]] = keyTable[0][i]
	}

	return k
}

func (k *key) getStraight(symbol byte) byte {
	val := k.straight[symbol]
	return val
}

func (k *key) getInvert(symbol byte) byte {
	val := k.invert[symbol]
	return val
}

type SimpleSubstitutionCipher struct {
	key *key
}

func NewSimpleSubstitutionCipher(keyTable [2][]byte) *SimpleSubstitutionCipher {
	c := &SimpleSubstitutionCipher{
		key: newSimpleSubstitutionCipherKey(keyTable),
	}
	return c
}

func (c *SimpleSubstitutionCipher) Encrypt(text string) string {
	var cryptoText strings.Builder
	for _, el := range text {
		cryptoText.WriteByte(c.key.getStraight(byte(el)))
	}
	return cryptoText.String()
}

func (c *SimpleSubstitutionCipher) Decrypt(text string) string {
	var cryptoText strings.Builder
	for _, el := range text {
		cryptoText.WriteByte(c.key.getInvert(byte(el)))
	}
	return cryptoText.String()
}
