package tools

func ValidateAphineKey(alpha, alphabetSize int) bool {
	return Gcd(alpha, alphabetSize) == 1
}

func Gcd(a, b int) int {
	for a != 0 && b != 0 {
		if a > b {
			a %= b
		} else {
			b %= a
		}
	}
	return max(a, b)
}
