package main

import "fmt"

func fullName() (string, string) {
	return "novan", "golang"
}

func main() {
	firstName, lastName := fullName()
	fmt.Println(firstName, lastName)
}
