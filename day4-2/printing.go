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

func checkCell(paperMat [][]bool, x int, y int) bool {
	numRolls := 0
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
	return numRolls < 4
}

func main() {
	file, err := os.Open("../day4/paper")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalf("Error closing file: %v", err)
		}
	}()
	removedRolls := 0
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

	removedRollsIter := 0
	for {
		for x, rollRow := range paperMat {
			for y, _ := range rollRow {
				if paperMat[x][y] && checkCell(paperMat, x, y) {
					paperMat[x][y] = false
					removedRollsIter++
				}
			}
		}
		if removedRollsIter == 0 {
			break
		} else {
			removedRolls += removedRollsIter
		}
		removedRollsIter = 0
	}
	fmt.Println(removedRolls)
}
