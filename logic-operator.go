package main

import "fmt"

func LogicOperator() {
	// Mendefinisikan variabel boolean
	a := true
	b := false

	// Operator AND (&&)
	andResult := a && b
	fmt.Printf("a && b: %v\n", andResult) // false

	// Operator OR (||)
	orResult := a || b
	fmt.Printf("a || b: %v\n", orResult) // true

	// Operator NOT (!)
	notA := !a
	notB := !b
	fmt.Printf("!a: %v, !b: %v\n", notA, notB) // false, true

	// Contoh penggunaan dalam kondisi
	if a && !b {
		fmt.Println("a adalah true dan b adalah false")
	}

	if a || b {
		fmt.Println("Setidaknya satu dari a atau b adalah true")
	}
}
