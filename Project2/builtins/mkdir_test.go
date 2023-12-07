// builtins/mkdir_test.go

package builtins

import (
	"os"
	"testing"
)

func TestMakeDirectory(t *testing.T) {
	t.Run("CreateNewDirectory", func(t *testing.T) {
		testDir := "test_directory"
		defer os.RemoveAll(testDir)

		if err := MakeDirectory(testDir); err != nil {
			t.Errorf("Failed to create directory: %v", err)
		}

		stat, err := os.Stat(testDir)
		if err != nil {
			t.Errorf("Error checking directory: %v", err)
		}

		if !stat.IsDir() {
			t.Errorf("Expected %s to be a directory", testDir)
		}
	})

	t.Run("CreateExistingDirectory", func(t *testing.T) {
		existingDir := "existing_directory"
		defer os.RemoveAll(existingDir)

		if err := MakeDirectory(existingDir); err != nil {
			t.Errorf("Failed to create directory: %v", err)
		}

		// Attempt to create the same directory again
		if err := MakeDirectory(existingDir); err == nil {
			t.Errorf("Expected an error for existing directory, but got none")
		}
	})

	t.Run("CreateInvalidPath", func(t *testing.T) {
		invalidPath := "/invalid/path/with/special/characters"
		err := MakeDirectory(invalidPath)
		if err == nil {
			t.Errorf("Expected an error for an invalid path, but got none")
		}
	})

	// Additional test cases ...

	t.Run("CreateChildDirectoryNonExistentParent", func(t *testing.T) {
		parentDir := "parent_directory"
		childDir := "parent_directory/child_directory"
		defer os.RemoveAll(parentDir)

		err := MakeDirectory(childDir)
		if err == nil {
			t.Errorf("Expected an error for non-existent parent directory, but got none")
		}
	})

	t.Run("CreateDirectoryReadOnlyParent", func(t *testing.T) {
		readOnlyDir := "read_only_directory"
		childDir := readOnlyDir + "/child_directory"
		defer os.RemoveAll(readOnlyDir)
		os.Mkdir(readOnlyDir, 0444) // Read-only permission

		err := MakeDirectory(childDir)
		if err == nil {
			t.Errorf("Expected an error for read-only parent directory, but got none")
		}
	})
}
