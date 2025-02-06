package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func TestVersionUsageFlag(t *testing.T) {
	var bufferOut bytes.Buffer
	var bufferErr bytes.Buffer

	expectedOutput := VersionUsage

	versionCommand := NewVersionCommand(&bufferOut, &bufferErr)

	versionCommand.flags.Usage()

	output := strings.TrimSpace(bufferOut.String())

	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}
}

func TestVersionCommandFlag(t *testing.T) {
	var bufferOut bytes.Buffer
	var bufferErr bytes.Buffer
	versionCommand := NewVersionCommand(&bufferOut, &bufferErr)

	if versionCommand.flags.Name() != "version" {
		t.Errorf("NewVersionCommand flag name = %q, want it to be %q", versionCommand.flags.Name(), "version")
	}
}

func TestVersionInfo(t *testing.T) {
	var bufferOut bytes.Buffer
	var bufferErr bytes.Buffer

	expectedOutput := "goDo vblank"

	versionCommand := NewVersionCommand(&bufferOut, &bufferErr)

	versionCommand.execute(versionCommand, nil)

	output := strings.TrimSpace(bufferOut.String())

	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}

}

func TestVersionInfoVerbose(t *testing.T) {
	var bufferOut bytes.Buffer
	var bufferErr bytes.Buffer

	expectedOutput := "goDo vblank, build: blank"

	versionCommand := NewVersionCommand(&bufferOut, &bufferErr)

	versionCommand.Init([]string{"-verbose=true"})
	versionCommand.Run()

	output := strings.TrimSpace(bufferOut.String())

	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}
}
