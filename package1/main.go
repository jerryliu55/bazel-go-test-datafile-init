package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"path/filepath"
)

func init() {
	dat, err := ioutil.ReadFile("./etc/data")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Print(string(dat))
}

func TestSomething(t *testing.T) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)

	dat, err := ioutil.ReadFile("./etc/data")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Print(string(dat))
}
