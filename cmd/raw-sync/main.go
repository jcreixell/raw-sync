package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/jcreixell/raw-sync/pkg/utils"
)

var cfg struct {
	jpgPath          *string
	rawPath          *string
	rawExtension     *string
	destionationPath *string
	dryRun           *bool
}

func main() {
	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	cfg.jpgPath = fs.String("jpg-path", ".", "Path to folder containing JPG files.")
	cfg.rawPath = fs.String("raw-path", ".", "Path to folder containing RAW files.")
	cfg.rawExtension = fs.String("raw-format", "raf", "Extension of the raw files.")
	cfg.destionationPath = fs.String("destination-path", "", "Path were duplicated files will be stored.")
	cfg.dryRun = fs.Bool("dry-run", false, "Run the script in a non-destructive way.")

	_ = fs.Parse(os.Args[1:])

	if !*cfg.dryRun && len(*cfg.destionationPath) == 0 {
		log.Fatal("destination path not specified, aborting.")
	}

	if !*cfg.dryRun {
		if _, err := os.Stat(*cfg.destionationPath); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(*cfg.destionationPath, os.ModePerm)
			if err != nil {
				log.Fatalf("destination path does not exist and could not be created: %v.", err)
			}
		}
	}

	err := filepath.WalkDir(*cfg.rawPath, processFolder)
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
		if ext == "."+*cfg.rawExtension {
			jpg := filepath.Join(strings.Replace(dir, *cfg.rawPath, *cfg.jpgPath, 1), strings.Replace(base, ext, ".jpg", 1))
			_, err = os.Stat(jpg)
			if os.IsNotExist(err) {
				if *cfg.dryRun {
					fmt.Println(path)
				} else {
					dest := filepath.Join(*cfg.destionationPath, base)
					err = utils.Move(path, dest)
					if err != nil {
						return fmt.Errorf("could not move %v: %v", path, err)
					}
				}
			}
		}
	}
	return nil
}
