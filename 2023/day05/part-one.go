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

	groups := make(map[string][][]int)
	var setName string
	var seeds []int

	for fileScanner.Scan() {
		t := fileScanner.Text()
		if strings.Contains(t, ":") {
			f := strings.Split(t, ":")
			if f[0] == "seeds" {
				seeds = append(generateArray(strings.Fields(f[1])))
				continue
			}
			setName = strings.TrimSpace(strings.TrimSuffix(f[0], " map"))
		} else {
			n := generateArray(strings.Fields(t))
			if len(n) > 0 {
				groups[setName] = append(groups[setName], n)
			}
		}
	}

	fmt.Println(seeds, groups)

	fmt.Println("end")

}
