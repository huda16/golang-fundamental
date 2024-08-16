package main

import (
	"fmt"
)

func Pointer() {
	// Deklarasi variabel integer
	num := 10
	fmt.Println("Original value of num:", num)

	// Mendapatkan alamat memori dari variabel num
	ptr := &num
	fmt.Println("Pointer to num:", ptr)
	fmt.Println("Value at pointer address:", *ptr)

	// Mengubah nilai variabel melalui pointer
	*ptr = 20
	fmt.Println("Updated value of num:", num)
	fmt.Println("Value at pointer address after update:", *ptr)

	// Menggunakan pointer dalam fungsi
	increment(&num)
	fmt.Println("Value of num after increment function:", num)
}

// Fungsi yang menerima pointer ke integer
func increment(n *int) {
	*n++ // Mengubah nilai di alamat yang ditunjuk oleh pointer
}
