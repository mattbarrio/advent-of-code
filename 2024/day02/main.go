package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// The levels are either all increasing or all decreasing.
// AND
// Any two adjacent levels differ by at least one and at most three.

func safe(report []int) bool {
	var increasing bool = report[0] < report[1]

	for i, v := range report {
		// i'm sure there is a more efficient way to combine all these...
		if i == len(report)-1 {
			if v > report[i-1] && v-report[i-1] < 4 && v-report[i-1] > 0 && increasing {
				return true
			} else if v < report[i-1] && report[i-1]-v < 4 && report[i-1]-v > 0 && !increasing {
				return true
			} else {
				return false
			}
			return true
		}

		if v > report[i+1] && v-report[i+1] < 4 && v-report[i+1] > 0 && (i == 0 || v < report[i-1]) {
			continue
		} else if v < report[i+1] && report[i+1]-v < 4 && report[i+1]-v > 0 && (i == 0 || v > report[i-1]) {
			continue
		} else {
			return false
		}
	}
	return true
}

func main() {

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var valid int
	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
		line := fileScanner.Text()

		var report []int
		r := strings.Split(line, " ")
		for _, v := range r {
			val, _ := strconv.Atoi(v)
			report = append(report, val)
		}

		if safe(report) {
			fmt.Println("adding")
			valid++
		} else {
			// part 2 shit
		retry:
			for i, _ := range report {
				// wtf Matt... is there no better way to copy slices to one another?
				// also, why does using slices.Delete() fuck the original slice?
				retryReport := append(append([]int{}, report[:i]...), report[i+1:]...)
				fmt.Println(retryReport)
				if safe(retryReport) {
					valid++
					fmt.Println("retry succeeded")
					break retry
				}
			}
		}
	}
	fmt.Println(valid)
	fmt.Println("end")

}
