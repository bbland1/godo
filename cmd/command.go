package cmd

import (
	"flag"
)

// Command defines a structure for a command that has flags and an execution function.
// `flags` field contains the parsed command-line arguments
// `Execute` function runs the logic associated with the command
type Command struct {
	flags    *flag.FlagSet
	Execute func(cmd *Command, args []string)
}

/* cmd is a method receiver that works like `self` or `this` in JS 
basically saying do the method on the passed thing/object of the method. 
I
n this case it will be a command, the *Command is making the type the pointer of the Command struct so that changes made are saved made to the command instance itself

returns an error if the flags can't be parsed */

// Init initializes the command by parsing the provided arguments.
// Sets up the flags using the arguments and returns an error if parsing fails.
func (cmd *Command) Init(args []string) error {
	return cmd.flags.Parse(args)
}

// Called checks if the command's flags have been parsed.
func (cmd *Command) Called() bool {
	return cmd.flags.Parsed()
}

// Run executes the command by calling the Execute function with the command instance
// and the parsed arguments. It triggers the execution logic that is passed during the command setup.
func (cmd *Command) Run() {
	cmd.Execute(cmd, cmd.flags.Args())
}