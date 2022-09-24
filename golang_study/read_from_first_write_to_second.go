package main

import (
	"fmt"
	"time"
)

func main() {
	inputStream := make(chan string)
	outputStream := make(chan string)
	go removeDuplicates(inputStream, outputStream)
	go func() {
		for _, r := range "112334456" {
			inputStream <- string(r)
		}
	}()
	<-time.After(time.Second * 5)
}

func removeDuplicates(ch1 chan string, ch2 chan string) {
	var tmp string
	go func() {
		for {
			v, ok := <-ch1
			if !ok {
				close(ch2)
				break
			} else if tmp != v {
				ch2 <- v
				tmp = v
			}
			//close(ch2)
		}
	}()
	for x := range ch2 {
		fmt.Println(x)
	}
}
