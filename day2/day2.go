package day2

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

const testInput = `abcdef
bababc
abbcde
abcccd
aabcdd
abcdee
ababab`

const testInput2 = `abcde
fghij
klmno
pqrst
fguij
axcye
wvxyz`

// SolutionPart1 is the solution for /2018/day/2 part1
func SolutionPart1() {
	file, _ := ioutil.ReadFile("day2/input.txt")
	lines := strings.Split(strings.TrimSuffix(string(file), "\n"), "\n")
	finalCounts := make(map[int]int)
	checkSum := 1

	for _, line := range lines {
		uniqCounts := make(map[int]bool)
		counts := make(map[string]int)

		for _, c := range line {
			letter := string(c)
			count := strings.Count(line, letter)
			if count > 1 {
				counts[letter] = count
			}
		}

		for _, val := range counts {
			if _, ok := uniqCounts[val]; !ok {
				uniqCounts[val] = true
			}
		}

		for c := range uniqCounts {
			finalCounts[c]++
		}
	}

	for _, val := range finalCounts {
		checkSum *= val
	}
	fmt.Printf("Counts: %+v, Checksum: %d\n", finalCounts, checkSum)
}

// SolutionPart2 is the solution for /2018/day/2 part2
func SolutionPart2() {
	file, _ := ioutil.ReadFile("day2/input.txt")
	lines := strings.Split(strings.TrimSuffix(string(file), "\n"), "\n")
	var word1, word2 string

	for _, line1 := range lines {
		for _, line2 := range lines {
			if line1 != line2 {
				diff := 0
				for k := 0; k < len(line1); k++ {
					if line1[k] != line2[k] {
						diff++
					}
				}
				if diff == 1 {
					word1, word2 = line1, line2
				}
			}
		}
	}

	var finalWord bytes.Buffer
	for k := 0; k < len(word1); k++ {
		if word1[k] == word2[k] {
			finalWord.WriteString(string(word1[k]))
		}
	}
	fmt.Println("Final word:", finalWord.String())
}
