package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	rs := []rune(strings.ReplaceAll(text, " ", ""))
	rsNew := rs[:len(rs)-2]
	k := 0
	//if len(rsNew)%2 != 0 {
	for i := 0; i < (len(rsNew)-1)/2; i++ {
		if rsNew[i] == rs[len(rsNew)-1-i] {
			k++
		}
	}
	if k == (len(rsNew)-1)/2 {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
	//}
}
