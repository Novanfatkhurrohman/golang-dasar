package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "novan"
	names[1] = "script"
	names[2] = "stone"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90, 95, 80,
	}
	fmt.Println((values))
	fmt.Println(len(values))

}
