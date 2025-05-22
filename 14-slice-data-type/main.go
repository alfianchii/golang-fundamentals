package main

import "fmt"

func main() {
	// The slice method
	names := [...]string{"Taka", "Ogi", "Kanji", "Nugraha", "Moe"}
	firstSlice := names[1:3]
	secondSlice := names[3:]
	thirdSlice := names[:3]
	fourthSlice := names[:]

	fmt.Println("The slice:", names)
	fmt.Println("----------------")
	fmt.Println("First slice:", firstSlice)
	fmt.Println("Second slice:", secondSlice)
	fmt.Println("Third slice:", thirdSlice)
	fmt.Println("Fourth slice:", fourthSlice)
	fmt.Println("========================================================")
}