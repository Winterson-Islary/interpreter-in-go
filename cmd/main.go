package main

import (
	"fmt"
	"interpreter/pkg/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! Welcome to the R-E-P Loop!\n", user.Username)
	fmt.Printf("Type in the commands\n")
	repl.REPLStart(os.Stdin, os.Stdout)
}
