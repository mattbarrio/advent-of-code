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

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	var input []string
	for fileScanner.Scan() {
		rucksack := fileScanner.Text()
		fmt.Println(rucksack)
		firstHalf := rucksack[:len(rucksack)/2]
		secondHalf := rucksack[len(rucksack)/2:]
		for _, v := range firstHalf {
			if strings.Contains(firstHalf[v], secondHalf) {
				printf("found match %s", firstHalf[v])
			}
			break
		}
		//input = append(input, currentVal)
		fmt.Printf("1st=%s 2nd=%s %s\n", firstHalf, secondHalf, rucksack)
	}

	readFile.Close()

	fmt.Println("end")

	dict := make(map[string]int)
	for _, abc := range input {
		dict[abc] = dict[abc]
	}
	fmt.Println(dict)
}
