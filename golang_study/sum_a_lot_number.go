package main

import (
	"fmt"
)

func main() {
	a, b := sumInt(1, 0)
	fmt.Println(a, b)
}

func sumInt(a ...int) (int, int) {
	var k, sum int
	for i := 0; i < len(a); i++ {
		k++
		sum += a[i]
	}
	return k, sum
}
