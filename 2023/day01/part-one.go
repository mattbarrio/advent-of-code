package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var numbers []int
	for fileScanner.Scan() {
		// fmt.Println(fileScanner.Text())
		row := fileScanner.Text()
		var line []int
		for _, char := range row {
			if unicode.IsDigit(char) {
				digit := int(char - '0')
				line = append(line, digit)
			}
		}
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

}
