package recurrentaphine

import (
	"fmt"
	"hw/pkg/alphabet"
	"hw/pkg/tools"
)

type alphaBetta struct {
	alpha int
	betta int
}

type aphineKey struct {
	keys []*alphaBetta
}

func newAphineKey(alpha1, betta1, alpha2, betta2 int) *aphineKey {
	return &aphineKey{
		keys: []*alphaBetta{
			{
				alpha: alpha1,
				betta: betta1,
			},
			{
				alpha: alpha2,
				betta: betta2,
			},
		},
	}
}

func (k *aphineKey) addKey(alphabetSize int) {
	keys1 := k.keys[len(k.keys)-1]
	keys2 := k.keys[len(k.keys)-2]

	k.keys = append(k.keys, &alphaBetta{
		alpha: (keys1.alpha * keys2.alpha) % alphabetSize,
		betta: (keys1.betta + keys2.betta) % alphabetSize,
	})
}

func (k *aphineKey) getKeysById(id, alphabetSize int) *alphaBetta {
	for id >= len(k.keys) {
		k.addKey(alphabetSize)
	}

	return k.keys[id]
}

type RecurrentAphineCipher struct {
	key      *aphineKey
	alphabet *alphabet.Alphabet
}

func NewRecurrentAphineCipher(alpha1, betta1, alpha2, betta2 int) (*RecurrentAphineCipher, error) {
	c := &RecurrentAphineCipher{
		alphabet: alphabet.NewAlphabet(),
	}

	if !tools.ValidateAphineKey(alpha1, c.alphabet.GetSize()) {
		return nil, fmt.Errorf("НОД %d и размера алфавита(=%d) должен быть = 1", alpha1, c.alphabet.GetSize())
	}
	if !tools.ValidateAphineKey(alpha2, c.alphabet.GetSize()) {
		return nil, fmt.Errorf("НОД %d и размера алфавита(=%d) должен быть = 1", alpha2, c.alphabet.GetSize())
	}

	c.key = newAphineKey(alpha1, betta1, alpha2, betta2)
	return c, nil

}

func (c *RecurrentAphineCipher) Encrypt(text string) string {
	encrypted := make([]byte, len(text))

	for i := 0; i < len(text); i++ {
		id := c.alphabet.GetIdByByte(text[i])
		key := c.key.getKeysById(i, c.alphabet.GetSize())
		encryptedId := (key.alpha*id + key.betta) % c.alphabet.GetSize()
		encrypted[i] = c.alphabet.GetSymbolById(encryptedId)
	}

	return string(encrypted)
}

func (c *RecurrentAphineCipher) Decrypt(text string) string {
	decrypted := make([]byte, len(text))
	for i := 0; i < len(text); i++ {
		key := c.key.getKeysById(i, c.alphabet.GetSize())
		invertAlpha := tools.GetInvertNumberByMod(key.alpha, c.alphabet.GetSize())

		symbolId := c.alphabet.GetIdByByte(text[i])
		minus := (symbolId - key.betta + c.alphabet.GetSize()) % c.alphabet.GetSize()
		decryptedId := (invertAlpha * minus) % c.alphabet.GetSize()
		decrypted[i] = c.alphabet.GetSymbolById(decryptedId)
	}
	return string(decrypted)
}
