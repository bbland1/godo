package cmd

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"text/tabwriter"
)

func TestHelpUsageFlag(t *testing.T) {
	var bufferOut bytes.Buffer
	var bufferErr bytes.Buffer
	var exitCode int

	expectedOutput := HelpUsage

	helpCommand := NewHelpCommand(&bufferOut, &bufferErr, &exitCode)

	helpCommand.flags.Usage()

	output := strings.TrimSpace(bufferOut.String())

	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}
}

func TestDisplayUserManual(t *testing.T) {
	var bufferOut bytes.Buffer

	exitCode := DisplayUserManual(&bufferOut)

	if exitCode != 0 {
		t.Errorf("Expected exit code to be: 0, got: %d", &exitCode)
	}

	expectedPhrases := []string{
		"Usage:\n  goDo [command] [options]",
		"Options:",
		"\nCommands:",
	}

	for _, phrase := range expectedPhrases {
		if !bytes.Contains(bufferOut.Bytes(), []byte(phrase)) {
			t.Errorf("Expected output to contain '%s', but it was missing", phrase)
		}
	}

}

func TestDisplayGreeting(t *testing.T) {
	var bufferOut bytes.Buffer
	var exitCode int

	expectedOutput := Greeting

	exitCode = DisplayGreeting(&bufferOut)

	output := strings.TrimSpace(bufferOut.String())

	if exitCode != 0 {
		t.Errorf("Expected exit code to be: 0, got: %d", &exitCode)
	}

	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}

}

func TestHelpCommandFlag(t *testing.T) {
	var bufferOut bytes.Buffer
	var bufferErr bytes.Buffer
	var exitCode int

	helpCommand := NewHelpCommand(&bufferOut, &bufferErr, &exitCode)

	if helpCommand.flags.Name() != "help" {
		t.Errorf("NewHelpCommand flag name = %q, want it to be %q", helpCommand.flags.Name(), "help")
	}
}

func TestHelpCommandOutput(t *testing.T) {
	var bufferOut bytes.Buffer
	var bufferErr bytes.Buffer
	var exitCode int
	
	helpCommand := NewHelpCommand(&bufferOut, &bufferErr, &exitCode)

	helpCommand.execute(helpCommand, nil)

	var buffExpectedOutput bytes.Buffer

	tw := tabwriter.NewWriter(&buffExpectedOutput, 0, 8, 2, ' ', 0)

	fmt.Fprintln(tw, "Usage:\n  goDo [command] [options]")

	fmt.Fprintln(tw, "\nOptions:")
	fmt.Fprintln(tw, "  -h\tShow more information about a command")
	fmt.Fprintln(tw, "  -verbose\tPrint detailed output when available")
	fmt.Fprintln(tw, "\nCommands:")

	for _, cmd := range registeredCommands {
		fmt.Fprintf(tw, "  %s\t- %s\n", cmd.GetName(), cmd.GetDescription())
	}

	tw.Flush()

	output := strings.TrimSpace(bufferOut.String())
	expectedOutput := strings.TrimSpace(buffExpectedOutput.String())

	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}
}
