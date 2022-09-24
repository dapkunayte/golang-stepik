package main

import "fmt"

func main() {
	channel := make(chan string)
	go task2(channel, "bebra")
	//task2(channel, "bebra")
	for i := 0; i < 5; i++ {
		s := <-channel
		fmt.Println(s)
	}
}

func task2(ch chan string, s string) {
	for i := 0; i < 5; i++ {
		ch <- s + " "
	}
}
