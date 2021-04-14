package main

import "fmt"

func main() {
	person := map[string]string{
		"name":   "novan",
		"addres": "Tegal",
	}
	person["title"] = "Golang"
	delete(person, "title")

	fmt.Println(person)
	fmt.Println(len(person["name"]))
	fmt.Println(person["addres"])
	fmt.Println(person["title"])
}
