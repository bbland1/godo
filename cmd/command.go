package cmd

import (
	"flag"
)

type Command struct {
	flag    *flag.FlagSet
	Execute func(cmd *Command, args []string)
}
