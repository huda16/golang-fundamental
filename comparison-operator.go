package main

import (
	"fmt"
)

func ComparisonOperator() {
    a := 10
    b := 20
    c := 10

    // Menggunakan operator perbandingan
    fmt.Println("a == b:", a == b) // false
    fmt.Println("a != b:", a != b) // true
    fmt.Println("a > b:", a > b)   // false
    fmt.Println("a < b:", a < b)   // true
    fmt.Println("a >= c:", a >= c) // true
    fmt.Println("a <= c:", a <= c) // true

    // Menggunakan operator perbandingan untuk kondisi
    if a < b {
        fmt.Println("a lebih kecil dari b")
    }

    if a == c {
        fmt.Println("a sama dengan c")
    }
}
