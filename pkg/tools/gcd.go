package tools

func ValidateAphineKey(alpha, alphabetSize int) bool {
	return gcd(alpha, alphabetSize) == 1
}

func gcd(a, b int) int {
	for a != 0 && b != 0 {
		if a > b {
			a %= b
		} else {
			b %= a
		}
	}
	return max(a, b)
}
