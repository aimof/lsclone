package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	flag.Parse()
	targetPath, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	isDir, err := isDirectory(targetPath)
	if err != nil {
		log.Fatalln(err)
	}
	if !isDir {
		log.Fatalln(err)
	}
	files, err := ioutil.ReadDir(targetPath)
	if err != nil {
		log.Fatalln(err)
	}
	for _, file := range files {
		var fileName string = file.Name()
		if file.IsDir() {
			fileName = fileName + "/"
		}
		if *pOpt && (*aOpt || *allOpt) && file.Name()[0] == byte('.') {
			fmt.Printf(fileName + "  ")
		} else if file.Name()[0] != byte('.') {
			fmt.Printf(fileName + "  ")
		}
	}
	fmt.Printf("\n")
}

func isDirectory(path string) (isDir bool, err error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return info.IsDir(), nil
}
