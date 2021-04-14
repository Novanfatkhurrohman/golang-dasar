package main

import "fmt"

func main() {
	name := "novan"

	switch name {
	case "novan":
		fmt.Println("hello novan")
	case "stone":
		fmt.Println("Hello stone")
	default:
		fmt.Println("tidak ada nama")
	}
	//sederhana
	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("nama terlalu panjang")
	// case false:
	// 	fmt.Println("nama sudah benar")
	//}
	//versi sederhana
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("nama terlalu panjang")
	case length > 5:
		fmt.Println("nama lumayan panjang")

	default:
		fmt.Println("nama sudah benar")
	}
}
