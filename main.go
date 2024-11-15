package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bbland1/goDo/cmd"
)

func main() {

	if len(os.Args) < 2 {
		cmd.DisplayGreeting()
		os.Exit(0)
	}

	command := os.Args[1]

	switch command {
	case "help":
		cmd.DisplayUserManual()
	default:
		fmt.Printf("unknown command passed to goDo: %s\n\n", command)
		flag.Usage()
		os.Exit(1)
	}

}
