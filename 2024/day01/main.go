package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func findNums(arr []int, num int) int {
	var foundNums int
	for _, arrNum := range arr {
		if arrNum == num {
			foundNums++
		}
	}
	return foundNums
}

func main() {

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var left []int
	var right []int

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
		line := fileScanner.Text()

		before, after, _ := strings.Cut(line, "   ")

		leftNum, _ := strconv.Atoi(before)
		rightNum, _ := strconv.Atoi(after)

		left = append(left, leftNum)
		right = append(right, rightNum)
	}

	slices.Sort(left)
	slices.Sort(right)

	var similarity int
	for _, leftNum := range left {
		similarity += findNums(right, leftNum) * leftNum
	}

	fmt.Println(left)
	fmt.Println(right)
	fmt.Println(similarity)

	fmt.Println("end")

}
