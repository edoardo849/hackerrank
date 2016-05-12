package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//A Discrete Mathematics professor has a class of  students. Frustrated with their lack of discipline, he decides to cancel class if fewer than  students are present when class starts.
//
//Given the arrival time of each student, determine if the class is canceled.
//
//Input Format
//
//The first line of input contains , the number of test cases.
//
//Each test case consists of two lines. The first line has two space-separated integers,  (students in the class) and (the cancelation threshold).
//The second line contains  space-separated integers () describing the arrival times for each student.
//
//Note: Non-positive arrival times () indicate the student arrived early or on time; positive arrival times () indicate the student arrived  minutes late.
//
//Output Format
//
//For each test case, print the word YES if the class is canceled or NO if it is not.
//
//Constraints
//
//Note
//If a student arrives exactly on time , the student is considered to have entered before the class started.

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))

	s.Split(bufio.ScanLines)

	// Advance to the first word
	// which is the number of test cases
	s.Scan()

	// convert the first word to an integer
	n, _ := strconv.Atoi(s.Text())

	// Loop through each word
	for i := 0; i < n; i++ {

		// Scan the requirements
		s.Scan()
		nk := convToInt(strings.Split(s.Text(), " "))
		//nStudents := nk[0]
		//k := nk[1]

		// Scan the students
		s.Scan()
		attendance := convToInt(strings.Split(s.Text(), " "))
		onTime := 0
		for _, v := range attendance {
			if v <= 0 {
				onTime++
			}
		}

		if onTime < nk[1] {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}

}

func convToInt(s []string) []int {
	i := []int{}
	for _, v := range s {
		val, _ := strconv.Atoi(v)
		i = append(i, val)
	}
	return i
}
