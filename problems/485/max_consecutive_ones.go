package maxconsec

// MaxConsecutiveOnes returns maximum number of contagious 1s in an array of binary numbers
func MaxConsecutiveOnes(number []int) int {
	var max, now int
	for _, num := range number {
		if num == 1 {
			now++
		} else {
			now = 0
		}
		max = maxOf(now, max)
	}
	return max
}

// maxOf returns maximum value among a and b
func maxOf(a, b int) int {
	if a > b {
		return a
	}
	return b
}
