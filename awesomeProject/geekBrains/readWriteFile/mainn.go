package main

import "os"

func main() {
	os.OpenFile("file.txt", os.O_CREATE, 777)
}
