package tools

import "sort"

type SymbolFrequencyInfo struct {
	Frequency float64
	Symbol    byte
}

func GetFrequencyInfo(text string) []SymbolFrequencyInfo {
	m := make(map[byte]int)

	for i := 0; i < len(text); i++ {
		_, ok := m[text[i]]
		if !ok {
			m[text[i]] = 0
		}
		m[text[i]]++
	}

	result := make([]SymbolFrequencyInfo, 0)
	for k, v := range m {
		result = append(result, SymbolFrequencyInfo{
			Frequency: (float64(v) * 100.0 / float64(len(text))),
			Symbol:    k,
		})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Frequency > result[j].Frequency
	})

	return result
}
