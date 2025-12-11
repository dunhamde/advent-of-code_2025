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

func doPartOne(file *os.File) {
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
	total := computeMathColumns(mathProblems)
	fmt.Println("total:", total)
}

func processColumn(lines []string, start int, end int) int {
	columnNumbers := make([]string, end-start)
	for rowIdx, line := range lines {
		if rowIdx == len(lines)-1 {
			continue
		}
		for i := start; i < end; i++ {
			char := line[i]
			if char != ' ' {
				columnNumbers[i-start] += string(char)
			}
		}
	}
	columnNumbersInt := []int{}
	for _, strNum := range columnNumbers {
		num, _ := strconv.Atoi(strNum)
		columnNumbersInt = append(columnNumbersInt, num)
	}
	total := 0
	if lines[len(lines)-1][start] == '+' {
		for _, num := range columnNumbersInt {
			total += num
		}
	} else if lines[len(lines)-1][start] == '*' {
		total = 1
		for _, num := range columnNumbersInt {
			total *= num
		}
	}
	return total
}

func doPartTwo(file *os.File) {
	scanner := bufio.NewScanner(file)
	// identify start and end of each column
	// first get the x positions of each operator in the last line
	operatorPositions := []int{}
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	lastLine := lines[len(lines)-1]
	for idx, char := range lastLine {
		if char == '+' || char == '*' {
			operatorPositions = append(operatorPositions, idx)
		}
	}

	// use operator positions to process each column
	total := 0
	for opIdx, opPos := range operatorPositions {
		opIdxEnd := 0
		if opIdx == len(operatorPositions)-1 {
			opIdxEnd = len(lines[len(lines)-2])
		} else {
			opIdxEnd = operatorPositions[opIdx+1] - 1
		}
		columnTot := processColumn(lines, opPos, opIdxEnd)
		total += columnTot
	}
	fmt.Println("total:", total)
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

	// doPartOne(file)
	doPartTwo(file)
}

// part two:
// each operator indicates where a column of numbers starts
// the next operator is 2 spaces away from the last number in the previous column
