package main

import (
	"fmt"
	"os"
	"strconv"
)

var pc [256]byte
func init(){
	for i:= range pc{
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int{
	return int(pc[byte(x>>(0*8))] + 
	pc[byte(x>>(1*8))] + 
	pc[byte(x>>(2*8))] + 
	pc[byte(x>>(3*8))] + 
	pc[byte(x>>(4*8))] + 
	pc[byte(x>>(5*8))] + 
	pc[byte(x>>(6*8))] +
	pc[byte(x>>(7*8))])
}

// Task 2.3: go run . 12
// Перепишите функцию PopCount так, чтобы она использовала 
// цикл вместо единого выражения. Сравните производительность двух версий.
func PopCount23(x uint64) int{
	var res byte = 0

	for i:=0; i < 8; i++{
		res += pc[byte(x>>(i*8))]
	}

	return int(res)
}

// Task 2.4: go run . 12
// Напишите версию PopCount, которая подсчитывает биты с помощью 
// сдвига аргумента по всем 64 позициям, проверяя при каждом сдвиге крайний 
// справа бит.
func PopCount24(x uint64) int{
	res := 0
	for i := 0; i < 64; i++{
		res += int((x) & 1)
		x = x >> 1
	} 
	return res
}

// Task 2.5:  go run . 12
// Выражение х&(х-1) сбрасывает крайний справа ненулевой 
// бит х. Напишите версию PopCount, которая подсчитывает биты с использованием 
// этого факта, и оцените ее производительность.
func PopCount25(x uint64) int{
	res := 0
	for x > 0{
		res++
		x = x&(x-1)
	}
	return res
}

func main(){
	for _, n := range os.Args[1:]{
		v, _ := strconv.ParseUint(n, 10, 64)
		fmt.Println(PopCount23(v))
		fmt.Println(PopCount24(v))
		fmt.Println(PopCount25(v))
	}
}