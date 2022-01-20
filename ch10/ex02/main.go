package main

import (
	"archive/tar"
	"archive/zip"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fileName := os.Args[1]
	ext := filepath.Ext(fileName)
	f, err := os.Open(fileName)
	defer f.Close()
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}
}

func archiveReader(fileName string) error {
	ext := filepath.Ext(fileName)
	if ext != "zip" && ext != "tar" {
		return fmt.Errorf("Can not use %s archive format", ext)
	}

	switch ext {
	case "zip":
		r, err := zip.OpenReader(fileName)
	case "tar":
		r, err := tar.OpenReader(fileName)
	}

}
