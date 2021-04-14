package main

import "fmt"

func recursiveFunc(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * recursiveFunc(value-1)
	}
}
func main() {
	recrusive := recursiveFunc(5)
	fmt.Println(recrusive)

}
