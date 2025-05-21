package main

import "fmt"

func main() {
	var firstName string = "Alfian"
	var lastName string = "Taka"

	var isFirstnameEqualsLastname bool = firstName == lastName
	fmt.Println(isFirstnameEqualsLastname)

	var isFirstnameNotEqualsLastName bool = firstName != lastName
	fmt.Println(isFirstnameNotEqualsLastName)
	fmt.Println("----------")

	var str1 string = "A"
	var str2 string = "B"

	var isStr1MoreThanStr2 bool = str1 > str2
	fmt.Println(isStr1MoreThanStr2)
	fmt.Println(str1[0], str2[0])
	fmt.Println("----------")

	var a uint8 = 10
	var b uint8 = 5

	fmt.Println("The value of 'a':", a)
	fmt.Println("The value of 'b':", b)
	
	var isAMoreThanB bool = a > b
	fmt.Println("Is 'a' more than 'b'?", isAMoreThanB)
	var isALessThanB bool = a < b
	fmt.Println("Is 'a' less than 'b'?", isALessThanB)
	var isAMoreThanEqualsB bool = a >= b
	fmt.Println("Is 'a' more than equals 'b'?", isAMoreThanEqualsB)
	var isALessThanEqualsB bool = a <= b
	fmt.Println("Is 'a' less than equals 'b'?", isALessThanEqualsB)
	var isAEqualsB bool = a == b
	fmt.Println("Is 'a' equals 'b'?", isAEqualsB)
	var isANotEqualsB bool = a != b
	fmt.Println("Is 'a' not equals 'b'?", isANotEqualsB)
}