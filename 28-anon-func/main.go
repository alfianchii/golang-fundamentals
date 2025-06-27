package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, isBlacklist Blacklist) {
	if isBlacklist(name) {
		fmt.Println("You are blocked, " + name + ".")
	} else {
		fmt.Println("Welcome home, " + name + "<3.")
	}
}

func main() {
	var blacklistedUser string = "Galih"
	
	// This is an anon func
	var isBlacklist Blacklist = func(name string) bool {
		return name == blacklistedUser
	}

	registerUser("Taka", isBlacklist)
	// This is also an anon func
	registerUser("Galih", func(name string) bool {
		return name == blacklistedUser
	})
}