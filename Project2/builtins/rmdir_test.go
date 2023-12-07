// builtins/rmdir_test.go

package builtins

import (
	"os"
	"testing"
)

func TestRemoveDirectory(t *testing.T) {
	// Define a test directory name
	testDir := "test_directory"

	// Clean up in case the directory already exists
	defer os.RemoveAll(testDir)

	// Create the directory for testing
	err := os.Mkdir(testDir, os.ModePerm)
	if err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}

	tests := []struct {
		name     string
		args     []string
		wantErr  bool
		expected string
	}{
		{
			name:     "Success",
			args:     []string{testDir},
			wantErr:  false,
			expected: "",
		},
		{
			name:     "DirectoryDoesNotExist",
			args:     []string{"nonexistent_directory"},
			wantErr:  true,
			expected: "Directory nonexistent_directory does not exist",
		},
		{
			name:     "MissingArgument",
			args:     []string{},
			wantErr:  true,
			expected: "Usage: rmdir <directory>",
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := RemoveDirectory(tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveDirectory() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr && err.Error() != tt.expected {
				t.Errorf("RemoveDirectory() got = %v, want %v", err.Error(), tt.expected)
			}
		})
	}
}
