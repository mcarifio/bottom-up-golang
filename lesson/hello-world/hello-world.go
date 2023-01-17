package main

import (
	"fmt"
	"os"
	"strings"
)

func defaultWho() string {
	return "world"
}

func main() {
	var who = strings.Join(os.Args[1:], " ") // pass the name on the command line e.g. `go run hello-world.go carif`
	if len(who) == 0 {
		// No command line arguments, therefore "world"
		who = defaultWho()
	}
	fmt.Println("hello, " + who) // traditional first program
}
