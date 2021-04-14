package main

import "fmt"

func main() {
	name := "script stone"
	if name == "novan" {
		fmt.Println("hello novan")
	} else if name == "script stone" {
		fmt.Println("hey stone")
	} else {
		fmt.Println("tidak ada data")
	}
	if length := len(name); length > 5 {
		fmt.Println("nama terlalu panjang ")
	} else {
		fmt.Println("nama sudah benar")
	}

}
