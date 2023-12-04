package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func handViable(h []string) bool {
	var red int
	var green int
	var blue int

	for _, r := range h {
		color := strings.Fields(strings.TrimSpace(r))[1]
		count, _ := strconv.Atoi(strings.Fields(strings.TrimSpace(r))[0])
		switch color {
		case "red":
			red += count
			if red > 12 {
				return false
			}
		case "green":
			green += count
			if green > 13 {
				return false
			}
		case "blue":
			blue += count
			if blue > 14 {
				return false
			}
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

	sum := 0
	for fileScanner.Scan() {
		// fmt.Println(fileScanner.Text())
		t := fileScanner.Text()
		handBroken := false
		game, _ := strconv.Atoi(strings.Fields(strings.Split(t, ":")[0])[1])
		fields := strings.Split(t, ":")[1]
		hands := strings.Split(fields, ";")
		for _, h := range hands {
			hand := strings.Split(h, ",")
			if !handViable(hand) {
				handBroken = true
				break
			}
		}
		if !handBroken {
			sum += game
		}

		fmt.Printf("Game: %d, Fields: %q, Sum: %d, Broken?: %t\n", game, fields, sum, handBroken)
	}
	fmt.Println("Sum:", sum)
	readFile.Close()

	fmt.Println("end")

}
