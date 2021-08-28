package reverse

// ReverseString reverses an array without using any auxilary space
func ReverseString(input []string) (reversed []string) {
	if len(input) <= 1 {
		return input
	}

	inputLength := len(input) - 1
	timesRun := inputLength / 2
	for i := 0; i <= timesRun; i++ {
		input[i], input[inputLength-i] = input[inputLength-i], input[i]
	}
	return input
}
