package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strings"
)

type Card struct {
	winningNumbers map[string]bool
	scratchNumbers map[string]bool
}

func (c *Card) countWinningNumbers() int {
	sIntersection := map[string]bool{}

	for k := range c.winningNumbers {
		if _, ok := c.scratchNumbers[k]; ok {
			sIntersection[k] = true
		}
	}

	return len(sIntersection)
}

func readFile(filename string) []string {
	actualFilePath := filepath.Join("..", "..", "inputs", "day_04", filename)
	dat, err := os.ReadFile(actualFilePath)

	if err != nil {
		panic(err)
	}

	stringDat := string(dat)

	stringDat = strings.ReplaceAll(stringDat, "\r\n", "\n")

	splitFn := func(c rune) bool {
		return c == '\n'
	}

	return strings.FieldsFunc(stringDat, splitFn)
}

func parseLine(line string) Card {
	splitLine := strings.Split(line, ":")

	numbers := splitLine[1]

	splitNumbers := strings.Split(numbers, "|")

	winningNumbers := map[string]bool{}

	splitFn := func(c rune) bool {
		return c == ' '
	}

	for _, number := range strings.FieldsFunc(splitNumbers[0], splitFn) {
		winningNumber := strings.TrimSpace(number)
		winningNumbers[winningNumber] = true
	}

	scratchNumbers := map[string]bool{}

	for _, number := range strings.FieldsFunc(splitNumbers[1], splitFn) {
		scratchNumber := strings.TrimSpace(number)
		scratchNumbers[scratchNumber] = true
	}

	return Card{winningNumbers, scratchNumbers}
}

func part_one() {
	readData := readFile("input.txt")

	total := 0

	for _, line := range readData {
		card := parseLine(line)

		winningCount := card.countWinningNumbers()

		if winningCount < 1 {
			continue
		}

		total += int(math.Pow(2, float64(winningCount-1)))
	}

	fmt.Println("Part one:", total)
}

func pop(slice *[]int) (int, []int) {
	head := (*slice)[0]
	tail := (*slice)[1:]

	return head, tail
}

func winCopies(
	winCountMap *map[int]int,
	cards *[]int,
) int {
	total := 0

	for len(*cards) > 0 {
		head, tail := pop(cards)
		*cards = tail

		winCountForCard := (*winCountMap)[head]

		if winCountForCard == 0 {
			total += 1
			continue
		}

		wonCards := []int{}

		for i := 1; i <= winCountForCard; i++ {
			wonCard := head + i

			if _, ok := (*winCountMap)[wonCard]; ok {
				wonCards = append(wonCards, wonCard)
			} else {
				break
			}
		}

		*cards = append(wonCards, *cards...)
		total += 1
	}

	return total
}

func part_two() {
	readData := readFile("input.txt")

	winCountMap := map[int]int{}

	cards := []int{}

	for index, line := range readData {
		card := parseLine(line)

		winningCount := card.countWinningNumbers()

		winCountMap[index+1] = winningCount

		cards = append(cards, index+1)
	}

	totalCount := winCopies(&winCountMap, &cards)

	fmt.Println("Part two:", totalCount)

}

func main() {
	part_one()
	part_two()
}
