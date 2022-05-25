package main

import (
	"fmt"
	"os"
	"strings"
)

// cd ./chapter_1/1.2

func main(){
// Task 1.1: go run main.go 1 2 3
// Измените программу echo так, чтобы она выводила также os.Args[0],
// имя выполняемой команды
	fmt.Println(strings.Join(os.Args[0:], " "))

// Task 1.2: go run main.go 1 2 3
// Измените программу echo так, чтобы она выводила индекс и 
// значение каждого аргумента по одному аргументу в строке.
	for i, v := range os.Args{
		fmt.Printf("Args[%d] = %v\n", i, v)
	}
}