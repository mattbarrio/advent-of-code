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

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var times []int
	var distances []int

	for fileScanner.Scan() {
		t := fileScanner.Text()
		f := strings.Split(t, ":")

		switch f[0] {
		case "Time":
			s, _ := strconv.Atoi(strings.Join(strings.Fields(f[1]), ""))
			times = append(times, s)
		case "Distance":
			s, _ := strconv.Atoi(strings.Join(strings.Fields(f[1]), ""))
			distances = append(distances, s)
		}
	}

	var winCounter []int
	for ti, time := range times {
		distance := distances[ti]
		var wins int
		for i := 0; i <= time; i++ {
			if (time-i)*i > distance {
				fmt.Println((time-i)*i, i, time, distance)
				wins++
			}
		}
		winCounter = append(winCounter, wins)
	}

	answer := 1
	for _, w := range winCounter {
		answer *= w
	}
	fmt.Println(answer)

	fmt.Println("end")

}
