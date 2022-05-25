package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main(){
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:]{
		go fetch(url, ch)
	}
	// Task 1.10: go run main.go https://mail.ru https://mail.ru https://mail.ru
	// Измените fetchall так, чтобы вывод осуществлялся
	// в файл и чтобы затем можно было его изучить.
	f, err := os.Create("task1.10")
	if(err != nil){
		fmt.Printf("Error when creating file: %v", err)
		os.Exit(1) 
	}
	fmt.Fprintln(f, "time\tbytes\turl")
	for range os.Args[1:]{
		fmt.Fprintln(f, <-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	f.Close()
}

func fetch(url string, ch chan<- string){
	start:= time.Now()
	resp, err := http.Get(url)
	if err != nil{
		ch <- fmt.Sprintf("%v", err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil{
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%v\t%7d\t%s", secs, nbytes, url)
}