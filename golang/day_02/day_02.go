package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Reveal struct {
	count int
	color string
}

type Game struct {
	gameId  int
	reveals []Reveal
}

func readFile(filename string) []string {
	actualFilePath := filepath.Join("..", "..", "inputs", "day_02", filename)
	dat, err := os.ReadFile(actualFilePath)

	if err != nil {
		panic(err)
	}

	stringDat := string(dat)

	stringDat = strings.ReplaceAll(stringDat, "\r\n", "\n")

	return strings.Split(stringDat, "\n")
}

func processLine(line string) Game {
	splitLine := strings.SplitN(line, ":", 2)

	gameId := splitLine[0]
	result := splitLine[1]

	gameId = strings.SplitN(gameId, " ", 2)[1]
	gameId = strings.TrimSpace(gameId)
	gameIdInt, err := strconv.Atoi(gameId)
	if err != nil {
		panic(err)
	}

	revealSplit := strings.Split(result, ";")

	reveals := make([]Reveal, len(revealSplit))

	for _, reveal := range revealSplit {
		reveal = strings.TrimSpace(reveal)
		if reveal == "" {
			continue
		}
		revealSplit := strings.Split(reveal, ", ")

		for _, revealPart := range revealSplit {
			revealPartSplit := strings.Split(revealPart, " ")
			countStr := revealPartSplit[0]
			count, err := strconv.Atoi(countStr)
			if err != nil {
				panic(err)
			}
			color := revealPartSplit[1]
			reveals = append(reveals, Reveal{count: count, color: color})
		}
	}

	return Game{gameId: gameIdInt, reveals: reveals}
}

func validateGame(game Game) bool {
	reveals := game.reveals
	for _, reveal := range reveals {
		if reveal.color == "red" && reveal.count > 12 {
			return false
		}
		if reveal.color == "green" && reveal.count > 13 {
			return false
		}
		if reveal.color == "blue" && reveal.count > 14 {
			return false
		}
	}
	return true
}

func partOne() {
	readData := readFile("input.txt")

	var total int

	for _, line := range readData {
		if line == "" {
			continue
		}
		current := processLine(line)

		if validateGame(current) {
			total += current.gameId
		}

	}

	fmt.Println(total)

}

func partTwo() {
	readData := readFile("input.txt")

	var total int

	for _, line := range readData {
		if line == "" {
			continue
		}

		colorCountMap := make(map[string]int)

		colorCountMap["red"] = 0
		colorCountMap["green"] = 0
		colorCountMap["blue"] = 0

		var current Game
		current = processLine(line)

		for _, reveal := range current.reveals {
			color := reveal.color
			count := reveal.count

			if count > colorCountMap[color] {
				colorCountMap[color] = count
			}
		}

		total += colorCountMap["red"] * colorCountMap["green"] * colorCountMap["blue"]

	}

	fmt.Println(total)

}

func main() {
	partOne()
	partTwo()
}
