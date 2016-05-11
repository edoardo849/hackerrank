package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	n := s.Text()
	tArr := strings.Split(n, ":")
	hh := tArr[0]
	mm := tArr[1]
	ss := tArr[2]

	vInt, _ := strconv.Atoi(hh)
	if strings.Contains(ss, "AM") {
		ss = strings.TrimSuffix(ss, "AM")
		if vInt == 12 {
			hh = "00"
		} else {

			hh = strconv.FormatInt(int64(vInt), 10)
			if (vInt) < 10 {
				hh = "0" + hh
			}

		}
	} else if strings.Contains(ss, "PM") {

		ss = strings.TrimSuffix(ss, "PM")
		if vInt != 12 {
			hh = strconv.FormatInt(int64(vInt)+12, 10)
		}
	}
	fmt.Println(fmt.Sprintf("%s:%s:%s", hh, mm, ss))
}
