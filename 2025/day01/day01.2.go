package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// readFile, err := os.Open("test.txt")
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	// size := 100 // 0 - 99
	current := 50
	counter := 0

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())

		row := fileScanner.Text()

		direction := strings.Split(row, "")[0]
		_, t, _ := strings.Cut(row, direction)
		turn, _ := strconv.Atoi(t)
		fmt.Printf("*turn: %d\n", turn)

		switch direction {
		case "L":
			for turn > 0 {
				if current == 0 {
					current = 99
				} else {
					current--
					if current == 0 {
						counter++
					}
				}
				turn--
			}
		case "R":
			for turn > 0 {
				if current == 99 {
					current = 0
					counter++
				} else {
					current++
				}
				turn--
			}
		}

		fmt.Printf("***counter: %d, current: %d\n", counter, current)
	}
	fmt.Printf("Zeros: %d\n", counter)
	fmt.Println("end")

}
