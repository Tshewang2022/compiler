package main

import (
	"compiler/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s ! This is the monkey programming language !\n", user.Username)

	fmt.Printf("Fell free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)

}

// now begins the chapter 2
// parsing=> is a software components that takes input data and builds a data structure, giving the structural representation of the input
// checking the correctness of syntax in the process
// changing the particular inputs into the particular data structure is called parsing;
// when writing parsing, two approach are being used, top down ,and bottom up approach
// expression produce values
// identifier => dont produce values
