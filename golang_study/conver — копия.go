package main

import (
	"fmt"
	"strconv"
)

func main() {
	fn := func(a uint) uint {
		res := 0
		var resStr string
		//resSlice := strconv.Itoa(int(a))
		//fmt.Println(res)
		for a != 0 {
			res = int(a % 10)
			if res%2 == 0 && res != 0 {
				resStr += strconv.Itoa(res)
			}
			a = a / 10
		}
		resUint, _ := strconv.ParseUint(resStr, 10, 64)
		if len(resStr) == 0 {
			return 100
		} else {
			return uint(resUint)
		}
	}
	fmt.Println(fn(717173))
}
