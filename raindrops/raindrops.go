package raindrops

import (
	"strconv"
)

func Convert(num int) string { // function converts number to string depending on factors
	str := ""
	if num%3 == 0 {
		str += "Pling"
	}
	if num%5 == 0 {
		str += "Plang"
	}
	if num%7 == 0 {
		str += "Plong"
	}
	if len(str) == 0 {
		return strconv.Itoa(num)
	}
	return str
}
