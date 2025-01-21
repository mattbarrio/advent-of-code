package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// StackItem represents an item in our processing stack
type StackItem struct {
	n   int
	op  string
	res int
	i   int
}

// printSolution wraps result printing
func printSolution(result int) {
	fmt.Printf("Solution: %d\n", result)
}

func part1(values []int, equations [][]int) int {
	p1 := 0

	for idx, v := range values {
		stack := make([]StackItem, 0)
		equation := equations[idx]
		i := 0

		stack = append(stack, StackItem{equation[i], "+", 0, 0})
		stack = append(stack, StackItem{equation[i], "*", 1, 0})

		for len(stack) > 0 {
			lastIdx := len(stack) - 1
			item := stack[lastIdx]
			stack = stack[:lastIdx]

			n, op, res, i := item.n, item.op, item.res, item.i

			if op == "+" {
				res += n
			} else {
				res *= n
			}
			i++

			if i == len(equation) {
				if res == v {
					p1 += v
					break
				}
			} else {
				if i < len(equation) {
					stack = append(stack, StackItem{equation[i], "+", res, i})
					stack = append(stack, StackItem{equation[i], "*", res, i})
				}
			}
		}
	}

	printSolution(p1)
	return p1
}

func part2(values []int, equations [][]int) int {
	p2 := 0

	for idx, v := range values {
		stack := make([]StackItem, 0)
		equation := equations[idx]
		i := 0

		stack = append(stack, StackItem{equation[i], "+", 0, 0})
		stack = append(stack, StackItem{equation[i], "*", 1, 0})
		stack = append(stack, StackItem{equation[i], "||", 0, 0})

		for len(stack) > 0 {
			lastIdx := len(stack) - 1
			item := stack[lastIdx]
			stack = stack[:lastIdx]

			n, op, res, i := item.n, item.op, item.res, item.i

			if op == "+" {
				res += n
			} else if op == "||" {
				// Convert to string, concatenate, and back to int
				resStr := fmt.Sprintf("%d%d", res, n)
				res, _ = strconv.Atoi(resStr)
			} else {
				res *= n
			}
			i++

			if i == len(equation) {
				if res == v {
					p2 += v
					break
				}
			} else {
				if i < len(equation) {
					stack = append(stack, StackItem{equation[i], "+", res, i})
					stack = append(stack, StackItem{equation[i], "*", res, i})
					stack = append(stack, StackItem{equation[i], "||", res, i})
				}
			}
		}
	}

	printSolution(p2)
	return p2
}

func readInputFile(filename string) ([]int, [][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var values []int
	var equations [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")

		// Parse value
		value, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, nil, err
		}
		values = append(values, value)

		// Parse equation
		numStrs := strings.Fields(strings.TrimSpace(parts[1]))
		equation := make([]int, len(numStrs))
		for i, numStr := range numStrs {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, nil, err
			}
			equation[i] = num
		}
		equations = append(equations, equation)
	}

	return values, equations, scanner.Err()
}

func main() {
	values, equations, err := readInputFile("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	part1(values, equations)
	part2(values, equations)
}
