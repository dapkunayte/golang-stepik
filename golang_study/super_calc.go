package main

import (
	fmt "fmt"
	"os"
)

func readTask() (value1, value2, operation interface{}) {

	return 5.0, 5, "/" //тут играемся с параметрами

}

func main() {
	value1, value2, operation := readTask()
	if v, ok := value1.(float64); ok {
		if v1, ok1 := value2.(float64); ok1 {
			if v2, ok2 := operation.(string); ok2 {
				operation_func(v2, v, v1)
			} else {
				fmt.Println("неизвестная операция")
			}
		} else {
			fmt.Printf("value=%v: %T", value2, value2)
		}
	} else {
		fmt.Printf("value=%v: %T", value1, value1)
	}
	os.Exit(0)
}

func operation_func(s string, a, b float64) {
	switch s {
	case "+":
		fmt.Printf("%.4f", a+b)
	case "-":
		fmt.Printf("%.4f", a-b)
	case "/":
		fmt.Printf("%.4f", a/b)
	case "*":
		fmt.Printf("%.4f", a*b)
	default:
		fmt.Println("неизвестная операция")
	}
}
