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

		// Find the maximum 12-digit number by selecting 12 digits in order
		// Use greedy algorithm: at each position, pick the largest digit available
		// while ensuring we can still pick the remaining required digits
		selectedDigits := make([]string, 0, 12)
		lastIndex := -1
		remainingDigits := 12

		for pos := 0; pos < 12; pos++ {
			// We need to pick from [lastIndex+1, len(bank)-(remainingDigits-1)]
			// to ensure we can still pick the remaining digits
			startIdx := lastIndex + 1
			endIdx := len(bankSlice) - (remainingDigits - 1)

			// Find the maximum digit in this range
			maxDigit := -1
			maxDigitIdx := -1
			for i := startIdx; i < endIdx; i++ {
				digit, _ := strconv.Atoi(bankSlice[i])
				if digit > maxDigit {
					maxDigit = digit
					maxDigitIdx = i
				}
			}

			selectedDigits = append(selectedDigits, bankSlice[maxDigitIdx])
			lastIndex = maxDigitIdx
			remainingDigits--
		}

		// Convert selected digits to number
		maxJoltageStr := strings.Join(selectedDigits, "")
		maxJoltage, _ := strconv.Atoi(maxJoltageStr)

		fmt.Printf("Bank: %s, Max joltage: %d\n", bank, maxJoltage)
		totalJoltage += maxJoltage
	}

	fmt.Printf("Total output joltage: %d\n", totalJoltage)
}
