package main

import (
	"fmt"
)

func main() {
	//for
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}
	fmt.Println("-------")
	//for statement
	for wes := 1; wes <= 10; wes++ {
		fmt.Println("Perulangan ke ", wes)
	}
	fmt.Println("-------")

	slice := []string{"novan", "script", "stone", "golang"}

	for i := 0; i < len(slice); i++ {
		fmt.Println("Perulangan ke ", slice)
	}

	fmt.Println("-------")
	//for rank

	names := []string{"novan", "script", "stone", "golang"}
	for index, name := range names {
		fmt.Println("index ke ", index, "=", name)

	}
	fmt.Println("-------")
	for _, value := range slice {
		fmt.Println(value)
	}
	fmt.Println("-------")
	person := make(map[string]string)
	person["name"] = "novan"
	person["title"] = "golang"

	for key, values := range person {
		fmt.Println(key, "=", values)
	}

}
