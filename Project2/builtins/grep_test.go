// builtins/grep_test.go

package builtins

import (
	"bytes"
	"os"
	"testing"
)

func TestGrep(t *testing.T) {
	// Create a temporary test file
	testFile := "test_file.txt"
	content := `Line 1: Hello, World!
Line 2: This is a test.
Line 3: Testing Grep function.
Line 4: End of test.`

	err := os.WriteFile(testFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Error creating test file: %v", err)
	}
	defer os.Remove(testFile)

	// Capture stdout for testing
	var out bytes.Buffer

	// Test Grep with a valid pattern and file
	err = Grep(&out, "Testing", testFile)
	if err != nil {
		t.Errorf("Grep failed: %v", err)
	}

	// Check if the output contains the expected line
	expectedOutput := "Line 3: Testing Grep function.\n"
	if out.String() != expectedOutput {
		t.Errorf("Grep output does not match expected content:\nExpected: %s\nActual: %s",
			expectedOutput, out.String())
	}
	out.Reset()

	// Test Grep with a non-matching pattern
	err = Grep(&out, "NonExistentPattern", testFile)
	if err != nil {
		t.Errorf("Grep failed: %v", err)
	}

	// Check if the output is empty for non-matching pattern
	if out.String() != "" {
		t.Errorf("Grep output should be empty for non-matching pattern")
	}
	out.Reset()

	// Test Grep with an invalid number of arguments
	err = Grep(&out, "PatternOnly")
	if err == nil {
		t.Errorf("Expected an error for invalid number of arguments, but got none")
	}
}
