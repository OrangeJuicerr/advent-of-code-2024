package main

import (
	"fmt"

	"github.com/cpalsa/advent-of-code-2024/util"
)

// var foo = 7

// type coordinate struct {
// 	x int
// 	y int
// }

func main() {
	// bar := 5
	// fmt.Println("Hello World")
	// fmt.Printf("The number is %d \n", add(foo, bar))

	// coord := coordinate{}
	// coord.x = 10
	// coord.y = 5
	// fmt.Printf("The number is %+v \n", coord)

	// bytes := 64
	// char := rune(bytes)
	// fmt.Println(char)

	input, _ := util.LoadInput("go.mod")
	str := string(input)

	fmt.Println(str)
}

// func add(a int, b int) int {
// 	return a + b
// }
