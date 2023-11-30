package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	//читаем содержимое текущего и всех вложенных каталогов
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}
