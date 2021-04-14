package main

import "fmt"

func main() {
	//array ini array := [5]int	{1,2,3,4,5}
	//slice ini slice :=[]int{1,2,3,4,5}
	var (
		months = [...]string{
			"januari",
			"febuari",
			"maret",
			"april",
			"mei",
			"juni",
			"juli",
			"agustus",
			"september",
			"oktobe",
			"novamber",
			"desember",
		}
		slice1 = months[4:7]
		slice2 = months[10:]
		slice3 = append(slice2, "novan")
	)
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	months[5] = "diubah"
	fmt.Println(slice1)

	fmt.Println("-------")
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println("-------")
	slice3[1] = "bukan"
	fmt.Println("-------")
	fmt.Println(slice3)
	fmt.Println("-------")
	fmt.Println(slice2)
	fmt.Println(months)
	//penambahan lebiih aman
	newSlice := make([]string, 2, 5)
	newSlice[0] = "script"
	newSlice[1] = "stone"

	fmt.Println("-------")
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	//copy slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println("-------")
	fmt.Println(copySlice)
}
