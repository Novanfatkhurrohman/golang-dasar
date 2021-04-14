package main

import "fmt"

func main() {
	type (
		NoKtp   string
		Merried bool
	)
	var (
		ktpNovan   NoKtp   = "111111111111"
		merriedKuy Merried = true
	)
	fmt.Println(ktpNovan)
	fmt.Println(merriedKuy)
	fmt.Println(NoKtp("2222222222222"))
}
