package two

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func wordToRunes(str string) []int {
	letters := make([]int, len(str))
	for i, r := range str {
		letters[i] = int(r - '0')
	}
	return letters
}

func calculateOffByOne(freq1 []int, freq2 []int) bool {
	differentLetters := 0
	for i := 0; i < len(freq1); i++ {
		if freq1[i] != freq2[i] {
			differentLetters++
		}
		if differentLetters > 1 {
			return false
		}
	}
	if differentLetters == 0 {
		return false
	}
	return true
}

// SolveTwoB solves part 2b of the challenge.
func SolveTwoB() {
	content, err := ioutil.ReadFile("2/input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	lines := strings.Split(string(content), "\n")
	lines = lines[:len(lines)-1]
	for _, line1 := range lines {
		for _, line2 := range lines[1 : len(lines)-2] {
			if calculateOffByOne(wordToRunes(line1), wordToRunes(line2)) {
				fmt.Println(line1)
				fmt.Println(line2)
				break
			}
		}
	}
}
