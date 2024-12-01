package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"reflect"
	"strings"
)

func checkRow(image [][]string, rowIndex int) bool {
	for i := 0; i < len(image[0]); i++ {
		if image[rowIndex][i] != "." {
			fmt.Printf("Row %d in Column %d doesn't contain .\n", rowIndex, i)
			return false
		}
	}
	return true
}

func checkColumn(image [][]string, columnIndex int) bool {
	for i := 0; i < len(image); i++ {
		if image[i][columnIndex] != "." {
			fmt.Printf("Column %d in Row %d doesn't contain .\n", columnIndex, i)
			return false
		}
	}
	return true
}

func inArray(num int, arr []int) bool {
	for _, n := range arr {
		if n == num {
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

	var image [][]string
	var galaxies [][]int

	// build 2d array
	for fileScanner.Scan() {
		t := fileScanner.Text()
		f := strings.Split(t, "")
		image = append(image, f)
	}

	// find galaxies & empty rows
	var emptyRows []int
	var emptyCols []int
	for ri, row := range image {
		if checkRow(image, ri) {
			emptyRows = append(emptyRows, ri)
		}
		for ci, s := range row {
			if s == "#" {
				coordinates := []int{ri, ci}
				galaxies = append(galaxies, coordinates)
			}
		}
	}
	// check empty colns
	for i := 0; i < len(image); i++ {
		if checkColumn(image, i) {
			emptyCols = append(emptyCols, i)
		}
	}

	fmt.Println(galaxies, emptyRows, emptyCols)

	distances := 0

	// I'm not sure how to not include an already calculate pair?
	// I could do another array that contains already calculates pairs and just ignore those in each iteration...
	// too lazy though, just assume each is calculated twice and divide by two
	for _, galaxy := range galaxies {
		for _, innerGalaxy := range galaxies {
			// skip inner galaxy if it is the current galaxy
			if !reflect.DeepEqual(galaxy, innerGalaxy) {
				// check rows
				start := int(math.Min(float64(galaxy[0]), float64(innerGalaxy[0]))) + 1
				end := int(math.Max(float64(galaxy[0]), float64(innerGalaxy[0])))
				for i := start; i <= end; i++ {
					if !inArray(i, emptyRows) {
						distances += 1
					} else {
						distances += 1000000
					}
				}

				// check columns
				start = int(math.Min(float64(galaxy[1]), float64(innerGalaxy[1]))) + 1
				end = int(math.Max(float64(galaxy[1]), float64(innerGalaxy[1])))
				for i := start; i <= end; i++ {
					if !inArray(i, emptyCols) {
						distances += 1
					} else {
						distances += 1000000
					}
				}
			}
		}
	}

	fmt.Println(distances / 2)

	readFile.Close()

	fmt.Println("end")

}
