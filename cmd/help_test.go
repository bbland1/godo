package cmd

import (
	// "fmt"
	"bytes"
	"os"
	"testing"
)

// TestDisplayUserManual tests if DisplayUserManual prints the correct output.
func TestDisplayUserManual(t *testing.T) {
	expectedOutput := UserManual + "\n"

	// Create a pipe to capture the output
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Failed to create pipe: %v", err)
	}

	// Save the original os.Stdout and defer restoration
	oldOut := os.Stdout
	os.Stdout = w
	defer func() {
		os.Stdout = oldOut
		w.Close()
		r.Close()
	}()

	// Run the function that outputs to os.Stdout
	DisplayUserManual()

	// Close the writer to complete the capture, and read the output
	w.Close()
	var buffer bytes.Buffer
	_, err = buffer.ReadFrom(r)
	if err != nil {
		t.Fatalf("Failed to read from buffer: %v", err)
	}

	// Check if the output is as expected
	if buffer.String() != expectedOutput {
		t.Errorf("DisplayUserManual() = %q, want %q", buffer.String(), expectedOutput)
	}

}

// TestDisplayGreeting tests if DisplayGreeting prints the correct output.
func TestDisplayGreeting(t *testing.T) {
	expectedOutput := Greeting + "\n"

	// Create a pipe to capture the output
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Failed to create pipe: %v", err)
	}

	// Save the original os.Stdout and defer restoration
	oldOut := os.Stdout
	os.Stdout = w
	defer func() {
		os.Stdout = oldOut
		w.Close()
		r.Close()
	}()

	// Run the function that outputs to os.Stdout
	DisplayGreeting()

	// Close the writer to complete the capture, and read the output
	w.Close()
	var buffer bytes.Buffer
	_, err = buffer.ReadFrom(r)
	if err != nil {
		t.Fatalf("Failed to read from buffer: %v", err)
	}

	// Check if the output is as expected
	if buffer.String() != expectedOutput {
		t.Errorf("Display Greeting() = %q, want %q", buffer.String(), expectedOutput)
	}

}

// TestHelpCommandFlag tests if the help command for the application is properly set as a flag
func TestHelpCommandFlag(t *testing.T) {
	helpCommand := NewHelpCommand()

	if helpCommand.flags.Name() != "help" {
		t.Errorf("NewHelpCommand flag name = %q, want it to be %q", helpCommand.flags.Name(), "help")
	}
}

// TestHelpCommandOutput tests if the help command properly prints the UserManual when Executed.
func TestHelpCommandOutput(t *testing.T) {
	helpCommand := NewHelpCommand()

	expectedOutput := UserManual + "\n"
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Failed to create pipe: %v", err)
	}

	oldOut := os.Stdout
	os.Stdout = w
	defer func() {
		os.Stdout = oldOut
		w.Close()
		r.Close()
	}()

	helpCommand.Execute(helpCommand, nil)

	w.Close()
	var buffer bytes.Buffer
	_, err = buffer.ReadFrom(r)
	if err != nil {
		t.Fatalf("Failed to read from buffer: %v", err)
	}

	if buffer.String() != expectedOutput {
		t.Errorf("When the NewHelpCommand is used it should print out UserManual = %q, want %q", buffer.String(), expectedOutput)
	}
}
