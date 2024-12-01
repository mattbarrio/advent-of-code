package main

// This was frustrating. I wish I had the second part of the puzzle, I don't think it is
// solvable with the way I implemented the first part so this needs to be redone.
// If I come back to this puzzle, use the gear to scan around for matching numbers,
// might need to dedup if I do that though.

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func seenCord(seenCords [][]int, cord []int) (seen bool) {
	for _, s := range seenCords {
		if s[0] == cord[0] && s[1] == cord[1] {
			fmt.Println("skipping:", s[0], s[1])
			return true
		}
	}
	return false
}

func main() {

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var schematic [][]string

	// build 2d array
	for fileScanner.Scan() {
		t := fileScanner.Text()
		f := strings.Split(t, "")
		schematic = append(schematic, f)
	}

	fmt.Println(schematic)

	answers := []int{}

	// find symbols
	for r, row := range schematic {
		for c, char := range row {
			if char == "*" {
				var sets []int
				var seenCords [][]int
				// check the nine spots around the symbol
				rowOffsets := []int{-1, 0, 1}
				colOffsets := []int{-1, 0, 1}
				// 6 1 8 . .
				// . . * . .
				// 7 8 6 . .
				for _, rowOffset := range rowOffsets {
					for _, colOffset := range colOffsets {
						ri, ci := r+rowOffset, c+colOffset
						var num string
						// Check if the new position is within bounds
						if ri >= 0 && ri < len(schematic) && ci >= 0 && ci < len(schematic[0]) {
							if unicode.IsDigit([]rune(schematic[ri][ci])[0]) && !seenCord(seenCords, []int{ri, ci}) {
								seenCords = append(seenCords, []int{ri, ci})
								num += schematic[ri][ci]
								fmt.Println("r:", r, "c:", c, "ri:", ri, "ci:", ci, "num:", num)
								for ci > 0 && unicode.IsDigit([]rune(schematic[ri][ci-1])[0]) {
									fmt.Println("left:", ci, schematic[ri][ci])
									ci--
									seenCords = append(seenCords, []int{ri, ci})
									num = schematic[ri][ci] + num
								}
								fmt.Println("ci:", ci, "num:", num, "seen:", seenCords)
								for ci+1 < len(schematic[0]) && unicode.IsDigit([]rune(schematic[ri][ci+1])[0]) {
									if seenCord(seenCords, []int{ri, ci + 1}) {
										ci++
										continue
									}
									fmt.Println("right:", ci, schematic[ri][ci])
									ci++
									seenCords = append(seenCords, []int{ri, ci})
									num = num + schematic[ri][ci]
								}
								set, _ := strconv.Atoi(num)
								sets = append(sets, set)
								fmt.Println("ci:", ci, "num:", num, "sets:", sets, "seen", seenCords)
							}
						}
					}
				}
				if len(sets) == 2 {
					s := sets[0] * sets[1]
					fmt.Println(s, sets[0], sets[1])
					answers = append(answers, s)
				}
			}
		}
	}

	answer := 0
	for _, r := range answers {
		answer += r
	}

	fmt.Println(answer)

	readFile.Close()

	fmt.Println("end")

}
