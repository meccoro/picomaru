package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	dastname := "tempfile.txt"
	file, err := os.Create("old.txt")
	if err != nil {
		panic(err)
	}
	file.Write([]byte("picopico"))
	file.Close()
	copy(dastname, "old.txt")
	result, err := ioutil.ReadFile(dastname)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", string(result))
}

func copy(dstName, srcName string) {
	src, err := os.Open(srcName)
	if err != nil {
		panic(err)
	}
	defer src.Close()
	dst, err := os.Create(dstName)
	if err != nil {
		panic(err)
	}
	defer dst.Close()
	io.Copy(dst, src)
}
