package main

import (
	"fmt"
	"strconv"
)

func main() {
	fn := func(a uint) uint {
		//res := 0
		var resStr string
		resSlice := []rune(strconv.Itoa(int(a)))
		//fmt.Println(res)
		for i := range resSlice {
			res, _ := strconv.Atoi(string(resSlice[i]))
			if res%2 == 0 && res != 0 {
				resStr += strconv.Itoa(res)
			}
			//a = a / 10*/
			//fmt.Println(string(resSlice[i]))
		}
		resUint, _ := strconv.ParseUint(resStr, 10, 64)
		if len(resStr) == 0 {
			return 100
		} else {
			return uint(resUint)
		}
	}
	fmt.Println(fn(727183))
}
