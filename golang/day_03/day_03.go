package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type NumberNode struct {
	value string
	row   int
	col   int
}

func (n *NumberNode) getStart() int {
	return n.col
}

func (n *NumberNode) getEnd() int {
	return n.col + len(n.value) - 1
}

type SymbolNode struct {
	value string
	row   int
	col   int
}

func (s *SymbolNode) isAdjacentToNumber(number NumberNode) bool {
	if s.row < number.row-1 || s.row > number.row+1 {
		return false
	}

	if s.col >= number.getStart()-1 && s.col <= number.getEnd()+1 {
		return true
	}

	return false
}

func readFile(filename string) []string {
	actualFilePath := filepath.Join("..", "..", "inputs", "day_03", filename)
	dat, err := os.ReadFile(actualFilePath)

	if err != nil {
		panic(err)
	}

	stringDat := string(dat)

	stringDat = strings.ReplaceAll(stringDat, "\r\n", "\n")

	return strings.Split(stringDat, "\n")
}

func part_one() {
	readData := readFile("input.txt")

	var numberInProgress *NumberNode

	numbers := []NumberNode{}
	symbols := []SymbolNode{}

	for row, line := range readData {
		if line == "" {
			continue
		}

		characters := strings.Split(line, "")

		for col, character := range characters {
			if character == "" {
				continue
			}

			if _, err := strconv.Atoi(character); err == nil {
				if numberInProgress == nil {
					numberInProgress = &NumberNode{value: character, row: row, col: col}
				} else {
					numberInProgress.value += character
				}
			} else {
				// Check if we have a number in progress
				if numberInProgress != nil {
					numbers = append(numbers, *numberInProgress)
					numberInProgress = nil
				}
				if character == "." {
					continue
				}

				symbols = append(symbols, SymbolNode{value: character, row: row, col: col})
			}

		}
	}

	if numberInProgress != nil {
		numbers = append(numbers, *numberInProgress)
		numberInProgress = nil
	}

	total := 0

NUMBERS:
	for _, number := range numbers {
		for _, symbol := range symbols {
			if symbol.isAdjacentToNumber(number) {
				if actualNumber, err := strconv.Atoi(number.value); err == nil {
					total += actualNumber
					continue NUMBERS
				} else {
					panic(err)
				}
			}
		}
	}

	fmt.Println(total)
}

func part_two() {
	readData := readFile("input.txt")

	var numberInProgress *NumberNode

	numbers := []NumberNode{}
	symbols := []SymbolNode{}

	for row, line := range readData {
		if line == "" {
			continue
		}

		characters := strings.Split(line, "")

		for col, character := range characters {
			if character == "" {
				continue
			}

			if _, err := strconv.Atoi(character); err == nil {
				if numberInProgress == nil {
					numberInProgress = &NumberNode{value: character, row: row, col: col}
				} else {
					numberInProgress.value += character
				}
			} else {
				// Check if we have a number in progress
				if numberInProgress != nil {
					numbers = append(numbers, *numberInProgress)
					numberInProgress = nil
				}
				if character == "." {
					continue
				}

				symbols = append(symbols, SymbolNode{value: character, row: row, col: col})
			}

		}
	}

	if numberInProgress != nil {
		numbers = append(numbers, *numberInProgress)
		numberInProgress = nil
	}

	total := 0

	for _, symbol := range symbols {
		if symbol.value != "*" {
			continue
		}

		gearNumbers := []NumberNode{}

		for _, number := range numbers {
			if symbol.isAdjacentToNumber(number) {
				gearNumbers = append(gearNumbers, number)
			}
		}

		if len(gearNumbers) == 2 {
			firstNumber, err := strconv.Atoi(gearNumbers[0].value)
			if err != nil {
				panic(err)
			}
			secondNumber, err := strconv.Atoi(gearNumbers[1].value)
			if err != nil {
				panic(err)
			}
			total += firstNumber * secondNumber
		}
	}

	fmt.Println(total)
}

func main() {
	part_one()
    part_two()
}
