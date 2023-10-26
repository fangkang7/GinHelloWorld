package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		go test()
	}

	for i := 0; i < 10; i++ {
		fmt.Println("hello world goland", i)
	}
}

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello world", i)
	}
}
