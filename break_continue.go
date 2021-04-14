package main

import "fmt"

func main() {

	//break
	for i := 0; i < 5; i++ {
		if i == 5 {
			break
		}
		fmt.Println("perulangan ke ", i)
	}
	fmt.Println("--------------")
	//continue
	for l := 0; l < 10; l++ {
		if l%2 == 1 {
			continue
		}
		fmt.Println("perulangan ke ", l)
	}
}
