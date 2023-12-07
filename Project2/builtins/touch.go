// builtins/touch.go

package builtins

import (
	"fmt"
	"os"
)

// TouchFile creates an empty file with the given name.
func TouchFile(args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Usage: touch <file>")
	}

	filePath := args[0]

	// Open the file or create it if it doesn't exist.
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return fmt.Errorf("Error creating or updating file %s: %w", filePath, err)
	}
	defer file.Close()

	return nil
}
