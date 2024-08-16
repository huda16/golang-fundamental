package main

import (
	"fmt"
)

func Array() {
	// Deklarasi array dengan ukuran 5
	var arr1 [5]int
	fmt.Println("Array 1 (default values):", arr1)

	// Mengisi nilai ke dalam array
	arr1[0] = 10
	arr1[1] = 20
	arr1[2] = 30
	arr1[3] = 40
	arr1[4] = 50
	fmt.Println("Array 1 (after filling):", arr1)

	// Deklarasi dan inisialisasi array sekaligus
	arr2 := [3]string{"Apple", "Banana", "Cherry"}
	fmt.Println("Array 2:", arr2)

	// Mengakses elemen array
	fmt.Println("Element at index 1 in Array 2:", arr2[1])

	// Panjang array
	fmt.Println("Length of Array 2:", len(arr2))
}
