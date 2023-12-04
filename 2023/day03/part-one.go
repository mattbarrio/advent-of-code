package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// Helper function to calculate the absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func nearbySymbol(target []int, symCoordinates [][]int) bool {

	for _, pair := range symCoordinates {
		// Check if pair coordinates are within a range around the target coordinates
		if abs(pair[0]-target[0]) <= 1 && abs(pair[1]-target[1]) <= 1 {
			return true, pair
		}
	}
	return false, nil
}

func main() {

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var schematic [][]string

	// build 2d array
	for fileScanner.Scan() {
		t := fileScanner.Text()
		f := strings.Split(t, "")
		schematic = append(schematic, f)
	}

	fmt.Println(schematic)

	n := []int{}
	symCoordinates := [][]int{}
	var s string

	// find symbols
	for r, row := range schematic {
		for c, char := range row {
			if char != "." && !unicode.IsDigit([]rune(char)[0]) {
				symCoordinates = append(symCoordinates, []int{r, c})
			}
		}
	}
	fmt.Println(symCoordinates)

	// find part numbers
	for r, row := range schematic {
		var nearSymbol bool
		for c, char := range row {
			if unicode.IsDigit([]rune(char)[0]) {
				s += char
				if !nearSymbol {
					nearSymbol = nearbySymbol([]int{r, c}, symCoordinates)
				}
			} else if char == "." && len(s) == 0 {
				continue
			}

			if (!unicode.IsDigit([]rune(char)[0]) && len(s) > 0) || (c == len(row)-1 && len(s) > 0) {
				if nearSymbol {
					num, _ := strconv.Atoi(s)
					n = append(n, num)
				}
				// we know this is the end of the number, reset for the next set
				nearSymbol = false
				s = ""
			}
		}
		fmt.Println(n)
	}

	var answer int
	for _, num := range n {
		answer += num
	}
	fmt.Println(answer)

	readFile.Close()

	fmt.Println("end")

}
