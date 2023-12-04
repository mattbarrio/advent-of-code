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
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var sum int

	for fileScanner.Scan() {
		t := fileScanner.Text()
		f := strings.Split(t, ":")

		//game := f[0]
		card := f[1]
		winners := strings.Fields(strings.Split(card, "|")[0])
		current := strings.Fields(strings.Split(card, "|")[1])
		counter := 0

		fmt.Println(winners, current)
		for _, n := range current {
			for _, w := range winners {
				if n == w {
					if counter == 0 {
						counter++
					} else {
						counter = counter * 2
					}
				}
			}
		}
		sum += counter
	}

	fmt.Println(sum)

	readFile.Close()

	fmt.Println("end")

}
