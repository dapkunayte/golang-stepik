package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	buf, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	firstTime, err := time.Parse(time.RFC3339, strings.TrimSpace(buf))
	if err != nil {
		panic(err)
	}
	//fmt.Println(firstTime.Unix())
	fmt.Println(firstTime.Format(time.UnixDate))
}
