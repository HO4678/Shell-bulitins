// builtins/pwd_test.go

package builtins

import (
	"bytes"
	"strings"
	"testing"
)

func TestPrintWorkingDirectory(t *testing.T) {
	// Redirect standard output to capture printed text
	var buf bytes.Buffer

	// Call the PrintWorkingDirectory function
	PrintWorkingDirectory(&buf) // Pass the bytes.Buffer directly

	// Get the printed output
	output := buf.String()

	// Check if the output is not empty
	if output == "" {
		t.Error("PrintWorkingDirectory did not print anything")
	}

	// Ensure there are no errors printed
	if containsError(output) {
		t.Errorf("PrintWorkingDirectory printed an error: %s", output)
	}

	// Add a custom writer to capture output
	var customWriter customWriter
	PrintWorkingDirectory(&customWriter)

	// Check if the output is not empty
	if customWriter.written == "" {
		t.Error("PrintWorkingDirectory did not write to the custom writer")
	}

	// Test output consistency between the standard and custom writers
	if output != customWriter.written {
		t.Error("PrintWorkingDirectory output inconsistency between standard and custom writers")
	}

	// Test if the output contains the correct path separator
	if !strings.Contains(output, "/") && !strings.Contains(output, "\\") {
		t.Error("PrintWorkingDirectory output does not contain path separator")
	}

	// TODO: Add more test cases as needed
}

// Helper function to check if an error is present in the printed output
func containsError(output string) bool {
	return len(output) > 7 && output[:7] == "Error: "
}

// customWriter is a custom io.Writer implementation for testing
type customWriter struct {
	written string
}

func (c *customWriter) Write(p []byte) (n int, err error) {
	c.written = string(p)
	return len(p), nil
}
