package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type myFile struct {
	name string
	path string
}

func main() {
	dir := "sample"
	var toRename []myFile
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if _, err := match(info.Name()); err == nil {
			toRename = append(toRename, myFile{
				name: info.Name(),
				path: path,
			})
		}

		return nil
	})

	for _, orig := range toRename {
		var n myFile
		var err error
		n.name, err = match(orig.name)
		if err != nil {
			fmt.Println("Error matching:", orig.path, err.Error())
			continue
		}
		n.path = filepath.Join(filepath.Dir(orig.path), n.name)
		fmt.Printf("mv %s => %s\n", orig.path, n.path)
		err = os.Rename(orig.path, n.path)
		if err != nil {
			fmt.Println("Error renaming:", orig.path, err.Error())
		}
	}
}

// func foo() {
// 	dir := "./sample"
// 	files, err := ioutil.ReadDir(dir)
// 	if err != nil {
// 		panic(err)
// 	}
// 	count := 0
// 	var toChange []myFile
// 	for _, file := range files {
// 		if !file.IsDir() {
// 			_, err := match(file.Name())
// 			if err == nil {
// 				count++
// 				toChange = append(toChange, myFile{
// 					name: file.Name(),
// 					path: filepath.Join(dir, file.Name()),
// 				})
// 			}
// 		}
// 	}

// 	for _, orgTarget := range toChange {
// 		newFileName, err := match(orgTarget.name)
// 		if err != nil {
// 			panic(err)
// 		}
// 		newPath := filepath.Join(dir, newFileName)
// 		err = os.Rename(orgTarget.path, newPath)
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Printf("mv %s => %s\n", orgTarget.path, newPath)
// 	}
// }

func match(fileName string) (string, error) {
	// "birthday_001.txt" => ["birthday", "001", "txt"]
	pieces := strings.Split(fileName, ".")
	fileExtension := pieces[len(pieces)-1]
	tmp := strings.Join(pieces[0:len(pieces)-1], ".")
	pieces = strings.Split(tmp, "_")
	name := strings.Join(pieces[0:len(pieces)-1], "_")
	number, err := strconv.Atoi(pieces[len(pieces)-1])
	if err != nil {
		return "", fmt.Errorf("%s didn't match our pattern", fileName)
	}
	return fmt.Sprintf("%s - %d.%s", name, number, fileExtension), nil
}
