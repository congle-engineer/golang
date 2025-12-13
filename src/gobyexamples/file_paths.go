package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	p := fmt.Println

	path := filepath.Join("dir1", "dir2", "filename")
	p("path:", path)

	p(filepath.Join("dir1//", "filename"))
	p(filepath.Join("dir1/../dir1", "filename"))

	p("Dir(path):", filepath.Dir(path))
	p("Base(path):", filepath.Base(path))

	p(filepath.IsAbs("dir/file"))
	p(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	ext := filepath.Ext(filename)
	p(ext)

	p(strings.TrimSuffix(filename, ext))

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	p(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	p(rel)
}
