// builtins/echo_test.go

package builtins

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEcho(t *testing.T) {
	// Redirect stdout to capture the output
	var buf bytes.Buffer

	// Test the Echo function with multiple arguments
	args := []string{"hello", "world"}
	Echo(&buf, args...)

	// Print the captured output for debugging
	fmt.Printf("Captured output: %q\n", buf.String())

	// Check if the output matches the expected result
	expectedOutput := "hello world\n"
	if buf.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, buf.String())
	}

	// Test the Echo function with a single argument
	buf.Reset()
	singleArg := []string{"single"}
	Echo(&buf, singleArg...)

	// Print the captured output for debugging
	fmt.Printf("Captured output: %q\n", buf.String())

	// Check if the output matches the expected result
	expectedSingleArg := "single\n"
	if buf.String() != expectedSingleArg {
		t.Errorf("Expected output: %s, got: %s", expectedSingleArg, buf.String())
	}

	// Test the Echo function with an empty argument list
	buf.Reset()
	emptyArgs := []string{}
	Echo(&buf, emptyArgs...)

	// Print the captured output for debugging
	fmt.Printf("Captured output: %q\n", buf.String())

	// Check if the output matches the expected result (empty line)
	expectedEmptyArgs := "\n"
	if buf.String() != expectedEmptyArgs {
		t.Errorf("Expected output: %s, got: %s", expectedEmptyArgs, buf.String())
	}

	// Test the Echo function with special characters
	buf.Reset()
	specialChars := []string{"$PATH", "%%", "#comment"}
	Echo(&buf, specialChars...)

	// Print the captured output for debugging
	fmt.Printf("Captured output: %q\n", buf.String())

	// Check if the output matches the expected result
	expectedSpecialChars := "$PATH %% #comment\n"
	if buf.String() != expectedSpecialChars {
		t.Errorf("Expected output: %s, got: %s", expectedSpecialChars, buf.String())
	}
	buf.Reset()
	numericArgs := []string{"123", "456", "789"}
	Echo(&buf, numericArgs...)
	expectedNumericArgs := "123 456 789\n"
	if buf.String() != expectedNumericArgs {
		t.Errorf("Expected output: %s, got: %s", expectedNumericArgs, buf.String())
	}
	buf.Reset()
	mixedArgs := []string{"abc", "123", "!@#$"}
	Echo(&buf, mixedArgs...)
	expectedMixedArgs := "abc 123 !@#$\n"
	if buf.String() != expectedMixedArgs {
		t.Errorf("Expected output: %s, got: %s", expectedMixedArgs, buf.String())
	}
	buf.Reset()
	repeatedArgs := []string{"apple", "orange", "apple", "banana"}
	Echo(&buf, repeatedArgs...)
	expectedRepeatedArgs := "apple orange apple banana\n"
	if buf.String() != expectedRepeatedArgs {
		t.Errorf("Expected output: %s, got: %s", expectedRepeatedArgs, buf.String())
	}
	buf.Reset()
	longArgs := []string{"This", "is", "a", "long", "argument", "with", "multiple", "words"}
	Echo(&buf, longArgs...)
	expectedLongArgs := "This is a long argument with multiple words\n"
	if buf.String() != expectedLongArgs {
		t.Errorf("Expected output: %s, got: %s", expectedLongArgs, buf.String())
	}

}
