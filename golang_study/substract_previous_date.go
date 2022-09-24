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
	timeArr := strings.Split(strings.TrimSpace(buf), ",")
	firstTime, err := time.Parse("02.01.2006 15:04:05", timeArr[0])
	if err != nil {
		panic(err)
	}
	secondTime, err := time.Parse("02.01.2006 15:04:05", timeArr[1])
	if err != nil {
		panic(err)
	}
	if firstTime.Before(secondTime) {
		fmt.Println(secondTime.Sub(firstTime))
	} else {
		fmt.Println(firstTime.Sub(secondTime))
	}
	//fmt.Println(firstTime.Unix())
	//fmt.Println(firstTime.Format("2006-01-02 15:04:05"))
}
