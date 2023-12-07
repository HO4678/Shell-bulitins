// builtins/mkdir.go

package builtins

import (
	"fmt"
	"os"
)

// MakeDirectory creates a new directory with the given name.
func MakeDirectory(args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Usage: mkdir <directory>")
	}

	dirPath := args[0]

	// Create the directory with read, write, and execute permissions for the owner.
	err := os.Mkdir(dirPath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("Error creating directory %s: %w", dirPath, err)
	}

	return nil
}
