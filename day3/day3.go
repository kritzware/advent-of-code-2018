package day3

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

const testInput = `#1 @ 1,3: 4x4
#2 @ 3,1: 4x4
#3 @ 5,5: 2x2`

// SolutionPart1 is the solution for /2018/day/3 part1
func SolutionPart1() {
	file, _ := ioutil.ReadFile("day3/input.txt")
	lines := strings.Split(strings.TrimSuffix(string(file), "\n"), "\n")
	// lines := strings.Split(strings.TrimSuffix(testInput, "\n"), "\n")

	var re = regexp.MustCompile(`(\d+),(\d+): (\d+)x(\d+)`)
	cords := make(map[string]int)

	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		if len(re.FindStringIndex(line)) == 0 {
			panic("no matching substrings found")
		}
		left, _ := strconv.Atoi(matches[1])
		top, _ := strconv.Atoi(matches[2])
		width, _ := strconv.Atoi(matches[3])
		height, _ := strconv.Atoi(matches[4])

		// fmt.Printf("Line: \"%s\", left=%d, top=%d, width=%d, height=%d\n", line, left, top, width, height)
		for i := left + 1; i < (left + 1 + width); i++ {
			for k := top + 1; k < (top + 1 + height); k++ {
				hash := fmt.Sprintf("%d,%d", i, k)
				cords[hash] = cords[hash] + 1
			}
		}
	}

	total := 0
	for _, val := range cords {
		if val > 1 {
			total++
		}
	}
	fmt.Println("Total overlap:", total)
}
