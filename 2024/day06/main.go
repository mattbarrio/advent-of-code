package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

// 	   0 1 2 3 4 5 6 7 8 9
// [0 [. . . . # . . . . .]
//	1 [. . . . . . . . . #]
//	2 [. . . . . . . . . .]
//	3 [. . # . . . . . . .]
//	4 [. . . . . . . # . .]
//	5 [. . . . . . . . . .]
//	6 [. # . . ^ . . . . .]
//	7 [. . . . . . . . # .]
//	8 [# . . . . . . . . .]
//	9 [. . . . . . # . . .]]

var obstructions = "#"
var guard = "^"
var start = []int{}

func checkBounds(puzzle [][]string, y int, x int) bool {
	fmt.Printf("Checking bound [%d][%d]\n", y, x)
	if x < 0 || x >= len(puzzle) {
		return false
	}
	if y < 0 || y >= len(puzzle[x]) {
		return false
	}
	return true
}

func safeAhead(puzzle [][]string, y int, x int) bool {
	fmt.Printf("Checking safe  [%d][%d]\n", y, x)
	if puzzle[y][x] == obstructions {
		return false
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

	var puzzle [][]string

	// build 2d array
	for fileScanner.Scan() {
		t := fileScanner.Text()
		f := strings.Split(t, "")
		puzzle = append(puzzle, f)
	}

	for r, row := range puzzle {
		for c, col := range row {
			if col == guard {
				start = []int{r, c}
				fmt.Printf("Guard at [%d][%d]: %s\n", r, c, puzzle[r][c])
				fmt.Println(start)
			}
		}
	}

	var direction string = "up"
	var distinct []string
	distinct = append(distinct, strconv.Itoa(start[0])+","+strconv.Itoa(start[1]))
	next := []int{start[0] - 1, start[1]}
	for true {
		fmt.Println(distinct)
		// if out of bounds, end
		if !checkBounds(puzzle, next[0], next[1]) {
			fmt.Println("out of bounds")
			break
		}

		// if false turn right
		if safeAhead(puzzle, next[0], next[1]) {
			if !slices.Contains(distinct, strconv.Itoa(next[0])+","+strconv.Itoa(next[1])) {
				distinct = append(distinct, strconv.Itoa(next[0])+","+strconv.Itoa(next[1]))
			}
			switch direction {
			case "up":
				next = []int{next[0] - 1, next[1]}
			case "right":
				next = []int{next[0], next[1] + 1}
			case "down":
				next = []int{next[0] + 1, next[1]}
			case "left":
				next = []int{next[0], next[1] - 1}
			}
		} else {
			fmt.Println("not safe")
			switch direction {
			case "up":
				// turn right
				direction = "right"
				next = []int{next[0] + 1, next[1] + 1}
			case "right":
				// turn down
				direction = "down"
				next = []int{next[0] + 1, next[1] - 1}
			case "down":
				// turn left
				direction = "left"
				next = []int{next[0] - 1, next[1] - 1}
			case "left":
				// turn up
				direction = "up"
				next = []int{next[0] - 1, next[1] + 1}
			}
		}
		fmt.Println(direction)
	}
	fmt.Println(len(distinct))
	fmt.Println("end")

}
