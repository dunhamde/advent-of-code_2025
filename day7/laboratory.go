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

func sumTimelines(currentCol int, lines []string) int {
	// base case: if we reach the bottom of the diagram
	if len(lines) == 0 {
		return 1
	}
	// find next splitter in the current column
	for rowIdx, line := range lines {
		char := line[currentCol]
		if char == '^' {
			// splitter found, recursively take both outgoing beams
			leftTimelines := sumTimelines(currentCol-1, lines[rowIdx+1:])
			rightTimelines := sumTimelines(currentCol+1, lines[rowIdx+1:])
			return leftTimelines + rightTimelines
		}
	}
	// no splitter found in this column, beam reaches bottom
	return 1
}
func doPartTwo(initialCol int, lines []string) {
	totalTimelines := sumTimelines(initialCol, lines)
	fmt.Printf("Total timelines reaching bottom: %d\n", totalTimelines)
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
	lines := []string{}
	beamCols := make(map[int]bool)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
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
	doPartTwo(initialCol, lines[1:]) // pass lines excluding the first line with 'S'

}

// part one:
// compute total number of times the beam splits in the laboratory diagram

// part two:
// use recursion to take a beam until it reaches the next splitter or bottom
// when the beam reaches a splitter, recursively take the two outgoing beams
// when a beam reaches the bottom, return 1
// sum up all the returns from the bottom to get total splits
