package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func part1(lines []string) int {
	rowsMax := len(lines)
	colsMax := len(lines[0])
	output := 0

	for rows, line := range lines {
		for cols, val := range line {
			if !unicode.IsDigit(val) && val != '.' {
				// This is a special symbol location
				for b := rows - 1; b <= rows+1; b++ {
					for a := cols - 1; a <= cols+1; a++ {
						// Check out of bounds
						if b > rowsMax || a > colsMax || a < 0 || b < 0 {
							continue
						}
						if a == cols && b == rows {
							continue
						}
						if unicode.IsDigit(rune(lines[b][a])) {

							// Found a digit - need to go left the find the start of the digit
							leftMost_x := a
							for a2 := a; a2 >= 0; a2-- {
								if unicode.IsDigit(rune(lines[b][a2])) {
									leftMost_x = a2
								} else {
									break
								}
							}
							// Need to find the end of the value
							rightMost_x := a
							for a2 := a; a2 < colsMax; a2++ {
								if unicode.IsDigit(rune(lines[b][a2])) {
									rightMost_x = a2
								} else {
									break
								}
							}
							// Add the value to the sum
							partNum := lines[b][leftMost_x : rightMost_x+1]
							val, err := strconv.Atoi(partNum)
							if err != nil {
								panic(err)
							}
							output += val

							// Erase the value so we don't count it multiple times.
							for a2 := leftMost_x; a2 <= rightMost_x; a2++ {
								lines[b] = replaceAtIndex(lines[b], '.', a2)
							}
						}
					}
				}
			}
		}
	}

	return output
}

func findPartNumber(cols int, rows int, lines []string, symbol rune) int {
	rowsMax := len(lines)
	colsMax := len(lines[0])
	output := 0

	var partNums []int

	for b := rows - 1; b <= rows+1; b++ {
		for a := cols - 1; a <= cols+1; a++ {
			// Check out of bounds
			if b > rowsMax || a > colsMax || a < 0 || b < 0 {
				continue
			}
			if a == cols && b == rows {
				continue
			}
			if unicode.IsDigit(rune(lines[b][a])) {
				// Found a digit - need to go left the find the start of the digit
				leftMost_x := a
				for a2 := a; a2 >= 0; a2-- {
					if unicode.IsDigit(rune(lines[b][a2])) {
						leftMost_x = a2
					} else {
						break
					}
				}
				// Need to find the end of the value
				rightMost_x := a
				for a2 := a; a2 < colsMax; a2++ {
					if unicode.IsDigit(rune(lines[b][a2])) {
						rightMost_x = a2
					} else {
						break
					}
				}
				// Add the value to the sum
				partNum := lines[b][leftMost_x : rightMost_x+1]
				val, err := strconv.Atoi(partNum)
				if err != nil {
					panic(err)
				}
				partNums = append(partNums, val)

				// Erase the value so we don't count it multiple times.
				for a2 := leftMost_x; a2 <= rightMost_x; a2++ {
					lines[b] = replaceAtIndex(lines[b], '.', a2)
				}
			}
		}
	}

	// Multiply the two values if only 2 parts AND '*' symbol.
	if len(partNums) == 2 && symbol == '*' {
		output += partNums[0] * partNums[1]
	} else {
		for i := range partNums {
			output += i
		}
	}

	return output
}

func part2(lines []string) int {
	output := 0
	for rows, line := range lines {
		for cols, val := range line {
			if !unicode.IsDigit(val) && val != '.' {
				output += findPartNumber(cols, rows, lines, val)
			}
		}
	}

	return output
}

func main() {
	fmt.Println("Day 3!")
	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\r\n")
	result := part1(lines)
	fmt.Printf("Part 1: %d\n", result)

	lines = strings.Split(string(dat), "\r\n")
	result = part2(lines)
	fmt.Printf("Part 2: %d\n", result)

}
