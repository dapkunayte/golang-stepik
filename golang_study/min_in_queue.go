package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string
	var max = "0"
	fmt.Scan(&a)
	b := strings.Split(a, "")
	for i := range b {
		if b[i] >= max {
			max = b[i]
		}
	}
	fmt.Println(max)
}
