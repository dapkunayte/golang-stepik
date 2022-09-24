package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type battery struct {
	string
}

var t interface {
	fmt.Stringer
}

func (b battery) String() string {
	//var oneCounter,zeroCounter int
	oneCounter := strings.Count(b.string, "1")
	zeroCounter := strings.Count(b.string, "0")
	var result string
	for i := 0; i <= zeroCounter; i++ {
		if i == 0 {
			result += "["
		} else {
			result += " "
		}
	}
	for i := 0; i <= oneCounter; i++ {
		if i == oneCounter {
			result += "]"
		} else {
			result += "X"
		}
	}
	return result
}

func main() {
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	str = strings.Trim(str, "\n")
	batteryForTest := battery{str}
	fmt.Println(batteryForTest)
}
