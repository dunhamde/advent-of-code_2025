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
	file, err := os.Open("../input-data/battery")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	// Defer closing the file until the function exits.
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalf("Error closing file: %v", err)
		}
	}()

	joltage := 0

	// Read the file line by line.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		batteryLine := scanner.Text()
		fmt.Println(batteryLine)
		lineLen := len(batteryLine)

		joltageValue := 0

		for i := 0; i < lineLen; i++ {
			for j := i + 1; j < lineLen; j++ {
				joltageTestStr := string(batteryLine[i]) + string(batteryLine[j])
				joltageTestValue, _ := strconv.Atoi(joltageTestStr)
				if joltageTestValue > joltageValue {
					joltageValue = joltageTestValue
				}
			}
		}
		joltage += joltageValue
	}
	fmt.Println("joltage: ", joltage)
}
