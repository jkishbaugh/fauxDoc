package FileUtils

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func CreateFauxFile(document *Document) {

	dirPath := removeFileName(document.Path)
	dirErr := os.MkdirAll(dirPath, os.ModeDir)
	if isError(dirErr) {
		fmt.Println(dirErr.Error())
		return
	}
	chErr := os.Chdir(dirPath)
	if isError(chErr) {
		fmt.Println(chErr.Error())
		return
	}
	var file, err = os.Create(document.Title)
	if isError(err){
		fmt.Println("Create File Error: ", err.Error())
		return
	}
	file.Close()
	fmt.Println("File created successfully ", document.Path)
}

func WriteToFile(path string, content string){
	var file, err = os.OpenFile(path, os.O_CREATE, os.ModePerm)
	if isError(err){
		log.Fatal(err.Error())
		return
	}
	defer file.Close()
	fmt.Println("Writing to File: " , path)
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
	Content string
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

func removeFileName(path string) string{
	dir := strings.Split(path, "\\")
	return strings.Join(dir[:len(dir)-1], "\\")
}
