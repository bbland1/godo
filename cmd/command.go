package cmd

import (
	"flag"
	"fmt"
	"io"
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

func NewBaseCommand(name, description string, stdout, stderr io.Writer, execute func(cmd *BaseCommand, args []string)) *BaseCommand {
	return &BaseCommand{
		name:        name,
		description: description,
		output:      stdout,
		errOutput:   stderr,
		flags:       flag.NewFlagSet(name, flag.ExitOnError),
		execute:     execute,
	}
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
	fmt.Fprintf(w, "Available commands:\n")
	for _, cmd := range registeredCommands {
		fmt.Fprintf(w, " %s - %s\n", cmd.GetName(), cmd.GetDescription())
	}
}
