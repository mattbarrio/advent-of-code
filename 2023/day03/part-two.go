package main

// This was frustrating. I wish I had the second part of the puzzle, I don't think it is
// solvable with the way I implemented the first part so this needs to be redone.
// If I come back to this puzzle, use the gear to scan around for matching numbers,
// might need to dedup if I do that though.

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

func nearbyNumber(numCords [][]int, symbolCord []int, schematic [][]string) bool {
	rowOffsets := []int{-1, 0, 1}
	colOffsets := []int{-1, 0, 1}
	for _, rowOffset := range rowOffsets {
		for _, colOffset := range colOffsets {
			row, col := symbolCord[0]+rowOffset, symbolCord[1]+colOffset

			// Check if the new position is within bounds
			if row >= 0 && row < len(schematic) && col >= 0 && col < len(schematic[0]) {

			}
		}
	}
	return false
}

func nearbySymbol(target []int, symCoordinates [][]int) (nearSymbol bool, symbolCord []int) {
	for _, pair := range symCoordinates {
		// Check if pair coordinates are within a range around the target coordinates
		if abs(pair[0]-target[0]) <= 1 && abs(pair[1]-target[1]) <= 1 {
			return true, pair
		}
	}
	return false, nil
}

func main() {

	readFile, err := os.Open("test.txt")

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
			if char == "*" {
				symCoordinates = append(symCoordinates, []int{r, c})
			}
		}
	}
	fmt.Println(symCoordinates)

	// find part numbers near gears
	for r, row := range schematic {
		var nearSymbol bool
		var symbolCord []int
		var numCords [][]int
		for c, char := range row {
			if unicode.IsDigit([]rune(char)[0]) {
				s += char
				cord := []int{r, c}
				numCords = append(numCords, cord)
				if !nearSymbol {
					nearSymbol, symbolCord = nearbySymbol(cord, symCoordinates)
				}
			} else if char == "." && len(s) == 0 {
				continue
			}

			if (!unicode.IsDigit([]rune(char)[0]) && len(s) > 0) || (c == len(row)-1 && len(s) > 0) {
				num, _ := strconv.Atoi(s)
				if nearSymbol && nearbyNumber(numCords, symbolCord, schematic) {
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
