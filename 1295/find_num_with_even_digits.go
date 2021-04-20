package numeven

func NumWithEvenDigits(numbers []int) int {
	count := 0
	for _, elem := range numbers {
		if countDigits(elem) % 2 == 0 {
			count++
		}
	}
	return count
}

func countDigits(digit int) int {
	if (digit == 0) {
		return 0
	}
	// count := 0
	// for digit != 0 {
	// 	digit = digit/10
	// 	count++
	// }
	return 1 + countDigits(digit / 10)
}
