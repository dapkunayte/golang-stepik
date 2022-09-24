package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type T struct {
	ID       int    `json:"ID"`
	Number   string `json:"Number"`
	Year     int    `json:"Year"`
	Students []struct {
		LastName   string `json:"LastName,omitempty"`
		FirstName  string `json:"FirstName,omitempty"`
		MiddleName string `json:"MiddleName,omitempty"`
		Birthday   string `json:"Birthday,omitempty"`
		Address    string `json:"Address,omitempty"`
		Phone      string `json:"Phone,omitempty"`
		Rating     []int  `json:"Rating,omitempty"`
	} `json:"Students"`
}
type Answer struct {
	Average float64
}

func main() {
	file, _ := os.Open("text.json")
	data, _ := ioutil.ReadAll(file)
	var s T
	if err := json.Unmarshal(data, &s); err != nil {
		fmt.Println(err)
		return
	}
	var sum int
	for i := range s.Students {
		//fmt.Println(s.Students[i].Rating)
		sum += len(s.Students[i].Rating)
	}
	res := float64(sum) / float64(len(s.Students))
	result := Answer{Average: res}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", data)
	/*fmt.Printf("%s", data)
	rating := strings.Split(string(data), ",")
	fmt.Println(rating)*/
}
