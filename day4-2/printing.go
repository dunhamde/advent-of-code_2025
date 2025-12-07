package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func isInBounds(paperMat [][]bool, x int, y int) bool {
	maxIndexHoriz := len(paperMat[0]) - 1
	maxIndexVert := len(paperMat) - 1
	if x > maxIndexHoriz || x < 0 {
		return false
	} else if y > maxIndexVert || y < 0 {
		return false
	}
	return true
}

// x-1, y
// x-1, y-1
// x, y-1
// x+1, y+1
// x+1, y
// x+1, y-1
// x, y+1
// x-1, y+1

func checkCell(paperMat [][]bool, x int, y int) bool {
	numRolls := 0
	fmt.Printf("%d, %d \n", x, y)
	if isInBounds(paperMat, x-1, y) && paperMat[x-1][y] {
		numRolls++
	}
	if isInBounds(paperMat, x-1, y-1) && paperMat[x-1][y-1] {
		numRolls++
	}
	if isInBounds(paperMat, x, y-1) && paperMat[x][y-1] {
		numRolls++
	}
	if isInBounds(paperMat, x+1, y+1) && paperMat[x+1][y+1] {
		numRolls++
	}
	if isInBounds(paperMat, x+1, y) && paperMat[x+1][y] {
		numRolls++
	}
	if isInBounds(paperMat, x+1, y-1) && paperMat[x+1][y-1] {
		numRolls++
	}
	if isInBounds(paperMat, x, y+1) && paperMat[x][y+1] {
		numRolls++
	}
	if isInBounds(paperMat, x-1, y+1) && paperMat[x-1][y+1] {
		numRolls++
	}
	// fmt.Println(numRolls < 4)
	return numRolls < 4
}

func main() {
	file, err := os.Open("paper")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalf("Error closing file: %v", err)
		}
	}()
	totalRolls := 0
	paperMat := [][]bool{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		paperLine := scanner.Text()
		paperLineArr := []bool{}
		for _, roll := range paperLine {
			if string(roll) == "@" {
				paperLineArr = append(paperLineArr, true)
			} else {
				paperLineArr = append(paperLineArr, false)
			}
		}
		paperMat = append(paperMat, paperLineArr)
	}

	fmt.Println(paperMat[0])
	// fmt.Println("mat len: ", len(paperMat))
	for x, rollRow := range paperMat {
		for y, _ := range rollRow {
			// fmt.Println("row len: ", len(rollRow))
			// fmt.Println("roll valyue: ", paperMat[x][y])
			if paperMat[x][y] && checkCell(paperMat, x, y) {
				totalRolls++
			}
		}
	}
	fmt.Println(totalRolls)
}

// nested for loops through the whole matrix
// on each iteration check the cells surrounding 8 cells

// x-1, y
// x-1, y-1
// x, y-1
// x+1, y+1
// x+1, y
// x+1, y-1
// x, y+1
// x-1, y+1
