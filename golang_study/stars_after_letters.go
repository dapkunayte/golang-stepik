package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string
	fmt.Scan(&a)
	a = strings.ReplaceAll(a, "", "*")
	fmt.Println(string(a[1 : len(a)-1]))
}
