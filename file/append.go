package file

import (
	"fmt"
	"os"
)

func Append() {
    file, err := os.OpenFile("./file/output.txt", os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    _, err = file.WriteString("Appending this line.\n")
    if err != nil {
        fmt.Println("Error writing to file:", err)
    }
}
