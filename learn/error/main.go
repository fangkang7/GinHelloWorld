package main

import (
	"errors"
	"fmt"
)

func main()  {
	err := testError()
	fmt.Println(err)
}

func testError() error {
	return errors.New("wohsi kak ")
}
