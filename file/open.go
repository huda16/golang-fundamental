package file

import (
	"fmt"
	"os"
)

func Open() {
    file, err := os.Open("./file/output.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()
}
