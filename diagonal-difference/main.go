package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Given a square matrix of size N×NN×N, calculate the absolute difference
// between the sums of its diagonals.
func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))

	// ScanWords is a split function for a Scanner that returns each
	// space-separated word of text, with surrounding spaces deleted.
	// It will never return an empty string. The definition of space is set by
	// unicode.IsSpace.
	s.Split(bufio.ScanWords)

	// Advance to the first word
	// which is the number of integers to sum
	s.Scan()

	// convert the first char to an integer
	lines, _ := strconv.Atoi(s.Text())

	lPosition := 0
	lDiagonal := 0

	rPosition := lines - 1
	rDiagonal := 0

	diff := 0

	// build the matrix
	for l := 0; l < lines; l++ {
		// Loop throug each position
		for p := 0; p < lines; p++ {

			// Scan it
			s.Scan()
			v, _ := strconv.Atoi(s.Text())

			//
			if p == rPosition {
				rDiagonal += v
			}

			if p == lPosition {
				lDiagonal += v
			}
		}

		// move the left pointer one position on the left
		lPosition++

		// move the right pointer one position on the right
		rPosition--
	}

	diff = lDiagonal - rDiagonal

	if diff < 0 {
		diff = -diff
	}

	fmt.Printf("%v", diff)

}
