package two

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func countFrequencies(str string) (bool, bool) {
	frequency := make([]int, 256)
	countTwos := 0
	countThrees := 0
	for _, r := range str {
		frequency[r] = frequency[r] + 1
		if frequency[r] == 2 {
			countTwos++
		}
		if frequency[r] == 3 {

			countThrees++
			countTwos--
		}
		if frequency[r] == 4 {
			countThrees--
		}
	}
	return countTwos > 0, countThrees > 0
}

// SolveTwoA solves part 2a of the challenge.
func SolveTwoA() {
	content, err := ioutil.ReadFile("2/input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	lines := strings.Split(string(content), "\n")
	lines = lines[:len(lines)-1]
	totalTwoCount := 0
	totalThreeCount := 0
	for _, line := range lines {
		twoCountExists, threeCountExists := countFrequencies(line)
		if twoCountExists {
			totalTwoCount++
		}
		if threeCountExists {
			totalThreeCount++
		}
	}
	fmt.Println(strconv.Itoa(totalTwoCount * totalThreeCount))
}
