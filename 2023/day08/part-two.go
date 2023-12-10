package main

// I had to get help on this one, tried to brute force. Needed the tip of finding the lcm instead
// unfortunately go doesn't have any built in helper functions that I could find.

import (
	"bufio"
	"fmt"
	"math/big"
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

func gcd(a, b *big.Int) *big.Int {
	for b.Sign() != 0 {
		a, b = b, new(big.Int).Mod(a, b)
	}
	return a
}

func lcm(a, b *big.Int) *big.Int {
	g := gcd(a, b)
	return new(big.Int).Div(new(big.Int).Mul(a, b), g)
}

func calcLCM(numbers []*big.Int) *big.Int {
	result := new(big.Int)
	result.Set(numbers[0])

	for i := 1; i < len(numbers); i++ {
		result = lcm(result, numbers[i])
	}

	return result
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

	var starters []int
	line := 0

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
		isStarter := strings.HasSuffix(nodeName, "A")
		if isStarter {
			starters = append(starters, line)
		}
		line++
		networks = append(networks, [][]string{{nodeName}, {nodeStart, nodeEnd}})
	}

	var nexts []*big.Int

	fmt.Println("networks:", networks, "starters:", starters, "nexts:", nexts)

	for _, start := range starters {
		exitLoop := false
		i := 0
		var counter big.Int
		var next []string
		current := networks[start][0][0]
		fmt.Println("current:", current)
		for !exitLoop {
			counter.Add(&counter, big.NewInt(1))

			current = findNext(networks, current, controls[i])

			if strings.HasSuffix(current, "Z") {
				next = append(next, current)
				// back to start
				if next[0] == current {
					exitLoop = true
				}
			}

			fmt.Println("current:", current, "counter:", counter.String())

			i++
			if i == len(controls) {
				fmt.Println("resetting loop")
				i = 0
			}
		}
		fmt.Println("next:", next)
		nexts = append(nexts, new(big.Int).Set(&counter))

	}

	fmt.Println("nexts:", nexts)
	answer := calcLCM(nexts)

	fmt.Println(answer.String())

	fmt.Println("end")

}
