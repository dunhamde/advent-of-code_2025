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

// Return array contain numbers where num digits % n == 0
func getZeroModArray(num int) []int {
	numLength := getNumDigits(num)
	mods := []int{1}

	for i := 2; i <= numLength/2; i++ {
		if numLength%i == 0 {
			mods = append(mods, i)
		}
	}
	return mods
}

func getNumDigits(x int) int {
	s := strconv.Itoa(x)
	return len(s)
}

func incrementRepeat(repeatNum int, numDigits int) int {
	numIncrements := len(strconv.Itoa(repeatNum)) / numDigits
	repeatNumStr := strconv.Itoa(repeatNum)
	repeatToIncrement := repeatNumStr[:numDigits]
	repeatToIncrementInt, _ := strconv.Atoi(repeatToIncrement)
	repeatToIncrementInt++

	updatedRepeatNum := ""
	for i := 0; i < numIncrements; i++ {
		updatedRepeatNum += strconv.Itoa(repeatToIncrementInt)
	}
	if len(updatedRepeatNum) > len(repeatNumStr) {
		return getNextRepeat(powInt(10, len(repeatNumStr)), numDigits)
	}
	incrementedRepeatNum, _ := strconv.Atoi(updatedRepeatNum)

	return incrementedRepeatNum

}

func getNextRepeat(x int, numDigits int) int {
	baseStr := strconv.Itoa(x)

	if len(baseStr)%numDigits != 0 || len(baseStr) == 1 {
		x = powInt(10, len(baseStr))
		baseStr = strconv.Itoa((x))
	}

	mostSigRepeatStr := string(baseStr[0:numDigits])
	initialRepeatStr := ""

	for {
		if len(initialRepeatStr+mostSigRepeatStr) > len(baseStr) {
			break
		}
		initialRepeatStr += mostSigRepeatStr
	}

	initialRepeat, _ := strconv.Atoi(initialRepeatStr)

	if initialRepeat < x {
		mostSigRepeatInt, _ := strconv.Atoi(mostSigRepeatStr)
		mostSigRepeatInt++
		mostSigRepeatStr := strconv.Itoa(mostSigRepeatInt)
		initialRepeatStr = ""
		for {
			if len(initialRepeatStr+mostSigRepeatStr) > len(baseStr) {
				break
			}
			initialRepeatStr += mostSigRepeatStr
		}

		initialRepeat, _ := strconv.Atoi(initialRepeatStr)
		return initialRepeat

	}
	return initialRepeat
}

func main() {
	file, err := os.Open("product_id_ranges")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

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
		fmt.Printf("[[%d - %d]]\n", minInt, maxInt)

		minIntMods := getZeroModArray(minInt)
		maxIntMods := getZeroModArray(maxInt)

		modMap := make(map[int]int)
		addendMap := make(map[int]int)

		for _, mod := range minIntMods {
			modMap[mod] = mod
		}
		for _, mod := range maxIntMods {
			modMap[mod] = mod
		}

		modIntersect := []int{}
		for k := range modMap {
			modIntersect = append(modIntersect, k)
		}
		fmt.Println("mods to check: ", modIntersect)

		for _, mod := range modIntersect {
			fmt.Println("mod: ", mod)
			repeat := getNextRepeat(minInt, mod)
			for repeat >= minInt && repeat <= maxInt {
				fmt.Printf(" %d ", repeat)
				addendMap[repeat] = repeat
				repeat = incrementRepeat(repeat, mod)
			}
		}
		for addend := range addendMap {
			total += addend
		}
		fmt.Println("\n-----")

	}

	fmt.Println("total: ", total)
}
