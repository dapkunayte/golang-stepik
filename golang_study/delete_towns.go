package main

import "fmt"

func work(x int) int {
	var result int
	return result
}

func main() {
	groupCity := map[int][]string{
		10:   []string{"Деревня", "Село"},        // города с населением 10-99 тыс. человек
		100:  []string{"Город", "Большой город"}, // города с населением 100-999 тыс. человек
		1000: []string{"Миллионик"},              // города с населением 1000 тыс. человек и более
	}
	cityPopulation := map[string]int{
		"Село":      100,
		"Миллионик": 500,
		"Город":     100,
	}
	for key, _ := range cityPopulation {
		for key1, value1 := range groupCity {
			if key1 != 100 {
				for i := range value1 {
					if key == value1[i] {
						delete(cityPopulation, key)
						i++
					}
				}
			}
		}
	}
	fmt.Println(cityPopulation)
}
