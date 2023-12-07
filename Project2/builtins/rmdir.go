// builtins/rmdir.go

package builtins

import (
	"fmt"
	"os"
)

// RemoveDirectory removes the directory with the given name.
func RemoveDirectory(args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Usage: rmdir <directory>")
	}

	dir := args[0]

	// Check if the directory exists before attempting to remove it.
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return fmt.Errorf("Directory %s does not exist", dir)
	} else if err != nil {
		return fmt.Errorf("Error checking directory: %v", err)
	}

	// Remove the directory.
	err = os.Remove(dir)
	if err != nil {
		return fmt.Errorf("Error removing directory %s: %v", dir, err)
	}

	return nil
}
