package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func containsSubstring(s string, substrings []string) ([]string, bool) {
	var matchingKeys []string
	slen := len(s)
	for i := 0; i < slen; i++ {
		for _, sub := range substrings {
			sublen := len(sub)
			if i+sublen <= slen && s[i:i+sublen] == sub {
				matchingKeys = append(matchingKeys, sub)
			}
		}
	}
	return matchingKeys, len(matchingKeys) > 0
}

func main() {

	readFile, err := os.Open("input2.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var numbers []int
	numberMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	var keys []string
	for key := range numberMap {
		keys = append(keys, key)
	}
	fmt.Println(keys)

	for fileScanner.Scan() {
		// fmt.Println(fileScanner.Text())
		row := fileScanner.Text()
		var line []int
		var sNumFull string

		for _, char := range row {
			if unicode.IsDigit(char) {
				digit := int(char - '0')
				line = append(line, digit)
				sNumFull = ""
				continue
			}

			sNumFull += string(char)
			if matchingKeys, found := containsSubstring(sNumFull, keys); found {
				for _, matchingKey := range matchingKeys {
					digit, _ := numberMap[matchingKey]
					line = append(line, digit)
				}
			}
		}

		sNumFull = ""
		lastIndex := len(line) - 1
		number, _ := strconv.Atoi(strconv.Itoa(line[0]) + strconv.Itoa(line[lastIndex]))

		numbers = append(numbers, number)
	}
	fmt.Println(numbers)
	var answer int
	for _, num := range numbers {
		answer += num
	}
	fmt.Println(answer)

	readFile.Close()

	fmt.Println("end")

	if err := fileScanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}

}
