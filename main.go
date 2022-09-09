package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func ListFiles(string folder) {
	f := filepath.Join(os.Getenv("HOME"), folder)
	filepath.Walk(f, func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})

}

func main() {
	ListFiles()
}
