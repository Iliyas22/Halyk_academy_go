package main

import (
	"flag"
	"fmt"
)

type randomName struct {
	dbPassword string
	dbPort     int
}

func main() {

	cfg := randomName{}
	flag.StringVar(&cfg.dbPassword, "db-password", "notapassword", "password to connect to db")
	flag.IntVar(&cfg.dbPort, "db-port", 1431, "part to connect to db")

	flag.Parse()
	newArgFirst := flag.Args()
	//newArg := flag.Arg(0)
	fmt.Println(cfg)
	//println(newArg)
	fmt.Println(newArgFirst)

	fmt.Println("Enter Your First Roman values: ")

	// var then variable name then variable type
	var number string

	// Taking input from user
	fmt.Scanln(&number)
	fmt.Println("Enter Second Roman values: ")
	var second string
	fmt.Scanln(&second)

	// Print function is used to
	// display output in the same line
	fmt.Print("Your Full Name is: ")

	// Addition of two string
	fmt.Print(number + " " + second)

}
