package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	rs := []rune(text)
	if unicode.IsUpper(rs[0]) && string(rs[len(rs)-2]) == "." {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}
