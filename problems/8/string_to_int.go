package stringtoint

// Atoi converts a string to an integer.
// Edge cases are available at:
// https://leetcode.com/problems/string-to-integer-atoi/
func Atoi(str string) int {
	return myAtoi(str)
}

func myAtoi(str string) int {
	var isNeg bool

	if str == "" {
		return 0
	}

	trimmedStr := trimLeft(str)

	if trimmedStr == "" {
		return 0
	}

	trimmedStr = stripRight(trimmedStr)
	if trimmedStr == "" {
		return 0
	}

	var final int

	// deal with optional '+' or '-' sign
	if trimmedStr[0] == '-' {
		isNeg = true
		trimmedStr = trimmedStr[1:]
	} else if trimmedStr[0] == '+' {
		trimmedStr = trimmedStr[1:]
	}

	// do the conversion
	for i := 0; i < len(trimmedStr); i++ {
		// subtract the value of integer at '0' from the
		// value of integer at the char's
		final = final*10 + int(trimmedStr[i]-'0')
	}

	if isNeg {
		final = final * -1
	}

	// if exceeds 32bit signed, return the last valid bit
	if final < (-1 << 31) {
		return (-1 << 31)
	} else if final > (1<<31)-1 {
		return (1 << 31) - 1
	}

	return final
}

// trimLeft discards all the 32 (whitespace) at the left of
// the first encountered non-space character
func trimLeft(str string) string {
	var seenspace bool
	var ndigit, start int
	for i := 0; i < len(str); i++ {

	}

	// trimmed := ""

	// for _, c := range str {
	// 	if c != 32 {
	// 		trimmed = trimmed + string(c)
	// 	}
	// }

	// return trimmed
}

// stripRight discards all the character at right of
// first non digit encountered
// Expects input to be trimmed from left already
func stripRight(str string) string {
	final := ""
	var plusorminus byte
	if str[0] == '-' || str[0] == '+' {
		plusorminus = str[0]
		str = str[1:]
	}

	for _, c := range str {
		if c >= '0' && c <= '9' {
			final = final + string(c)
		} else {
			break
		}
	}

	if plusorminus != 0 {
		final = string(plusorminus) + final
	}

	return final
}
