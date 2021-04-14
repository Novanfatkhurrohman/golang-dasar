package main

import "fmt"

func getGodbye(name string) string {
	return "Godbye" + name
}

func main() {
	godbye := getGodbye
	fmt.Println(godbye(" novan"))
}
