package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func countRanges(ranges []string) int {
	total := 0
	for _, r := range ranges {
		bounds := strings.Split(r, "-")
		lower, _ := strconv.Atoi(strings.TrimSpace(bounds[0]))
		upper, _ := strconv.Atoi(strings.TrimSpace(bounds[1]))
		total += (upper - lower + 1)
	}
	return total
}

func combineRanges(first string, second string) []string {
	bounds1 := strings.Split(first, "-")
	lower1, _ := strconv.Atoi(strings.TrimSpace(bounds1[0]))
	upper1, _ := strconv.Atoi(strings.TrimSpace(bounds1[1]))

	bounds2 := strings.Split(second, "-")
	lower2, _ := strconv.Atoi(strings.TrimSpace(bounds2[0]))
	upper2, _ := strconv.Atoi(strings.TrimSpace(bounds2[1]))

	if (lower1 <= lower2 && upper1 >= lower2) || (lower2 <= lower1 && upper2 >= lower1) {
		newLower := lower1
		if lower2 < lower1 {
			newLower = lower2
		}
		newUpper := upper1
		if upper2 > upper1 {
			newUpper = upper2
		}
		return []string{fmt.Sprintf("%d-%d", newLower, newUpper)}
	}
	return []string{first, second}
}

func removeRange(ranges []string, toDelete string) []string {
	updatedRanges := ranges
	for i, r := range ranges {
		if r == toDelete {
			updatedRanges = append(updatedRanges[:i], updatedRanges[i+1:]...)
			break
		}
	}
	return updatedRanges
}

func totalFreshIngredients(freshRanges []string) int {
	testRanges := freshRanges
	combinedRanges := []string{}
	i := 0
	for {
		r1 := testRanges[i]
		for j := i + 1; j < len(testRanges); j++ {
			r2 := testRanges[j]
			combined := combineRanges(r1, r2)
			if len(combined) == 1 {
				combinedRanges = append(combinedRanges, combined[0])
				testRanges = removeRange(testRanges, r2)
				testRanges = removeRange(testRanges, r1)
				i--
				break
			}
		}
		i++
		if i+1 == len(testRanges) {
			if len(combinedRanges) == 0 {
				break
			}
			testRanges = append(combinedRanges, testRanges...)
			combinedRanges = []string{}
			i = 0
			continue
		}
	}
	return countRanges(testRanges)
}

func isFresh(ingredient string, freshRanges []string) bool {
	for _, freshRange := range freshRanges {
		bounds := strings.Split(freshRange, "-")
		fmt.Println(bounds)
		lowerBound, _ := strconv.Atoi(strings.TrimSpace(bounds[0]))
		upperBound, _ := strconv.Atoi(strings.TrimSpace(bounds[1]))
		ingredientValue, _ := strconv.Atoi(ingredient)
		if ingredientValue >= lowerBound && ingredientValue <= upperBound {
			return true
		}
	}
	return false
}

func main() {
	file, err := os.Open("ingredients")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalf("Error closing file: %v", err)
		}
	}()
	freshRanges := []string{}
	ingredients := []string{}
	numFresh := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "-") {
			freshRanges = append(freshRanges, line)
		} else if len(line) > 0 {
			ingredients = append(ingredients, line)
		}
	}

	for _, ingredient := range ingredients {
		if isFresh(ingredient, freshRanges) {
			numFresh++
		}
	}
	totalFresh := totalFreshIngredients(freshRanges)
	fmt.Println("totalFresh: ", totalFresh)
	fmt.Println("numFresh: ", numFresh)
}
