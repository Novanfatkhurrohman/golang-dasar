package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You Bane", name)
	} else {
		fmt.Println("Welcome")
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "anjing"
	}
	registerUser("anjing", blacklist)
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}
