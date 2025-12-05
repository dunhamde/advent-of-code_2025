package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Open the file for reading.
	file, err := os.Open("data_sequence")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	// Defer closing the file until the function exits.
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalf("Error closing file: %v", err)
		}
	}()

	currentPosition := 50
	solution := 0

	// Read the file line by line.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Parse the line to get our direction and distance
		line := scanner.Text()
		direction := line[0:1]
		distanceStr := line[1:]

		distanceNumRaw, err := strconv.Atoi(distanceStr)
		if err != nil {
			log.Fatalf("Failed to convert string to int: %v", err)
		}

		// Add full dial spins to the solution total and
		// remove full spins from the distance we want to move

		distanceNum := distanceNumRaw % 100

		// fmt.Printf("Extra spins: %d\n", (distanceNumRaw-distanceNum)/100)
		solution += (distanceNumRaw - distanceNum) / 100

		if direction == "L" {
			distanceNum = distanceNum * -1
		}

		origPos := currentPosition
		// Calculate raw dial position
		currentPosition += distanceNum

		// Compute "real" dial position
		if currentPosition < 0 {
			currentPosition = 100 + currentPosition

			// If we started on zero, we already added that to the solution total
			if origPos != 0 {
				solution++
			}
		}
		if currentPosition > 99 {
			currentPosition = currentPosition - 100

			// Passed zero, add it to the solution total
			if currentPosition > 0 {
				solution++
			}
		}
		if currentPosition == 0 {
			solution++
		}
	}
	fmt.Printf("solution: %d\n", solution)
}
