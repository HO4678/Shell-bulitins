// builtins/cat_test.go

package builtins

import (
	"bytes"
	"os"
	"testing"
)

func assertConcatenatedOutput(t *testing.T, out *bytes.Buffer, expected string) {
	t.Helper()
	if out.String() != expected {
		t.Errorf("ConcatenateFiles output does not match expected content:\nExpected: %s\nActual: %s",
			expected, out.String())
	}
}

func TestConcatenateFiles(t *testing.T) {
	t.Run("ConcatenateTwoFiles", func(t *testing.T) {
		file1 := "test_file1.txt"
		file2 := "test_file2.txt"

		content1 := "Content of file 1."
		content2 := "Content of file 2."

		err := os.WriteFile(file1, []byte(content1), 0644)
		if err != nil {
			t.Fatalf("Error creating test file 1: %v", err)
		}
		defer os.Remove(file1)

		err = os.WriteFile(file2, []byte(content2), 0644)
		if err != nil {
			t.Fatalf("Error creating test file 2: %v", err)
		}
		defer os.Remove(file2)

		var out bytes.Buffer
		err = ConcatenateFiles(&out, file1, file2)
		if err != nil {
			t.Errorf("ConcatenateFiles failed: %v", err)
		}
		assertConcatenatedOutput(t, &out, content1+content2)
	})

	t.Run("ConcatenateEmptyFile", func(t *testing.T) {
		emptyFile := "empty_file.txt"

		err := os.WriteFile(emptyFile, nil, 0644)
		if err != nil {
			t.Fatalf("Error creating empty test file: %v", err)
		}
		defer os.Remove(emptyFile)

		var out bytes.Buffer
		err = ConcatenateFiles(&out, emptyFile)
		if err != nil {
			t.Errorf("ConcatenateFiles failed: %v", err)
		}
		assertConcatenatedOutput(t, &out, "")
	})

	t.Run("ConcatenateNonExistentFile", func(t *testing.T) {
		var out bytes.Buffer
		err := ConcatenateFiles(&out, "nonexistent_file.txt")
		if err == nil {
			t.Error("Expected an error for non-existent file, but got none")
		}
	})

	t.Run("ConcatenateMultipleFiles", func(t *testing.T) {
		file1 := "file1.txt"
		file2 := "file2.txt"
		file3 := "file3.txt"

		content1 := "Content of file 1."
		content2 := "Content of file 2."
		content3 := "Content of file 3."

		err := os.WriteFile(file1, []byte(content1), 0644)
		if err != nil {
			t.Fatalf("Error creating test file 1: %v", err)
		}
		defer os.Remove(file1)

		err = os.WriteFile(file2, []byte(content2), 0644)
		if err != nil {
			t.Fatalf("Error creating test file 2: %v", err)
		}
		defer os.Remove(file2)

		err = os.WriteFile(file3, []byte(content3), 0644)
		if err != nil {
			t.Fatalf("Error creating test file 3: %v", err)
		}
		defer os.Remove(file3)

		var out bytes.Buffer
		err = ConcatenateFiles(&out, file1, file2, file3)
		if err != nil {
			t.Errorf("ConcatenateFiles failed: %v", err)
		}
		assertConcatenatedOutput(t, &out, content1+content2+content3)
	})

	// Add more test cases as needed...
}
