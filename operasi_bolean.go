package main

import "fmt"

func main() {
	var (
		nilaiAkhir = 90
		nilaiAbsen = 80

		lulusNilaiAkhir bool = nilaiAkhir >= 80
		lulusNilaiAbsen bool = nilaiAbsen >= 80
		lulus           bool = lulusNilaiAkhir && lulusNilaiAbsen
	)
	fmt.Println(lulus)
	fmt.Println(nilaiAkhir >= 75 && nilaiAbsen >= 70)

}
