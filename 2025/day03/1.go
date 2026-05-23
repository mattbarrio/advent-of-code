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

	totalJoltage := 0
	for fileScanner.Scan() {
		bank := fileScanner.Text()
		bankSlice := strings.Split(bank, "")

		// Find the maximum two-digit number by trying all pairs
		maxJoltage := 0
		for i := 0; i < len(bankSlice); i++ {
			for j := i + 1; j < len(bankSlice); j++ {
				// Form number: first digit * 10 + second digit
				first, _ := strconv.Atoi(bankSlice[i])
				second, _ := strconv.Atoi(bankSlice[j])
				joltage := first*10 + second
				if joltage > maxJoltage {
					maxJoltage = joltage
				}
			}
		}

		fmt.Printf("Bank: %s, Max joltage: %d\n", bank, maxJoltage)
		totalJoltage += maxJoltage
	}

	fmt.Printf("Total output joltage: %d\n", totalJoltage)
}
