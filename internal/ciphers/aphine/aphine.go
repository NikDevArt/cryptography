package aphine

import (
	"fmt"
	"hw/pkg/alphabet"
	"hw/pkg/tools"
)

type cipherKey struct {
	alpha int
	betta int
}

func newAphineCipherKey(alpha, betta int) *cipherKey {
	k := &cipherKey{
		alpha: alpha,
		betta: betta,
	}
	return k
}

type AphineCipher struct {
	key      *cipherKey
	alphabet *alphabet.Alphabet
}

func NewAphineCipher(alpha, betta int) (*AphineCipher, error) {
	c := &AphineCipher{
		alphabet: alphabet.NewAlphabet(),
	}
	if !tools.ValidateAphineKey(alpha, c.alphabet.GetSize()) {
		return nil, fmt.Errorf("НОД %d и размера алфавита(=%d) должен быть = 1", alpha, c.alphabet.GetSize())
	}
	c.key = newAphineCipherKey(alpha, betta)
	return c, nil
}

func (c *AphineCipher) Encrypt(text string) string {
	encrypted := make([]byte, len(text))
	for i := 0; i < len(text); i++ {
		id := c.alphabet.GetIdByByte(text[i])
		encryptedId := (c.key.alpha*id + c.key.betta) % c.alphabet.GetSize()
		encrypted[i] = c.alphabet.GetSymbolById(encryptedId)
	}
	return string(encrypted)
}

func (c *AphineCipher) Decrypt(text string) string {
	decrypted := make([]byte, len(text))
	invertAlpha := tools.GetInvertNumberByMod(c.key.alpha, c.alphabet.GetSize())
	for i := 0; i < len(text); i++ {
		symbolId := c.alphabet.GetIdByByte(text[i])
		minus := (symbolId - c.key.betta + c.alphabet.GetSize()) % c.alphabet.GetSize()
		decryptedId := (invertAlpha * minus) % c.alphabet.GetSize()
		decrypted[i] = c.alphabet.GetSymbolById(decryptedId)
	}
	return string(decrypted)
}
