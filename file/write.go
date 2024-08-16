package file

import (
	"fmt"
	"os"
)

func Write() {
    file, err := os.Create("./file/output.txt")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close()

    data := []byte("Hello, Golang!")
    _, err = file.Write(data)
    if err != nil {
        fmt.Println("Error writing to file:", err)
    }
}
