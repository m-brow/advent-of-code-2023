package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type race struct {
	time       int
	distance   int
	waysToBeat int
}

func (r race) calculateWins() int {
	wins := 0
	for i := 0; i < r.time; i++ {
		if i*(r.time-i) > r.distance {
			wins += 1
		}
	}
	r.waysToBeat = wins
	return wins
}

func main() {
	fmt.Println("Day 6")

	dat, err := os.ReadFile("./input_p2.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\r\n")
	times := strings.Fields(lines[0][9:])
	distances := strings.Fields(lines[1][9:])

	var races []race

	fmt.Println(times)
	fmt.Println(distances)

	for i := 0; i < len(times); i++ {
		if times[i] == "" {
			continue
		}
		t, _ := strconv.Atoi(times[i])
		d, _ := strconv.Atoi(distances[i])
		r := race{
			time:     t,
			distance: d,
		}

		races = append(races, r)
	}

	fmt.Println(races)
	result := 1
	for _, r := range races {
		result *= r.calculateWins()
	}
	fmt.Println(result)

}
