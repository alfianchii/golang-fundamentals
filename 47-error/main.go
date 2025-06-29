package main

import (
	"errors"
	"fmt"
)

func Divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("couldn't divide w/ zero")
	} else {
		return a / b, nil
	}
}

func main() {
	result, err := Divide(4, 0)
	if err == nil {
		fmt.Println(result)
		} else {
		fmt.Println(err.Error())
	}
}