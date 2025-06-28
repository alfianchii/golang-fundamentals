package main

import "fmt"

func logging() {
	fmt.Println("App is closed!")
}

func runApp() {
	defer logging()
	fmt.Println("Run the app...")
}

func main() {
	runApp()
}