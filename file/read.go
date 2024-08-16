package file

import (
	"fmt"
	"os"
)

func Read() {
    data, err := os.ReadFile("./file/output.txt")
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }
    fmt.Println(string(data))
}
