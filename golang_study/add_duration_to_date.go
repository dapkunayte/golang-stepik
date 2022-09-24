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
	res := strings.ReplaceAll(buf, " сек.", "s")
	res = strings.ReplaceAll(res, " мин. ", "m")
	//fmt.Println(res[:len(res)-2])
	dur, err := time.ParseDuration(res[:len(res)-2])
	if err != nil {
		panic(err)
	}
	const now = 1589570165
	nowTime := time.Unix(now, 0).UTC()
	nowTime = nowTime.Add(dur)
	fmt.Println(nowTime.Format(time.UnixDate))
}
