package reverseint

// Reverse reverses a 32-bit signed integer.
// Returns 0 on overflow.
func Reverse(x int) (reversed int) {
	var lastDigit int
	var isNeg bool

	if x < 0 {
		isNeg = true
		x = x * -1
	}

	for x > 0 {
		lastDigit = x % 10
		reversed = (reversed * 10) + lastDigit
		x = x / 10
	}

	if isNeg {
		reversed = reversed - reversed*2
	}

	// manage the overflow
	if x > (1<<31 - 1) {
		return 0
	}

	return
}
