package file

import (
	"fmt"
	"os"
)

func Metadata() {
    fileInfo, err := os.Stat("./file/output.txt")
    if err != nil {
        fmt.Println("Error getting file info:", err)
        return
    }

    fmt.Println("File Name:", fileInfo.Name())
    fmt.Println("File Size:", fileInfo.Size())
    fmt.Println("Last Modified:", fileInfo.ModTime())
}
