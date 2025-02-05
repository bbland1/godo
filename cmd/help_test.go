package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func TestHelpUsageFlag(t *testing.T) {
	var bufferOut bytes.Buffer
	var bufferErr bytes.Buffer

	expectedOutput := HelpUsage

	helpCommand := NewHelpCommand(&bufferOut, &bufferErr)

	helpCommand.flags.Usage()

	output := strings.TrimSpace(bufferOut.String())

	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}
}

func TestDisplayUserManual(t *testing.T) {
	var bufferOut bytes.Buffer

	expectedOutput := UserManual

	DisplayUserManual(&bufferOut)

	output := strings.TrimSpace(bufferOut.String())

	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}

}

func TestDisplayGreeting(t *testing.T) {
	var bufferOut bytes.Buffer

	expectedOutput := Greeting

	DisplayGreeting(&bufferOut)

	output := strings.TrimSpace(bufferOut.String())

	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}

}

func TestHelpCommandFlag(t *testing.T) {
	var bufferOut bytes.Buffer
	var bufferErr bytes.Buffer
	helpCommand := NewHelpCommand(&bufferOut, &bufferErr)

	if helpCommand.flags.Name() != "help" {
		t.Errorf("NewHelpCommand flag name = %q, want it to be %q", helpCommand.flags.Name(), "help")
	}
}

func TestHelpCommandOutput(t *testing.T) {
	var bufferOut bytes.Buffer
	var bufferErr bytes.Buffer
	helpCommand := NewHelpCommand(&bufferOut, &bufferErr)

	expectedOutput := UserManual

	helpCommand.execute(helpCommand, nil)

	output := strings.TrimSpace(bufferOut.String())

	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}
}
