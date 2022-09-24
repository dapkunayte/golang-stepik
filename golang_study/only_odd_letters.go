package main

import (
	"fmt"
	"strings"
)

func main() {
	var n string
	fmt.Scan(&n)
	for i := range n {
		if strings.Count(n, string(n[i])) < 2 {
			fmt.Print(string(n[i]))
		}
	}
}
