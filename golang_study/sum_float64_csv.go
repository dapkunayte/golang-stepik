package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//добавить панику на ошибки можно
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	str = strings.Trim(str, "\n")
	str = strings.ReplaceAll(str, " ", "")
	str = strings.ReplaceAll(str, ",", ".")
	number := strings.Split(str, ";")
	res1, _ := strconv.ParseFloat(number[0], 64)
	res2, _ := strconv.ParseFloat(number[1], 64)
	res := res1 / res2
	fmt.Printf("%.4f", res)
}
