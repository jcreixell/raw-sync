package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var JPGPath = os.Getenv("JPG_PATH")
var RAWPath = os.Getenv("RAW_PATH")

func main() {
	err := filepath.WalkDir(RAWPath, processFolder)
	if err != nil {
		log.Fatal(err)
	}
}

func processFolder(path string, d os.DirEntry, err error) error {
	info, err := d.Info()
	if err != nil {
		return err
	}

	dir := filepath.Dir(path)
	base := filepath.Base(path)
	ext := filepath.Ext(path)

	if !info.IsDir() {
		if ext == ".raf" {
			jpg := filepath.Join(strings.Replace(dir, RAWPath, JPGPath, 1), strings.Replace(base, ext, ".jpg", 1))
			_, err = os.Stat(jpg)
			if os.IsNotExist(err) {
				fmt.Println(jpg)
			}
		}
	}
	return nil
}
