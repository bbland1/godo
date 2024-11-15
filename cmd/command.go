package cmd

import (
	"flag"
)

type Command struct {
	flags    *flag.FlagSet
	Execute func(cmd *Command, args []string)
}

// command is a method receiver that works like self or this in JS saying do the method on the passed thing/object of the method. in this case it will be a command, the *Command is making the type the pointer of the Command struct so that changes made are saved made to the command instance itself

// returns an error if the flags can't be parsed

// sets up the commands by setting up the flags
func (cmd *Command) Init(args []string) error {
	return cmd.flags.Parse(args)
}

func (cmd *Command) Called() bool {
	return cmd.flags.Parsed()
}

// run the actual command passed with reference to the pointer and utilizing the args
func (cmd *Command) Run() {
	cmd.Execute(cmd, cmd.flags.Args())
}