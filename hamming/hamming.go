package hamming

import (
	"errors"
)

// Function finds the hamming distance given the sequence of two DNA strands
func Distance(a, b string) (int, error) {
	ctr := 0
	if len(a) != len(b) {
		return -1, errors.New("DNA strands of different length!")
		/* If the length of strands is not the same
		   an error is returned */
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			ctr = ctr + 1
		}
	}
	return ctr, nil
}
