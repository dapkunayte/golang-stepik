package main

import (
	"fmt"
	"math"
)

func main() {
	min := minimumFromFour()
	fmt.Println(min)
}

func minimumFromFour() int {
	var min int = math.MaxInt
	var a int
	for i := 0; i < 4; i++ {
		fmt.Scan(&a)
		if a < min {
			min = a
		}
	}
	return min
}
