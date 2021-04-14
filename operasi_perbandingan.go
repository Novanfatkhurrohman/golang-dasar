package main

import "fmt"

func main() {
	var (
		name1       = "novan"
		name2       = "stone"
		result bool = name1 == name2
		value1      = 100
		value2      = 200
	)
	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 >= value2)
	fmt.Println(value1 <= value2)
	fmt.Println(value1 != value2)
	fmt.Println(result)
}
