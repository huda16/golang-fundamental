package main

import (
	"errors"
	"fmt"
)

func Error() {
    err := errors.New("terjadi kesalahan")
    if err != nil {
        fmt.Println(err)
    }
}

type MyError struct {
    msg string
}

func (e *MyError) Error() string {
    return e.msg
}

func CustomError() {
    err := &MyError{"Operasi gagal"}
    if err != nil {
        fmt.Println(err)
    }
}

func divide(x, y int) (int, error) {
    if y == 0 {
        return 0, errors.New("pembagian dengan nol")
    }
    return x / y, nil
}

func ErrorExample() {
    result, err := divide(10, 0)
    if err != nil {
        fmt.Println("Terjadi kesalahan:", err)
        return
    }
    fmt.Println("Hasil:", result)
}