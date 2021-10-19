package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func GetAppPath() string {

	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	// path, _ := os.Getwd()
	index := strings.LastIndex(path, string("/"))

	log.Printf("path=%v, index=%v", path, index)

	return path[:index]
}
