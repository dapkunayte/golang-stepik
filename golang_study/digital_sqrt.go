package main

import "fmt"

func main() {
	var a, sum int
	fmt.Scan(&a)
	for a != 0 {
		sum += a % 10
		a /= 10
	}
	for sum != 0 {
		a += sum % 10
		sum /= 10
	}
	fmt.Println(a)
}
