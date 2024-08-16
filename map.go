package main

import (
	"fmt"
)

func Map() {
	// Deklarasi dan inisialisasi map dengan tipe key string dan value int
	m := make(map[string]int)

	// Menambahkan elemen ke map
	m["apple"] = 10
	m["banana"] = 20
	m["cherry"] = 30

	// Menampilkan map
	fmt.Println("Map:", m)

	// Mengakses nilai dari map menggunakan key
	fmt.Println("Value for 'banana':", m["banana"])

	// Menghapus elemen dari map
	delete(m, "banana")
	fmt.Println("Map after deleting 'banana':", m)

	// Mengecek apakah key ada di dalam map
	value, exists := m["cherry"]
	if exists {
		fmt.Println("Value for 'cherry':", value)
	} else {
		fmt.Println("'cherry' not found in map")
	}

	// Menambahkan elemen dengan key yang sama menggantikan nilai sebelumnya
	m["apple"] = 100
	fmt.Println("Map after updating 'apple':", m)

	// Iterasi melalui map
	fmt.Println("Iterating through map:")
	for key, value := range m {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}
