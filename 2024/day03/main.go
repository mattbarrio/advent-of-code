package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))

func main() {

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var instruction int = 0
	var enabled bool = true
	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
		line := fileScanner.Text()

		// r := strings.Split(line, "")
		r := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|(don't)\(\)|(do)\(\)`)
		matches := r.FindAllStringSubmatch(line, -1)
		fmt.Println(matches)
		for _, v := range matches {
			fmt.Println(v)
			if v[4] == "do" {
				enabled = true
			} else if v[3] == "don't" {
				enabled = false
			} else {
				if enabled {
					fmt.Println(v[1], v[2])
					left, _ := strconv.Atoi(v[1])
					right, _ := strconv.Atoi(v[2])
					instruction += left * right
				}
			}
		}
	}
	fmt.Println(instruction)
	fmt.Println("end")

}
