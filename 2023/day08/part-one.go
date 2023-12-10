package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findNext(arr [][][]string, current string, direction string) string {
	for nsi, networks := range arr {
		if networks[0][0] == current {
			switch direction {
			case "L":
				return arr[nsi][1][0]
			case "R":
				return arr[nsi][1][1]
			}
		}
	}
	return "ZZZ"
}

func main() {

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var networks [][][]string
	var controls []string
	isFirstLine := true

	for fileScanner.Scan() {
		t := fileScanner.Text()

		if isFirstLine {
			fmt.Println("This is the first line:", t)
			controls = strings.Split(t, "")
			isFirstLine = false
			continue
		}

		if strings.TrimSpace(t) == "" {
			continue
		}

		node := strings.Split(t, " = ")
		nodeStart := strings.Trim(strings.Split(node[1], ", ")[0], "(")
		nodeEnd := strings.Trim(strings.Split(node[1], ", ")[1], ")")
		nodeName := node[0]
		networks = append(networks, [][]string{{nodeName}, {nodeStart, nodeEnd}})
	}

	var counter int
	foundEnd := false
	i := 0
	next := "AAA"

	for !foundEnd {

		counter++
		next = findNext(networks, next, controls[i])
		fmt.Println("next:", next)
		if next == "ZZZ" {
			foundEnd = true
		}

		i++
		if i == len(controls) {
			fmt.Println("resetting loop")
			i = 0
		}

	}

	// var answer int
	// for _, a := range answers {
	// 	answer += a
	// }
	fmt.Println(controls, networks, counter)

	fmt.Println("end")

}
