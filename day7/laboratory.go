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

func memoizationKey(col int, row int) string {
	return fmt.Sprintf("%d,%d", col, row)
}

func getMemoizedValue(col int, row int) (int, bool) {
	key := memoizationKey(col, row)
	value, exists := memoizationCache[key]
	return value, exists
}

func setMemoizedValue(col int, row int, value int) {
	key := memoizationKey(col, row)
	memoizationCache[key] = value
}

func sumTimelines(currentCol int, lines []string, currentRowIdx int) int {
	// find next splitter in the current column
	splitterIdx := -1
	for rowIdx, line := range lines {
		char := line[currentCol]

		if char == '^' {
			splitterIdx = rowIdx
			// splitter found, recursively take both outgoing beams
			// check memoization cache
			if cachedValue, exists := getMemoizedValue(currentCol, currentRowIdx+rowIdx); exists {
				fmt.Printf("Using cached value for col=%d, row=%d: %d\n", currentCol, currentRowIdx+rowIdx, cachedValue)
				return cachedValue
			}
			leftTimelines := sumTimelines(currentCol-1, lines[rowIdx+1:], currentRowIdx+rowIdx+1)
			rightTimelines := sumTimelines(currentCol+1, lines[rowIdx+1:], currentRowIdx+rowIdx+1)
			setMemoizedValue(currentCol, currentRowIdx+rowIdx, leftTimelines+rightTimelines)
			return leftTimelines + rightTimelines
		}
	}

	setMemoizedValue(currentCol, splitterIdx+currentRowIdx, 1)
	// no splitter found in this column, beam reaches bottom
	return 1
}
func doPartTwo(initialCol int, lines []string) {
	totalTimelines := sumTimelines(initialCol, lines, 1)
	fmt.Printf("Total timelines reaching bottom: %d\n", totalTimelines)
}

var memoizationCache = make(map[string]int)

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
