package cmd

import (
	"flag"
	"fmt"
	"io"
	"text/tabwriter"
)

var registeredCommands = make(map[string]Command)

// Command interface to define the structure for all CLI commands
type Command interface {
	Init(args []string) error
	Run()
	Called() bool
	GetName() string
	GetDescription() string
}

// BaseCommand defines a structure for a command that has flags and an execution function.
type BaseCommand struct {
	name        string
	description string
	flags       *flag.FlagSet
	output      io.Writer
	errOutput   io.Writer
	execute     func(cmd *BaseCommand, args []string)
}

/* cmd is a method receiver that works like `self` or `this` in JS
basically saying do the method on the passed thing/object of the method.
In this case it will be a command, the *Command is making the type the pointer of the Command struct so that changes made are saved made to the command instance itself
returns an error if the flags can't be parsed */

// Init initializes the command by parsing the provided arguments.
// Sets up the flags using the arguments and returns an error if parsing fails.
func (cmd *BaseCommand) Init(args []string) error {
	return cmd.flags.Parse(args)
}

// Called checks if the command's flags have been parsed.
func (cmd *BaseCommand) Called() bool {
	return cmd.flags.Parsed()
}

// Run executes the command by calling the Execute function with the command instance
// and the parsed arguments. It triggers the execution logic that is passed during the command setup.
func (cmd *BaseCommand) Run() {
	cmd.execute(cmd, cmd.flags.Args())
}

func (cmd *BaseCommand) GetName() string {
	return cmd.name
}

func (cmd *BaseCommand) GetDescription() string {
	return cmd.description
}

func RegisterCommand(cmd Command) {
	registeredCommands[cmd.GetName()] = cmd
}

func GetCommand(name string) (Command, bool) {
	cmd, exists := registeredCommands[name]
	return cmd, exists
}

func ListCommands(w io.Writer) {
	tw := tabwriter.NewWriter(w, 0, 8, 2, ' ',0)
	fmt.Fprintf(w, "commands:\n")
	for _, cmd := range registeredCommands {
		fmt.Fprintf(tw, "  %s\t- %s\n", cmd.GetName(), cmd.GetDescription())
	}

	tw.Flush()
}
