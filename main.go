package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		log.Fatal("Must include directory path as input to program")
	}

	dir := args[0]

	goFiles, err := getGoFiles(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range goFiles {
		structs, err := FindStructsInFile(f)
		if err != nil {
			fmt.Println(err)
		}
		if len(structs) > 0 {
			fmt.Printf("\n%s\n", f)

			for _, s := range structs {
				fmt.Println(s)
			}
		}
	}

}

func getGoFiles(path string) ([]string, error) {
	files := []string{}
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		if !info.IsDir() {
			name := info.Name()
			ext := filepath.Ext(name)
			if ext == ".go" {
				files = append(files, path)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return files, err
	}
	return files, nil
}
