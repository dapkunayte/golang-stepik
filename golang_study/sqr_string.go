package main

import (
	"fmt"
)

func main() {
	var a string
	fmt.Scan(&a)
	b := []rune(a)
	for i := range b {
		tmp := b[i] - 48
		fmt.Print(tmp * tmp)
	}
}
