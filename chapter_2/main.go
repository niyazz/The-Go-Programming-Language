package main

import (
	"fmt"
	"os"
	"strconv"
	"github.com/niyazz/The-Go-Programming-Language/chapter_2/weight"
) 


// Task 2.2: cd ./chapter_2 ; go run .  12
// Напишите программу общего назначения для преобразования 
// единиц, аналогичную c f , которая считывает числа из аргументов командной строки 
// (или из стандартного ввода, если аргументы командной строки отсутствуют)
// и преобразует каждое число в другие единицы, как температуру — в градусы Цельсия и Фаренгейта, длину — в футы и метры, 
// вес — в фунты и килограммы и т.д.
func main(){
	for _, arg := range os.Args[1:]{
		t , err:= strconv.ParseFloat(arg, 64)
		if err != nil{
			os.Exit(1)
		}
		fmt.Println(weightconv.KgToP(weightconv.Kilo(t)))
		fmt.Println(weightconv.PToKg(weightconv.Pund(t)))
	}
}