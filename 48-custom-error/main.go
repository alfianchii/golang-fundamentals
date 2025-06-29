package main

import (
	"48-custom-error/user"
	"fmt"
)

type HttpError struct {
	Code int
	Message string
}

func (e *HttpError) Error() string {
	return e.Message
}

func main() {
	// var userID1 string = "taka"
	// var userID1 string = ""
	var userID1 string = "OK"

	var err error = user.Save(userID1, "OK")
	if err != nil {
		if validationErr, ok := err.(*user.ValidationError); ok { // Type assertions
			fmt.Println("Validation error:", validationErr.Message)
		} else if notFoundErr, ok := err.(*user.NotFoundError); ok { // Type assertions
			fmt.Println("Not found error:", notFoundErr.Message)
		} else {
			fmt.Println("Unknown error:", err.Error())
		}
	} else {
		fmt.Println("Success!")
	}

	// var userID2 string = "taka"
	// var userID2 string = "OK"
	var userID2 string = ""

	err = user.Save(userID2, "OK")
	if err != nil {
		switch finalError := err.(type) {
		case *user.ValidationError:
			fmt.Println("Validation error:", finalError.Error())
		case *user.NotFoundError:
			fmt.Println("Not found error:", finalError.Error())
		default:
			fmt.Println("Unknown error:", finalError.Error())
		}
	} else {
		fmt.Println("Success!")
	}
}