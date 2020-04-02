package others

import (
	"fmt"
)

func timeInWords(h int32, m int32) string {
	min := "minute"
	if m > 1 {
		min += "s"
	}

	if m == 30 || m == 15 || m == 45 {
		min = ""
	}

	if m == 0 {
		return fmt.Sprintf("%s o' clock", intToWord(h))
	} else if m <= 30 {
		if min == "" {
			return fmt.Sprintf("%s past %s", intToWord(m), intToWord(h))
		}
		return fmt.Sprintf("%s %s past %s", intToWord(m), min, intToWord(h))
	}
	if min == "" {
		return fmt.Sprintf("%s to %s", intToWord(60-m), intToWord(h+1))
	}
	return fmt.Sprintf("%s %s to %s", intToWord(60-m), min, intToWord(h+1))
}

func intToWord(i int32) string {
	if i == 30 {
		return "half"
	}

	if i == 15 || i == 45 {
		return "quarter"
	}

	if i == 0 {
		return ""
	}
	to19 := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve",
		"thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}

	tens := []string{"twenty", "thirty", "forty", "fifty"}

	tenDigit := int32(i / 10)
	digit := i - tenDigit*10

	if i < 20 {
		return to19[i-1]
	}

	if digit == 0 {
		return tens[tenDigit-1]
	}
	return fmt.Sprintf("%s %s", tens[tenDigit-2], to19[digit-1])
}
