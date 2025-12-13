package main

import (
	"embed"
	"fmt"
)

var fileString string

var fileByte []byte

var folder embed.FS

func main() {
	p := fmt.Println

	p(fileString)
	p(string(fileByte))

	content1, _ := folder.ReadFile("folder/file1.hash")
	p(string(content1))

	content2, _ := folder.ReadFile("folder/file2.hash")
	p(string(content2))
}
