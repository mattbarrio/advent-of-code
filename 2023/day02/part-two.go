package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func handValues(h []string) (red int, green int, blue int) {
	var r int
	var g int
	var b int

	for _, hand := range h {
		color := strings.Fields(strings.TrimSpace(hand))[1]
		count, _ := strconv.Atoi(strings.Fields(strings.TrimSpace(hand))[0])
		switch color {
		case "red":
			if r < count {
				r = count
			}
		case "green":
			if g < count {
				g = count
			}
		case "blue":
			if b < count {
				b = count
			}
		}
	}
	return r, g, b
}

func largest(arr []int) (smallest int) {
	// avoid O(n log n) with sort.Ints(arr)
	for _, num := range arr {
		if num > smallest {
			smallest = num
		}
	}

	return smallest
}

func main() {

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var red, green, blue []int
	var largestRed, largestGreen, largestBlue, sum int
	for fileScanner.Scan() {
		// fmt.Println(fileScanner.Text())
		t := fileScanner.Text()
		game, _ := strconv.Atoi(strings.Fields(strings.Split(t, ":")[0])[1])
		fields := strings.Split(t, ":")[1]
		hands := strings.Split(fields, ";")

		for _, h := range hands {
			hand := strings.Split(h, ",")

			r, g, b := handValues(hand)

			red = append(red, r)
			green = append(green, g)
			blue = append(blue, b)
			fmt.Println(red, green, blue)
		}

		largestRed = largest(red)
		largestGreen = largest(green)
		largestBlue = largest(blue)
		fmt.Println(largestRed, largestGreen, largestBlue)
		sum += largestRed * largestGreen * largestBlue
		red = []int{}
		green = []int{}
		blue = []int{}
		fmt.Printf("Game: %d, Fields: %q, Sum: %d\n", game, fields, sum)
	}
	fmt.Println("Sum:", sum)
	readFile.Close()

	fmt.Println("end")

}
