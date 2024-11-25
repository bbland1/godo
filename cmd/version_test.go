package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func TestVersionUsageFlag(t *testing.T) {
	var buffer bytes.Buffer

	expectedOutput := HelpUsage

	versionCommand := NewVersionCommand(&buffer)

	versionCommand.flags.Usage()

	output := strings.TrimSpace(buffer.String())

	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}
}