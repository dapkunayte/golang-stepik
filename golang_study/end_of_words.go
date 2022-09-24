package main

import "fmt"

func main() {
	var i int
	fmt.Scan(&i)
	switch {
	case i%10 == 1 && !(11 <= i && i <= 14):
		fmt.Println(i, "korova")
	case (i%10 == 3 || i%10 == 2 || i%10 == 4) && !(11 <= i && i <= 14):
		fmt.Println(i, "korovy")
	case i%10 == 0 || i%10 >= 5 || (11 <= i && i <= 14):
		fmt.Println(i, "korov")
	}
}
