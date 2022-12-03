package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	totalScorePt1 := 0
	totalScorePt2 := 0
	// shape you selected (1 for Rock, 2 for Paper, and 3 for Scissors) plus the score for the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won

	handScorePt1 := map[string]map[string]int{
		"X": { // rock
			"A": 4, // v rock (d)
			"B": 1, // v paper (l)
			"C": 7, // v scissors (w)
		},
		"Y": { // paper
			"A": 8, // v rock (w)
			"B": 5, // v paper (d)
			"C": 2, // v scissors (l)
		},
		"Z": { // scissors
			"A": 3, // v rock (l)
			"B": 9, // v paper (w)
			"C": 6, // v scissors (d)
		},
	}

	// day2
	// X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win. Good luck!"
	handScorePt2 := map[string]map[string]int{
		"X": { // loss
			"A": 3, // scissors v rock
			"B": 1, // rock v paper
			"C": 2, // paper v scissors
		},
		"Y": { // draw
			"A": 4, // rock v rock
			"B": 5, // paper v paper
			"C": 6, // scissors v scissors
		},
		"Z": { // win
			"A": 8, // paper v rock
			"B": 9, // scissors v paper
			"C": 7, // rock v scissors
		},
	}

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		hand := strings.Fields(fileScanner.Text())
		totalScorePt1 += handScorePt1[hand[1]][hand[0]]
		totalScorePt2 += handScorePt2[hand[1]][hand[0]]
	}

	fmt.Println(totalScorePt1, totalScorePt2)

	fmt.Println("end")

}
