package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func main() {
	// Open the file
	file, err := os.Open("product_id_ranges")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close() // Ensure the file is closed when the function exits

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	id_ranges := strings.Split(line, ",")
	fmt.Println(id_ranges)

	total := 0

	for _, id_range := range id_ranges {
		minMax := strings.Split(id_range, "-")
		min := minMax[0]
		minInt, _ := strconv.Atoi(min)
		max := minMax[1]
		maxInt, _ := strconv.Atoi(max)

		firstRepeat := 0
		if len(min)%2 == 0 {
			middle := len(min) / 2
			halfMin := min[:middle]
			convertedInt, _ := strconv.Atoi(halfMin + halfMin)
			if convertedInt < minInt {
				halfMinInt, _ := strconv.Atoi(halfMin)
				halfMinInt++
				halfMin = strconv.Itoa(halfMinInt)
				updatedMinRepeat, _ := strconv.Atoi(halfMin + halfMin)
				firstRepeat = updatedMinRepeat
			} else {
				firstRepeat = convertedInt
			}

		} else {
			baseAboveMin := powInt(10, len(min))
			baseAboveMinStr := strconv.Itoa(baseAboveMin)
			middle := len(baseAboveMinStr) / 2
			half := baseAboveMinStr[:middle]
			convertedInt, _ := strconv.Atoi(half + half)
			firstRepeat = convertedInt
		}

		if firstRepeat >= minInt && firstRepeat <= maxInt {
			currentRepeat := firstRepeat
			for currentRepeat <= maxInt {
				fmt.Printf(" %d ", currentRepeat)
				total += currentRepeat
				currentRepeatStr := strconv.Itoa(currentRepeat)

				half := currentRepeatStr[:len(currentRepeatStr)/2]
				halfInt, _ := strconv.Atoi(half)
				newHalf := halfInt + 1
				newRepeat, _ := strconv.Atoi(strconv.Itoa(newHalf) + strconv.Itoa(newHalf))
				currentRepeat = newRepeat
			}
		}
		fmt.Println("\n--------------------------------\n")

	}
	fmt.Println(total)

}
