package main

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(input, "\n")
	var left []string
	var right []string
	for _, line := range lines {
		pair := strings.Split(line, "   ")
		left = append(left, pair[0])
		right = append(right, pair[1])
	}

	// S* printf returns a formatted string
	// F* printf sends a formatted string to a writer interface
	// Printf sends a formatted string to std out

	// fmt.Println(fmt.Sprintf("My data is %#v", left))

	slices.Sort(left)
	slices.Sort(right)

	sum := 0.0
	for i := range len(left) {
		l, err := strconv.Atoi(left[i])
		if err != nil {
			fmt.Println("Conversion error:", err)
			return
		}
		r, err := strconv.Atoi(right[i])
		if err != nil {
			fmt.Println("Conversion error:", err)
			return
		}
		sum += math.Abs(float64(l) - float64(r))
	}

	fmt.Println(int(sum))
}
