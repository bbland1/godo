package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func TestVersionUsageFlag(t *testing.T) {
	var bufferOut bytes.Buffer
	var bufferErr bytes.Buffer
	var exitCode int

	expectedOutput := VersionUsage

	versionCommand := NewVersionCommand(&bufferOut, &bufferErr, &exitCode)

	versionCommand.flags.Usage()

	output := strings.TrimSpace(bufferOut.String())

	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}
}

func TestVersionCommandFlag(t *testing.T) {
	var bufferOut bytes.Buffer
	var bufferErr bytes.Buffer
	var exitCode int

	versionCommand := NewVersionCommand(&bufferOut, &bufferErr, &exitCode)

	if versionCommand.flags.Name() != "version" {
		t.Errorf("NewVersionCommand flag name = %q, want it to be %q", versionCommand.flags.Name(), "version")
	}
}

func TestVersionInfo(t *testing.T) {
	var bufferOut bytes.Buffer
	var bufferErr bytes.Buffer
	var exitCode int

	expectedOutput := "goDo vblank"

	versionCommand := NewVersionCommand(&bufferOut, &bufferErr, &exitCode)

	versionCommand.execute(versionCommand, nil)

	output := strings.TrimSpace(bufferOut.String())

	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}

}

func TestVersionInfoVerbose(t *testing.T) {
	var bufferOut bytes.Buffer
	var bufferErr bytes.Buffer
	var exitCode int


	expectedOutput := "goDo vblank, build: blank"

	versionCommand := NewVersionCommand(&bufferOut, &bufferErr, &exitCode)

	versionCommand.Init([]string{"-verbose=true"})
	versionCommand.Run()

	output := strings.TrimSpace(bufferOut.String())

	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}
}
