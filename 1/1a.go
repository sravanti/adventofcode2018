package one

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// SolveOneA solves part 1a of the challenge.
func SolveOneA() {
	file, _ := os.Open("1/input.txt")
	fscanner := bufio.NewScanner(file)
	counter := 0
	for fscanner.Scan() {
		i, err := strconv.Atoi(fscanner.Text())
		if err != nil {
			fmt.Printf("Error converting string to num")
		}
		counter += i
	}
	fmt.Println(strconv.Itoa(counter))
}
