package main

import (
	"fmt"
	"monkeylang/repl"
	"os"
	user "os/user"
)

func main() {
	us, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", us.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
