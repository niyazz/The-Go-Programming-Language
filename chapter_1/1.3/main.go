package main

import (
	"os"
	"bufio"
	"fmt"
)

// cd ./chapter_1/1.3

func main(){
// Task 1.4: go run main.go task1.4a task1.4b
// Измените программу dup2 так, чтобы она выводила имена всех 
// файлов, в которых найдены повторяющиеся строки.
	counts := make(map[string]int) 
	files := os.Args[1:] 
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
			for _, arg := range files { 
				f, err := os.Open(arg) 
				if err != nil {
					fmt.Fprintf(os.Stderr, "dup2: %v\n", err) 
					continue
				}
				countLines(f, counts) 
				f.Close()			
				for _, n := range counts { 
					if n > 1{
						fmt.Println(arg)
					}
				}
				counts = make(map[string]int) 
			}
		}
}

func countLines(f *os.File, counts map[string]int){
	input := bufio.NewScanner(f)
	for input.Scan(){
		counts[input.Text()]++
	}
}