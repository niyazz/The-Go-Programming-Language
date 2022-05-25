package main

import (
	"net/http"
	"io"
	"fmt"
	"os"
	"strings"
)
const schema = "http://"

func main(){
	for _, url := range os.Args[1:]{

// Task 1.8: go run main.go github.com
// Измените программу fetch так, чтобы к каждому аргументу 
// URL автоматически добавлялся префикс http:// в случае отсутствия в нем такового.
		if !strings.HasPrefix(url, schema){
			url = schema + url
		}

		resp, err := http.Get(url)
		if err != nil{
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		
// Task 1.7: go run main.go http://github.com
// Воспользуйтесь ею вместо ioutil.ReadAll для копирования тела 
// ответа в поток os.Stdout без необходимости выделения достаточно большого для 
// хранения всего ответа буфера.
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil{
			fmt.Fprintf(os.Stderr, "fetch: cannot read %s, because %v\n", url, err)
		}

// Task 1.8: go run main.go http://github.com, github.com
// Измените программу fetch так, чтобы она выводила код состояния HTTP, 
// содержащийся в resp.Status.
		fmt.Printf("%d", resp.StatusCode)
	}
}