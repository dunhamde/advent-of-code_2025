package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func findInitialBeamCol(line string) int {
	for col, char := range line {
		if char == 'S' {
			return col
		}
	}
	return -1
}
func processLine(line string, beamCols *map[int]bool, totalSplits *int) {
	for colIdx, char := range line {

		// we found a splitter
		if char == '^' {
			// check if there is a beam coming into this splitter
			if (*beamCols)[colIdx] {
				// beam splits here
				*totalSplits++
				// remove the incoming beam
				delete(*beamCols, colIdx)
				// add outgoing beams
				(*beamCols)[colIdx-1] = true
				(*beamCols)[colIdx+1] = true
			}
		}
	}

}

func main() {
	file, err := os.Open("diagram")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("failed to close file: %s", err)
		}
	}()
	scanner := bufio.NewScanner(file)
	totalSplits := 0
	initialCol := -1
	beamCols := make(map[int]bool)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		if initialCol == -1 {
			initialCol = findInitialBeamCol(line)
			beamCols[initialCol] = true
		} else {
			//process line to find beam splits
			processLine(line, &beamCols, &totalSplits)
		}

	}
	fmt.Printf("Total beam splits: %d\n", totalSplits)
	fmt.Println("-----")
	fmt.Println("beamCols:", beamCols)

	// doPartOne(file)
	// doPartTwo(file)
}

// part one:
// compute total number of times the beam splits in the laboratory diagram

// part two:
//
