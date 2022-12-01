package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	cals := []int{0}
	previousCal := 0
	sumCal := 0
	for fileScanner.Scan() {
		currentCal, _ := strconv.Atoi(fileScanner.Text())
		if currentCal == 0 {
			// zero here means we need to save the sum into the slice
			cals = append(cals, sumCal)
		} else if previousCal == 0 {
			// zero here means we need to reset the sumCal with the current new Elf's cal
			sumCal = currentCal
		} else {
			// compound
			sumCal += currentCal
		}
		// update the previousCal
		previousCal = currentCal
		fmt.Println(currentCal, sumCal)
	}

	readFile.Close()

	fmt.Println(cals)

}
