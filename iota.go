package main

import (
	"fmt"
)

const (
    _ = iota             // 0
    KB = 1 << (10 * iota) // 1 KB
    MB                    // 2 MB
    GB                    // 3 GB
)

func Iota() {
    fmt.Println("Nilai KB:", KB) // Mencetak nilai KB
    fmt.Println("Nilai MB:", MB) // Mencetak nilai MB
    fmt.Println("Nilai GB:", GB) // Mencetak nilai GB
}
