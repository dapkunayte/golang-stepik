package main

import (
	"fmt"
)

func main() {
	var n, min, sum, a int
	min = 100000
	fmt.Scan(&n)
	// считываем числа пока не будет введен 0
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		if min > a {
			min = a
			sum = 1
		} else if min == a {
			sum += 1
		}
	}
	fmt.Println(sum)
}
