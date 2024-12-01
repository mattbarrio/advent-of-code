package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findStart(arr [][]string) (int, int) {
	for ri, r := range arr {
		for pi, p := range r {
			if p == "S" {
				return ri, pi
			}
		}
	}
	return 0, 0
}

func main() {

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var field [][]string

	// build 2d array
	for fileScanner.Scan() {
		t := fileScanner.Text()
		f := strings.Split(t, "")
		field = append(field, f)
	}

	fmt.Println(field)

	startRow, startColumn := findStart(field)
	fmt.Println("startRow:", startRow, "startColumn:", startColumn)

	pipe := field[startRow][startColumn+1]
	currentPosition := [2]int{startRow, startColumn + 1}
	previousPosition := [2]int{startRow, startColumn}
	counter := 0
mainLoop:
	for {
		counter++
		fmt.Println("previousPosition:", previousPosition, "currentPosition:", currentPosition, "pipe:", pipe)

		switch pipe {
		case "|":
			if previousPosition[0] < currentPosition[0] {
				previousPosition = currentPosition
				currentPosition[0]++
				pipe = field[currentPosition[0]][currentPosition[1]]
			} else {
				previousPosition = currentPosition
				currentPosition[0]--
				pipe = field[currentPosition[0]][currentPosition[1]]
			}
		case "-":
			if previousPosition[1] < currentPosition[1] {
				previousPosition = currentPosition
				currentPosition[1]++
				pipe = field[currentPosition[0]][currentPosition[1]]
			} else {
				previousPosition = currentPosition
				currentPosition[1]--
				pipe = field[currentPosition[0]][currentPosition[1]]
			}
		case "L":
			if previousPosition[0] < currentPosition[0] {
				previousPosition = currentPosition
				currentPosition[1]++
				pipe = field[currentPosition[0]][currentPosition[1]]
			} else {
				previousPosition = currentPosition
				currentPosition[0]--
				pipe = field[currentPosition[0]][currentPosition[1]]
			}
		case "J":
			if previousPosition[1] < currentPosition[1] {
				previousPosition = currentPosition
				currentPosition[0]--
				pipe = field[currentPosition[0]][currentPosition[1]]
			} else {
				previousPosition = currentPosition
				currentPosition[1]--
				pipe = field[currentPosition[0]][currentPosition[1]]
			}
		case "7":
			if previousPosition[1] < currentPosition[1] {
				previousPosition = currentPosition
				currentPosition[0]++
				pipe = field[currentPosition[0]][currentPosition[1]]
			} else {
				previousPosition = currentPosition
				currentPosition[1]--
				pipe = field[currentPosition[0]][currentPosition[1]]
			}
		case "F":
			if previousPosition[0] > currentPosition[0] {
				previousPosition = currentPosition
				currentPosition[1]++
				pipe = field[currentPosition[0]][currentPosition[1]]
			} else {
				previousPosition = currentPosition
				currentPosition[0]++
				pipe = field[currentPosition[0]][currentPosition[1]]
			}
		case "S":
			break mainLoop
		}

	}

	fmt.Println(counter / 2)

	// [[. . F 7 .]
	//  [. F J | .]
	//  [S J . L 7]
	//  [| F - - J]
	//  [L J . . .]]
	// | is a vertical pipe connecting north and south.
	// - is a horizontal pipe connecting east and west.
	// L is a 90-degree bend connecting north and east.
	// J is a 90-degree bend connecting north and west.
	// 7 is a 90-degree bend connecting south and west.
	// F is a 90-degree bend connecting south and east.
	// . is ground; there is no pipe in this tile.
	// S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.

	readFile.Close()

	fmt.Println("end")

}
