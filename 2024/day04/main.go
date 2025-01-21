package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// [
//	[M M M S X X M A S M]
//	[M S A M X M S M S A]
//	[A M X S X M A A M M]
//	[M S A M A S M S M X]
//	[X M A S A M X A M M]
//	[X X A M M X X A M A]
//	[S M S M S A S X S S]
//	[S A X A M A S A A A]
//	[M A M M M X M M M M]
//	[M X M X A X M A S X]
// ]

//func isNearby(crossword, currX, currY, nextVal) bool {

//}

func main() {

	readFile, err := os.Open("test.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var crossword [][]string

	// build 2d array
	for fileScanner.Scan() {
		t := fileScanner.Text()
		f := strings.Split(t, "")
		crossword = append(crossword, f)
	}

	for i := 0; i < len(crossword); i++ {
		for j := 0; j < len(crossword); j++ {
			fmt.Printf("Element at [%d][%d]: %s\n", i, j, crossword[i][j])
		}
	}

	fmt.Println(crossword)
	fmt.Println("end")

}
