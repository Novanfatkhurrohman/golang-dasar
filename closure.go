package main

import "fmt"

func main() {
	name := "budi"
	counter := 0
	increment := func() {
		name = "novan"
		fmt.Println("increment")
		counter++
	}
	increment()
	fmt.Println(counter)
	fmt.Println(name)
}
