package FileUtils

import (
	"fmt"
	"log"
	"os"
)

func CreateFauxFile(path string) {
	var _, err = os.Stat(path)
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err){
			return
		}
		defer file.Close()
	}
	fmt.Println("File created successfully ", path)
}

func WriteToFile(path string, content string){
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err){
		log.Fatal(err.Error())
		return
	}
	defer file.Close()
	_, err = file.WriteString(content)

	if isError(err){
		log.Fatal(err.Error())
		return
	}
}

type Document struct {
	Title string
	Size  int
	Path  string
}

func newDocument(title string, size int, address string) *Document {
	doc := Document{Title: title, Size: size, Path: address}
	return &doc
}
func isError(err error) bool {
	if err != nil {
		return true
	}
	return false
}