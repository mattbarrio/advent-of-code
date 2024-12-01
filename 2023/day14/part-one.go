package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func generateArray(arr []string) []int {
	var a []int
	for _, n := range arr {
		n, _ := strconv.Atoi(n)
		a = append(a, n)
	}
	return a
}

func main() {

	readFile, err := os.Open("test.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var platform [][]string

	// build 2d array
	for fileScanner.Scan() {
		t := fileScanner.Text()
		f := strings.Split(t, "")

	}

	fmt.Println(platform)

	readFile.Close()

	fmt.Println("end")

}
