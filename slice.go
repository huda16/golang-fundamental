package main

import (
	"fmt"
)

func Slice() {
	// Deklarasi dan inisialisasi array
	arr := [6]int{1, 2, 3, 4, 5, 6}

	// Membuat slice dari array
	s1 := arr[1:4] // Mengambil elemen dari indeks 1 sampai 3 dari array arr
	fmt.Println("Slice 1:", s1)
	fmt.Printf("Length of Slice 1: %d\n", len(s1))
	fmt.Printf("Capacity of Slice 1: %d\n", cap(s1))

	// Mengubah nilai dalam slice
	s1[0] = 20
	fmt.Println("Array after modifying Slice 1:", arr)
	fmt.Println("Updated Slice 1:", s1)

	// Membuat slice dengan menggunakan make
	s2 := make([]int, 3, 5) // Membuat slice dengan length 3 dan capacity 5
	fmt.Println("\nSlice 2:", s2)
	fmt.Printf("Length of Slice 2: %d\n", len(s2))
	fmt.Printf("Capacity of Slice 2: %d\n", cap(s2))

	// Mengisi nilai ke dalam slice
	s2[0] = 10
	s2[1] = 20
	s2[2] = 30
	fmt.Println("Updated Slice 2:", s2)

	// Menambahkan elemen ke slice dengan append
	s2 = append(s2, 40, 50)
	fmt.Println("\nSlice 2 after append:", s2)
	fmt.Printf("Length of Slice 2 after append: %d\n", len(s2))
	fmt.Printf("Capacity of Slice 2 after append: %d\n", cap(s2))
}
