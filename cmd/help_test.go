package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func TestHelpUsageFlag(t *testing.T) {
	var buffer bytes.Buffer

	expectedOutput := HelpUsage

	helpCommand := NewHelpCommand(&buffer)

	helpCommand.flags.Usage()

	output := strings.TrimSpace(buffer.String())

	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}
}

func TestDisplayUserManual(t *testing.T) {
	var buffer bytes.Buffer

	expectedOutput := UserManual

	DisplayUserManual(&buffer)

	output := strings.TrimSpace(buffer.String())

	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}

}

func TestDisplayGreeting(t *testing.T) {
	var buffer bytes.Buffer

	expectedOutput := Greeting

	DisplayGreeting(&buffer)

	output := strings.TrimSpace(buffer.String())

	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}

}

func TestHelpCommandFlag(t *testing.T) {
	var buffer bytes.Buffer
	helpCommand := NewHelpCommand(&buffer)

	if helpCommand.flags.Name() != "help" {
		t.Errorf("NewHelpCommand flag name = %q, want it to be %q", helpCommand.flags.Name(), "help")
	}
}

func TestHelpCommandOutput(t *testing.T) {
	var buffer bytes.Buffer
	helpCommand := NewHelpCommand(&buffer)

	expectedOutput := UserManual

	helpCommand.Execute(helpCommand, nil)

	output := strings.TrimSpace(buffer.String())

	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}
}
