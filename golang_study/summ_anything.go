package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func adding(x string, y string) int64 {
	xRune := []rune(x)
	yRune := []rune(y)
	var first, second string
	for i := range xRune {
		if unicode.IsDigit(xRune[i]) {
			first += string(xRune[i])
		}
	}
	for i := range yRune {
		if unicode.IsDigit(yRune[i]) {
			second += string(yRune[i])
		}
	}
	x1, err := strconv.ParseInt(first, 10, 64)
	if err != nil {
		panic(err)
	}
	x2, err := strconv.ParseInt(second, 10, 64)
	if err != nil {
		panic(err)
	}
	return x1 + x2
}

func main() {
	fmt.Println(adding("2", "6"))
}
