package main

import (
	"fmt"
)

// Mendeklarasikan struct dengan nama 'Person'
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func Struct() {
	// Membuat instance struct 'Person'
	p1 := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}

	// Menampilkan nilai-nilai field dari struct 'Person'
	fmt.Println("Person 1:")
	fmt.Println("First Name:", p1.FirstName)
	fmt.Println("Last Name:", p1.LastName)
	fmt.Println("Age:", p1.Age)

	// Mengakses dan mengubah field dari struct
	p1.Age = 31
	fmt.Println("Updated Age:", p1.Age)
}
