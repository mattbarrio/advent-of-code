package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var cards [][][]string
	var winners []string
	var current []string

	// build cards
	for fileScanner.Scan() {
		t := fileScanner.Text()
		f := strings.Split(t, ":")

		card := f[1]
		winners = strings.Fields(strings.Split(card, "|")[0])
		current = strings.Fields(strings.Split(card, "|")[1])

		cards = append(cards, [][]string{winners, current})
		fmt.Println(winners, current)
	}

	var counts []int

	var wins int
	for i, card := range cards {
		if i >= len(counts) {
			counts = append(counts, make([]int, i+1-len(counts))...)
		}
		counts[i] += 1
		fmt.Println("starting:", card, counts)
		for _, n := range card[1] {
			for _, w := range card[0] {
				if n == w {
					// fmt.Println("win:", n, w)
					wins++
				}
			}
		}

		for win := 0; win < wins; win++ {
			if i+win+1 >= len(counts) {
				counts = append(counts, 0)
			}
			counts[i+win+1] += counts[i]
			// fmt.Println("counts:", i, i+win, counts)
		}
		wins = 0
	}

	var answer int
	for _, count := range counts {
		answer += count
	}

	fmt.Println("answer:", answer)

	fmt.Println("end")

}
