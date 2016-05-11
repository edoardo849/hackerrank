package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

	// convert the first word to an integer
	n, _ := strconv.Atoi(s.Text())

	var (
		neg = float64(0)
		pos = float64(0)
		irr = float64(0)
	)

	// Loop through each word
	for i := 0; i < n; i++ {
		// Scan it
		s.Scan()
		// convert to int
		v, _ := strconv.Atoi(s.Text())
		// sum it
		if v > 0 {
			pos++
		} else if v == 0 {
			irr++
		} else if v < 0 {
			neg++
		}

	}
	count := float64(n)
	// print the sum
	fmt.Println(fmt.Sprintf("%.6f", pos/count))
	fmt.Println(fmt.Sprintf("%.6f", neg/count))
	fmt.Println(fmt.Sprintf("%.6f", irr/count))
}
