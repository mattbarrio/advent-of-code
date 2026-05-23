package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	readFile, err := os.Open("test.txt")
	//readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	counter := 0

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())

		row := fileScanner.Text()

		ranges := strings.Split(row, ",")
		fmt.Printf("ranges: %s\n", ranges)
		for _, v := range ranges {
			fmt.Printf("range: %s\n", v)
			startingId, _ := strconv.Atoi(strings.Split(v, "-")[0])
			endingId, _ := strconv.Atoi(strings.Split(v, "-")[1])
			for i := startingId; i <= endingId; i++ {
				split := len(strconv.Itoa(i)) / 2

				sSlice := strings.Split(strconv.Itoa(i), "")

				for _, s := range sSlice {
					for n := 0; n < len(sSlice); n++ {

					}
				}

				// left := strconv.Itoa(i)[:split]
				// right := strconv.Itoa(i)[split:]
				// if left == right {
				// 	fmt.Printf("count: %d\n", i)
				// 	counter += i
				// }
			}
		}

	}
	fmt.Printf("counter: %d\n", counter)
	fmt.Println("end")

}
