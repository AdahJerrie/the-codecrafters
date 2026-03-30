package calculatorfunctions

func Pow(a, b int) int {
	if b == 0 {
		return 1
	} else if b < 0 {
		result := 1
		for i := 0; i < b; i++ {
			result *= a
		}
		final := 1 / result
		return final
	}

	result := 1
	for i := 0; i < b; i++ {
		result *= a
	}
	return result
}
