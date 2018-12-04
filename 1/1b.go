package one

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// SolveOneB solves part 1b of the challenge.
func SolveOneB() {
	content, err := ioutil.ReadFile("1/input.txt")
	if err != nil {
		fmt.Printf("error reading file")
	}
	lines := strings.Split(string(content), "\n")
	lines = lines[:len(lines)-1]
	counter := 0
	resultFound := false
	pastValues := make(map[int]bool)
	pastValues[counter] = true
	for resultFound == false {
		for _, num := range lines {
			i, err := strconv.Atoi(num)
			if err != nil {
				fmt.Println("error parsing int")
			}
			counter += i
			if pastValues[counter] {
				resultFound = true
				break
			}
			pastValues[counter] = true
		}
	}
	fmt.Println(counter)
}
