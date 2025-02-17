package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		usage()
	}
	cmd := os.Args[1]
	if cmd != "testfile" {
		usage()
	}

	path, err := filepath.Abs(os.Args[2])
	if err != nil {
		die(err)
	}
	info, err := os.Stat(path)
	if err == nil && info.IsDir() {
		fmt.Fprintln(os.Stderr, "aviator: path is a directory")
		usage()
	}
	ext := filepath.Ext(path)
	name, _, _ := strings.Cut(path, ext)
	fmt.Println(name + "_test" + ext)
}

func usage() {
	fmt.Fprintln(os.Stderr, "usage: aviator [testfile] [path to file]")
	os.Exit(2)
}

func die(err error) {
	fmt.Fprintf(os.Stderr, "aviator: %v\n", err)
	os.Exit(1)
}
