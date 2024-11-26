package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func TestAddUsageFlag(t *testing.T) {
	var buffer bytes.Buffer

	expectedOutput := AddUsage

	addCommand := NewAddCommand(&buffer)

	addCommand.flags.Usage()

	output := strings.TrimSpace(buffer.String())

	if output != expectedOutput {
		t.Errorf("Expected output: %q, got: %q", expectedOutput, output)
	}
}