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
	firstTime, err := time.Parse("2006-01-02 15:04:05", strings.TrimSpace(buf))
	if err != nil {
		panic(err)
	}
	//fmt.Println(firstTime.Unix())
	if firstTime.Hour() > 13 {
		future := firstTime.Add(time.Hour * 24)
		fmt.Println(future.Format("2006-01-02 15:04:05"))
	} else {
		fmt.Println(firstTime.Format("2006-01-02 15:04:05"))
	}
	//fmt.Println(firstTime.Format("2006-01-02 15:04:05"))
}
