package main

import "fmt"

// Mendefinisikan fungsi dengan nama greet yang tidak mengembalikan nilai
func greet(name string) {
	fmt.Println("Hello,", name)
}

// Mendefinisikan fungsi dengan nama add yang mengembalikan jumlah dua angka
func add(a int, b int) int {
	return a + b
}

func FunctionDataType() {
	// Memanggil fungsi greet
	greet("Alice") // Output: Hello, Alice

	// Memanggil fungsi add dan menyimpan hasilnya
	sum := add(5, 3)
	fmt.Println("Sum:", sum) // Output: Sum: 8
}
