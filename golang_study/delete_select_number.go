package main

import (
	"fmt"
)

func main() {
	var n, c string
	fmt.Scan(&n, &c)
	for i := 0; i < len(n); i++ {
		if string(n[i]) != string(c[0]) {
			fmt.Print(string(n[i]))
		}
	}
}
