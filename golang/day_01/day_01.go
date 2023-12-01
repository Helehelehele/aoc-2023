package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func readFile(filename string) []string {
	actualFilePath := filepath.Join("..", "..", "inputs", "day_01", filename)
	dat, err := os.ReadFile(actualFilePath)

	if err != nil {
		panic(err)
	}

	stringDat := string(dat)

	stringDat = strings.ReplaceAll(stringDat, "\r\n", "\n")

	return strings.Split(stringDat, "\n")
}

func processLine(line string) int {
	chars := strings.Split(line, "")

	var firstDigit, lastDigit int
	firstDigit = -1

	for _, char := range chars {
		digit, err := strconv.Atoi(char)

		if err != nil {
			continue
		}

		if firstDigit == -1 {
			firstDigit = digit
		}

		lastDigit = digit
	}

	return firstDigit*10 + lastDigit
}

func part_one() {
	readData := readFile("input.txt")

	var total int

	for _, line := range readData {
		if line == "" {
			continue
		}
		current := processLine(line)
		total += current
	}

	fmt.Println(total)
}

func processLinePartTwo(line string) int {
	line = strings.ReplaceAll(line, "one", "o1e")
	line = strings.ReplaceAll(line, "two", "t2o")
	line = strings.ReplaceAll(line, "three", "t3e")
	line = strings.ReplaceAll(line, "four", "f4r")
	line = strings.ReplaceAll(line, "five", "f5e")
	line = strings.ReplaceAll(line, "six", "s6x")
	line = strings.ReplaceAll(line, "seven", "s7n")
	line = strings.ReplaceAll(line, "eight", "e8t")
	line = strings.ReplaceAll(line, "nine", "n9e")

	return processLine(line)
}

func part_two() {
	readData := readFile("input.txt")

	var total int

	for _, line := range readData {
		if line == "" {
			continue
		}
		current := processLinePartTwo(line)
		total += current
	}

	fmt.Println(total)
}

func main() {
	part_one()
	part_two()
}
