package file

import (
	"fmt"
	"os"
)

func Delete() {
    err := os.Remove("./file/output.txt")
    if err != nil {
        fmt.Println("Error deleting file:", err)
    }
}