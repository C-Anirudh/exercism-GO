package scrabble

import (
	"strings"
)

func Score(txt string) int { // function returns the scrabble score of input string
	score := 0
	s := strings.ToLower(txt)
	str := []rune(s)
	for i := 0 ; i < len(str) ; i++ {
		if strings.Contains("aeioulnrst", string(str[i])) {
			score += 1
		} else if strings.Contains("dg", string(str[i])) {
			score += 2
		} else if strings.Contains("bcmp", string(str[i])) {
			score += 3
		} else if strings.Contains("fhvwy", string(str[i])) {
			score += 4
		} else if strings.Contains("k", string(str[i])) {
			score += 5
		} else if strings.Contains("jx", string(str[i])) {
			score += 8
		} else {
			score += 10
		}
	}
	return score
}