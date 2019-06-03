package raindrops

import (
	"strconv"
)

func Convert(num int) string { // function converts number to string depending on factors
	str := ""
	ctr := 0
	if num%3 == 0 {
		str += "Pling"
		ctr = 1
	}
	if num%5 == 0 {
		str += "Plang"
		ctr = 1
	}
	if num%7 == 0 {
		str += "Plong"
		ctr = 1
	}
	if ctr == 1 {
		return str
	}
	return strconv.Itoa(num)
}
