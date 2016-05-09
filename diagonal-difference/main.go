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

	matrix := [][]int{}

	// build the matrix
	for l := 0; l < lines; l++ {
		row := []int{}
		// Loop throug each position
		for p := 0; p < lines; p++ {
			// Scan it
			s.Scan()
			// convert to int
			v, _ := strconv.Atoi(s.Text())
			row = append(row, v)
		}
		matrix = append(matrix, row)
	}

	lPosition := 0
	lDiagonal := 0
	rPosition := lines - 1
	rDiagonal := 0

	diff := 0

	for _, v := range matrix {
		// lines
		lDiagonal += v[lPosition]
		rDiagonal += v[rPosition]

		lPosition++
		rPosition--
	}

	diff = lDiagonal - rDiagonal

	if diff < 0 {
		diff = -diff
	}

	fmt.Println("Sum", diff)

}
