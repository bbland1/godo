package cmd

import (
	"flag"
	"testing"
)

func TestCommandInit(t *testing.T) {
	cmd := &Command{
		flags: flag.NewFlagSet("tester", flag.ContinueOnError),
	}

	cmd.flags.String("name", "", "a test of string flag")

	sampleArgs := []string{"-name", "your name"}

	err := cmd.Init(sampleArgs)

	if err != nil {
		t.Errorf("Init returned and error in the set up of a command, %v", err)
	}

	nameFlag := cmd.flags.Lookup("name").Value.String()

	if nameFlag != "your name" {
		t.Errorf("Init did not properly parse the %q flag, got= %q, want= %q", "name", nameFlag, "your name")
	}

}

func TestCommandCalled(t *testing.T) {
	cmd := &Command{
		flags: flag.NewFlagSet("tester", flag.ContinueOnError),
	}

	cmd.flags.String("name", "", "a test of string flag")

	if cmd.Called() {
		t.Errorf("Called() should return false if method is before Init(), got= %t", cmd.Called())
	}

	sampleArgs := []string{"-name", "your name"}

	err := cmd.Init(sampleArgs)
	if err != nil {
		t.Errorf("Init returned and error in the set up of a command, %v", err)
	}

	if !cmd.Called() {
		t.Errorf("Called() should return true if method is after Init(), got= %t", cmd.Called())
	}

}

func TestCommandRun(t *testing.T) {
	var executed bool
	var passedArgs []string

	cmd := &Command{
		flags: flag.NewFlagSet("tester", flag.ContinueOnError),
		Execute: func(cmd *Command, args []string) {
			executed = true
			passedArgs = args
		},
	}

	cmd.flags.String("name", "", "a test of string flag")

	sampleArgs := []string{"-name", "your name", "extra", "non-parsed/non-flag", "args"}

	err := cmd.Init(sampleArgs)
	if err != nil {
		t.Errorf("Init() returned and error in the set up of a command, %v", err)
	}

	cmd.Run()

	if !executed {
		t.Errorf("Run() should utilize the Execute function setting 'executed' to true, got= %t, want= true", executed)
	}

	expectedArgs := []string{"extra", "non-parsed/non-flag", "args"}
	if len(passedArgs) != len(expectedArgs) {
		t.Errorf("Args not the expected values got= %v, want= %v", passedArgs, expectedArgs)
	}

	for i, arg := range expectedArgs {
		if passedArgs[i] != arg {
			t.Errorf("Expected arg[%d] to be %q, got %q", i, arg, passedArgs[i])
		}

	}

}
