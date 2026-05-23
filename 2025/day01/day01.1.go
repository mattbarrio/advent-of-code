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

	size := 100 // 0 - 99
	current := 50
	counter := 0

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())

		row := fileScanner.Text()

		direction := strings.Split(row, "")[0]
		_, t, _ := strings.Cut(row, direction)
		turn, _ := strconv.Atoi(t)

		switch direction {
		case "L":
			current = ((current-turn)%size + size) % size
		case "R":
			current = (current + turn) % size
		}

		if current == 0 {
			counter++
		}
	}
	fmt.Printf("Zeros: %d\n", counter)
	fmt.Println("end")

}
