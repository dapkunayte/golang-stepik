package main

import (
	"fmt"
)

func main() {
	voteResult := vote(0, 0, 1)
	fmt.Println(voteResult)
}

func vote(x int, y int, z int) int {
	var a []int
	a = append(a, x, y, z)
	var counterZeros, counterOnes, result int
	for i := 0; i < len(a); i++ {
		if a[i] == 0 {
			counterZeros++
		} else {
			counterOnes++
		}
	}
	if counterZeros > counterOnes {
		result = 0
	} else {
		result = 1
	}
	return result
}
