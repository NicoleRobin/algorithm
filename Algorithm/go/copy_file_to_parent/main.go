package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func main() {
	dirPath := "E:\\视频\\hh\\X-Art 2014"
	dirs := os.DirFS(dirPath)
	fs.WalkDir(dirs, ".", func(path string, d fs.DirEntry, err error) error {
		if d.Type().IsRegular() {
			oldPath := filepath.Join(dirPath, path)
			newPath := filepath.Join(dirPath, d.Name())
			fmt.Printf("oldPath:%s, newPath:%s\n", oldPath, newPath)
			err := os.Rename(oldPath, newPath)
			if err != nil {
				log.Fatal(err)
			}
		}
		return nil
	})
}
