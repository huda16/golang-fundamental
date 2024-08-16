package main

import "fmt"

func Panic() {
    x := -1
    if x < 0 {
        panic("Nilai x harus positif")
    }
    fmt.Println(x)
}