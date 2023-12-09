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

func differences(arr []int) []int {
	var diffs []int
	for ni, n := range arr {
		// don't calculate the last number in the array
		if ni != len(arr)-1 {
			diffs = append(diffs, arr[ni+1]-n)
		}
	}
	return diffs
}

func recurse(arr []int) [][]int {

	// check for non-zero
	zeros := true
	for _, num := range arr {
		if num != 0 {
			zeros = false
			break
		}
	}

	// if all zeros, stop recursion
	if zeros {
		return nil
	}
	diffs := differences(arr)

	nextDiffs := recurse(diffs)
	return append([][]int{diffs}, nextDiffs...)
}

func main() {

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var histories [][]int
	var answers []int

	for fileScanner.Scan() {
		t := fileScanner.Text()
		f := strings.Fields(t)

		histories = append(histories, generateArray(f))
	}

	for _, h := range histories {
		diffs := recurse(h)
		fullSet := append([][]int{h}, diffs...)

		// [[0 3 6 9 12 15] [3 3 3 3 3] [0 0 0 0]]
		// iterate over the arrays in reverse order
		for i := len(fullSet) - 1; i >= 0; i-- {
			// if this is the last array, append a zero
			if i == len(fullSet)-1 {
				fullSet[i] = append([]int{0}, fullSet[i]...)
				continue
			}
			// get the last item in the next array, add that to the last item in the current array
			appendNum := fullSet[i][0] - fullSet[i+1][0]
			fullSet[i] = append([]int{appendNum}, fullSet[i]...)
			if i == 0 {
				answers = append(answers, appendNum)
			}
		}

		fmt.Println("h", fullSet)
	}

	var answer int
	for _, a := range answers {
		answer += a
	}
	fmt.Println(answer)

	fmt.Println("end")

}
