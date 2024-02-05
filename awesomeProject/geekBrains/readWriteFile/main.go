package main

import (
	"os"
)

func main() {
	file, err := os.OpenFile("myfilNew.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 547)
	if err != nil {
		panic("ds")
	}
	defer file.Close()
	file.Write([]byte("Hey22"))
}
