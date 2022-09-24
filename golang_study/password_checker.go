package main

import (
	"fmt"
	"unicode"
)

func main() {
	var n string
	fmt.Scan(&n)
	result := true
	rs := []rune(n)
	if len(rs) >= 5 {
		for i := range rs {
			if unicode.Is(unicode.Latin, rs[i]) || unicode.IsDigit(rs[i]) {

			} else {
				result = false
			}
		}
		if result == true {
			fmt.Println("Ok")
		} else {
			fmt.Println("Wrong password")
		}
	} else {
		fmt.Println("Wrong password")
	}

}
