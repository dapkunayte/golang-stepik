package main

import "fmt"

func main() {

	var a, hours, minutes int
	fmt.Scan(&a) // считаем переменную 'a' с консоли
	hours = a / 30
	minutes = (a - hours*30) * 2
	fmt.Println("It is", hours, "hours", minutes, "minutes.")
}
