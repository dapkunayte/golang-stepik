package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err // Если по какой-то причине мы получили ошибку, проигнорируем эту итерацию
	}
	if info.IsDir() {
		return nil // Проигнорируем директории
	}
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	rd := bufio.NewReader(file)
	s, err := rd.ReadString('\n')
	result := strings.Split(s, ";")
	for i := range result {
		if result[i] == "0" {
			fmt.Println(i + 1)
		}
	}
	return nil
}

func main() {
	const root = "C:/Users/borov/Desktop/task_sep_1"
	if err := filepath.Walk(root, walkFunc); err != nil {
		fmt.Printf("Какая-то ошибка: %v\n", err)
	}
}
