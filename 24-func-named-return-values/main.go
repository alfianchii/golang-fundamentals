package main

import "fmt"

func getMyFunnyFriends() (firstF string, secondF string, thirdF string) {
	firstF = "Moe"
	secondF = "Galih"
	thirdF = "Nugraha"

	return firstF, secondF, thirdF
}

func main() {
	var a, b, c string = getMyFunnyFriends()
	
	fmt.Println(a, b, c)
}