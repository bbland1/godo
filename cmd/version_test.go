package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func TestVersionUsageFlag(t *testing.T) {
	var buffer bytes.Buffer

	expectedOutput := VersionUsage

	versionCommand := NewVersionCommand(&buffer)

	versionCommand.flags.Usage()

	output := strings.TrimSpace(buffer.String())

	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}
}

func TestVersionCommandFlag(t *testing.T) {
	var buffer bytes.Buffer
	versionCommand := NewVersionCommand(&buffer)

	if versionCommand.flags.Name() != "version" {
		t.Errorf("NewVersionCommand flag name = %q, want it to be %q", versionCommand.flags.Name(), "version")
	}
}

func TestVersionInfo(t *testing.T) {
	var buffer bytes.Buffer

	expectedOutput := "goDo vblank"

	versionCommand := NewVersionCommand(&buffer)

	versionCommand.Execute(versionCommand, nil)

	output := strings.TrimSpace(buffer.String())

	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}

}

func TestVersionInfoVerbose(t *testing.T) {
	var buffer bytes.Buffer

	expectedOutput := "goDo v%blank, build: %blank"


	versionCommand := NewVersionCommand(&buffer)

	versionCommand.Execute(versionCommand, []string{"-verbose"})

	output := strings.TrimSpace(buffer.String())

	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}
}