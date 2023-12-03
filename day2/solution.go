package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type hand struct {
	red   int
	green int
	blue  int
}

type game struct {
	id    int
	hands []hand
}

func parseInput(gameInfo string) []game {
	var output []game

	lines := strings.Split(gameInfo, "\r\n")
	for _, line := range lines {
		partition := strings.Split(line, ":")

		gameIdStr := partition[0][5:]
		gameId, err := strconv.Atoi(gameIdStr)
		if err != nil {
			panic(err)
		}
		gameHands := strings.Split(partition[1], ";")

		g := game{
			id: gameId,
		}

		for _, vals := range gameHands {
			h := hand{}
			reBlue := regexp.MustCompile(`(\d+) blue`)
			blueVals := reBlue.FindStringSubmatch(vals)
			if len(blueVals) > 1 {
				h.blue, _ = strconv.Atoi(blueVals[1])
			}

			reGreen := regexp.MustCompile(`(\d+) green`)
			greenVals := reGreen.FindStringSubmatch(vals)
			if len(greenVals) > 1 {
				h.green, _ = strconv.Atoi(greenVals[1])
			}

			reRed := regexp.MustCompile(`(\d+) red`)
			redVals := reRed.FindStringSubmatch(vals)
			if len(redVals) > 1 {
				h.red, _ = strconv.Atoi(redVals[1])
			}
			g.hands = append(g.hands, h)
		}

		output = append(output, g)
	}

	return output
}

func gamesPossible(redMax, greenMax, blueMax int, games []game) int {

	output := 0

	for _, g := range games {
		possible := true

		for _, h := range g.hands {
			if h.red > redMax {
				possible = false
			}
			if h.green > greenMax {
				possible = false
			}
			if h.blue > blueMax {
				possible = false
			}
		}

		if possible {
			output += g.id
		}
	}
	return output
}

func gameMinimums(games []game) int {
	output := 0
	for _, g := range games {
		var redMin, greenMin, blueMin int = 0, 0, 0
		for _, h := range g.hands {
			if redMin < h.red {
				redMin = h.red
			}
			if greenMin < h.green {
				greenMin = h.green
			}
			if blueMin < h.blue {
				blueMin = h.blue
			}
		}

		output += redMin * greenMin * blueMin
	}
	return output
}

func main() {
	fmt.Println("Day 2!")
	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	games := parseInput(string(dat))

	part1 := gamesPossible(12, 13, 14, games)
	fmt.Printf("Part 1: %d\n", part1)

	part2 := gameMinimums(games)
	fmt.Printf("Part 2: %d\n", part2)
}
