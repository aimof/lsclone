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
		if *pOpt && (*aOpt || *allOpt) && file.Name()[0] == byte('.') {
			fmt.Printf(file.Name() + "/  ")
		} else if (*aOpt || *allOpt) || file.Name()[0] == byte('.') {
			fmt.Printf(file.Name() + "  ")
		} else if file.Name()[0] != byte('.') {
			fmt.Printf(file.Name() + "  ")
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