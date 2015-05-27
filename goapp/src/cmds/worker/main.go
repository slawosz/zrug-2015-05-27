package main

import (
	"fmt"
	"github.com/benmanns/goworker"
)

func workerFunc(queue string, args ...interface{}) error {
	fmt.Printf("Hello, go. %v \n", args)
	return nil
}

func init() {
	goworker.Register("SimpleWorker", workerFunc)
}

func main() {
	if err := goworker.Work(); err != nil {
		fmt.Println("Error:", err)
	}
}
