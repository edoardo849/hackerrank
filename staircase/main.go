package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// the input is
// 6
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

	stairs := 1
	// Loop through each word
	for i := 0; i < n; i++ {

		stairsLine := ""
		whiteSpaces := n - stairs
		for j := 0; j < whiteSpaces; j++ {
			stairsLine += " "
		}
		for k := 0; k < stairs; k++ {
			stairsLine += "#"
		}

		fmt.Println(stairsLine)

		stairs++
	}

}
