package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int, 1)
	go func(ch chan int, i int) int {
		ch <- i + 1
		fmt.Println(<-ch)
		return <-ch
	}(channel, 5)
	<-time.After(time.Second * 5)
}

func task(ch chan int, i int) int {
	ch <- i + 1
	return <-ch
}
