package main

import (
	"fmt"
)

func main() {
	var n, max, sum int
	max = 0
	// считываем числа пока не будет введен 0
	for fmt.Scan(&n); n != 0; fmt.Scan(&n) {
		if max < n {
			max = n
			sum = 1
		} else if max == n {
			sum += 1
		}
	}
	fmt.Println(sum)
}
