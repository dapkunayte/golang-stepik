package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err // Если по какой-то причине мы получили ошибку, проигнорируем эту итерацию
	}
	if info.IsDir() {
		return nil // Проигнорируем директории
	}
	dataFromFile, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	if dataFromFile[0] != 0 && len(dataFromFile) == 290 {
		fmt.Println(path)
		fmt.Println(string(dataFromFile))
	}
	return nil
}

func main() {
	const root = "C:/Users/borov/Desktop/task_csv_1"
	if err := filepath.Walk(root, walkFunc); err != nil {
		fmt.Printf("Какая-то ошибка: %v\n", err)
	}
}
