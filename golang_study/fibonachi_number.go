package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	result := fibonacci(n)
	fmt.Println(result)
}

func fibonacci(n int) int {
	var a []int = []int{0, 1}
	for i := 0; i < n; i++ {
		a = append(a, a[i]+a[i+1])
	}
	return a[n]
}
