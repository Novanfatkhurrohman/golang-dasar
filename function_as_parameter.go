package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFilter := filter(name)
	fmt.Println("hello", nameFilter)
}
func spamFilter(name string) string {
	if name == "Anjing" {
		return "....."
	} else {
		return name
	}
}
func main() {
	filtere := spamFilter
	sayHelloWithFilter("novan", (filtere))
	fmt.Println("-----------")
	sayHelloWithFilter("Anjing", spamFilter)
}
