package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func findNumber_p1(calibrated string) int {
	var firstNum, lastNum rune
	var outputStr string
	var output int

	if calibrated == "" {
		return 0
	}

	for _, val := range calibrated {
		if unicode.IsDigit(val) {
			if firstNum == 0 {
				firstNum = val
			} else {
				lastNum = val
			}
		}
	}

	if firstNum != 0 {
		outputStr += string(firstNum)
		if lastNum != 0 {
			outputStr += string(lastNum)
		} else {
			outputStr += string(firstNum)
		}
	}

	output, err := strconv.Atoi(outputStr)
	if err != nil {
		fmt.Println(err)
	}

	return output
}

func part1() {
	dat, err := os.ReadFile("./input.txt")

	if err != nil {
		panic(err)
	}
	var total int = 0

	lines := strings.Split(string(dat), "\r\n")
	for _, line := range lines {
		total += findNumber_p1(line)
	}

	fmt.Printf("[Part 1] Total: %d\n", total)
}

func checkNumber(numberStr string) int {

	if len(numberStr) >= 3 {
		if numberStr[:3] == "one" {
			return 1
		}
		if numberStr[:3] == "two" {
			return 2
		}
		if numberStr[:3] == "six" {
			return 6
		}
	}
	if len(numberStr) >= 4 {
		if numberStr[:4] == "four" {
			return 4
		}
		if numberStr[:4] == "five" {
			return 5
		}
		if numberStr[:4] == "nine" {
			return 9
		}
	}
	if len(numberStr) >= 5 {
		if numberStr[:5] == "three" {
			return 3
		}
		if numberStr[:5] == "seven" {
			return 7
		}
		if numberStr[:5] == "eight" {
			return 8
		}
	}
	return 0
}

func findNumber_p2(calibrated string) int {
	var firstNum, lastNum rune
	var outputStr string
	var output int

	if calibrated == "" {
		return 0
	}

	for idx, val := range calibrated {
		if unicode.IsDigit(val) {
			if firstNum == 0 {
				firstNum = val
			} else {
				lastNum = val
			}
		} else if num := checkNumber(calibrated[idx:]); num > 0 {
			if firstNum == 0 {
				firstNum = rune(strconv.Itoa(num)[0])
			} else {
				lastNum = rune(strconv.Itoa(num)[0])
			}
		}
	}

	if firstNum != 0 {
		outputStr += string(firstNum)
		if lastNum != 0 {
			outputStr += string(lastNum)
		} else {
			outputStr += string(firstNum)
		}
	}

	output, err := strconv.Atoi(outputStr)
	if err != nil {
		fmt.Println(err)
	}

	return output
}

func part2() {
	dat, err := os.ReadFile("./input.txt")

	if err != nil {
		panic(err)
	}
	var total int = 0

	lines := strings.Split(string(dat), "\r\n")
	for _, line := range lines {
		total += findNumber_p2(line)
	}

	fmt.Printf("[Part 2] Total: %d\n", total)
}
func main() {
	fmt.Println("Day 1!")
	part1()
	part2()
}
