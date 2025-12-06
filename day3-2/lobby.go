package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Find the highest value int in batteryStr, with at least toRemain
// digits to the right
func findHighestWithRemains(batteryStr string, toRemain int) int {
	highestValue := 0
	highestIndex := -1
	for i := len(batteryStr) - (toRemain + 1); i >= 0; i-- {
		testStr := string(batteryStr[i])
		testInt, _ := strconv.Atoi(testStr)
		if testInt >= highestValue {
			highestValue = testInt
			highestIndex = i
		}
	}
	return highestIndex
}

func findHighestOfNumDigits(batteryStr string, toRemain int) string {
	// We found the last digit we needed on the last call, start combining the answer
	if toRemain == -1 {
		return ""
	} else {
		// Find the next digit in the sequence
		idx := findHighestWithRemains(batteryStr, toRemain)
		digitStr := string(batteryStr[idx])
		return digitStr + findHighestOfNumDigits(batteryStr[idx+1:], toRemain-1)
	}
}

func main() {
	file, err := os.Open("../input-data/battery")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalf("Error closing file: %v", err)
		}
	}()

	joltage := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		batteryLine := scanner.Text()
		highestStr := findHighestOfNumDigits(batteryLine, 11)
		joltageValue, _ := strconv.Atoi(highestStr)
		joltage += joltageValue
	}
	fmt.Println(joltage)
}
