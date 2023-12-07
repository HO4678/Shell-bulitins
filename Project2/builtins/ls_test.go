// builtins/ls_test.go

package builtins

import (
	"bytes"
	"os"
	"testing"
)

func assertListContains(t *testing.T, out *bytes.Buffer, expected string) {
	t.Helper()
	if !bytes.Contains(out.Bytes(), []byte(expected)) {
		t.Errorf("ListDirectory output does not contain expected entry: %s", expected)
	}
}

func TestListDirectory(t *testing.T) {
	t.Run("ListContentsOfTestDirectory", func(t *testing.T) {
		testDir := "test_directory"
		defer os.RemoveAll(testDir)

		err := os.Mkdir(testDir, os.ModePerm)
		if err != nil {
			t.Fatalf("Failed to create test directory: %v", err)
		}

		testFiles := []string{"file1.txt", "file2.txt", "file3.txt"}
		for _, file := range testFiles {
			filePath := testDir + "/" + file
			f, err := os.Create(filePath)
			if err != nil {
				t.Fatalf("Failed to create test file %s: %v", filePath, err)
			}
			defer f.Close()
		}

		var out bytes.Buffer
		err = ListDirectory(&out, testDir)
		if err != nil {
			t.Errorf("ListDirectory failed: %v", err)
		}

		for _, file := range testFiles {
			assertListContains(t, &out, file+"\n")
		}
	})

	t.Run("ListContentsOfCurrentDirectory", func(t *testing.T) {
		var out bytes.Buffer
		err := ListDirectory(&out)
		if err != nil {
			t.Errorf("ListDirectory failed: %v", err)
		}

		// Add assertions for the current directory as needed
	})
}
