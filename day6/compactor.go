package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getColumnNumbers(mathProblems [][]string, col int) []int {
	columnNumbers := []int{}
	for _, row := range mathProblems {
		if col < len(row) {
			num, err := strconv.Atoi(row[col])
			if err == nil {
				columnNumbers = append(columnNumbers, num)
			}
		}
	}
	return columnNumbers
}

func computeMathColumn(mathProblems [][]string, col int) int {
	total := 0
	if mathProblems[len(mathProblems)-1][col] == "+" {
		for _, num := range getColumnNumbers(mathProblems, col) {
			total += num
		}

	} else if mathProblems[len(mathProblems)-1][col] == "*" {
		total = 1
		for _, num := range getColumnNumbers(mathProblems, col) {
			total *= num
		}
	}
	return total
}

func computeMathColumns(mathProblems [][]string) int {
	total := 0
	for col := 0; col < len(mathProblems[0]); col++ {
		total += computeMathColumn(mathProblems, col)
	}
	return total
}

func main() {
	file, err := os.Open("problems")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalf("Error closing file: %v", err)
		}
	}()

	scanner := bufio.NewScanner(file)
	mathProblems := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		trimmedLine := []string{}
		splitLine := strings.Split(line, " ")
		for _, str := range splitLine {
			_, err := strconv.Atoi(str)
			if err == nil || str == "+" || str == "*" {
				trimmedLine = append(trimmedLine, str)
			}
		}
		mathProblems = append(mathProblems, trimmedLine)
	}
	fmt.Println("len(mathProblems):", len(mathProblems[0]))
	fmt.Println("mathProblems:", mathProblems)
	fmt.Println(computeMathColumns(mathProblems))
}
