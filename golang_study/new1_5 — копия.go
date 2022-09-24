package main

import (
	"fmt"
)

func main() {
	var workArray [10]uint8
	var pair_number [6]uint8
	var c uint8

	for idx := range workArray {
		fmt.Scan(&workArray[idx])
	}
	for idx := range pair_number {
		fmt.Scan(&pair_number[idx])
	}
	for i := 0; i < len(pair_number)-1; i += 2 {
		c = workArray[pair_number[i]]
		workArray[pair_number[i]] = workArray[pair_number[i+1]]
		workArray[pair_number[i+1]] = c
	}
	for idx := range workArray {
		fmt.Print(workArray[idx], " ")
	}
}
