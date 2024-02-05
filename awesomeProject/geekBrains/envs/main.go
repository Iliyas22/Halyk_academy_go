package main

import (
	"fmt"
	"os"
)

func main() {

	ShowPath := os.Getenv("GOPATH")
	os.Setenv("myval", "myvar")
	fmt.Println(os.Getenv("myval"))
	fmt.Println(ShowPath, ": ")
}
