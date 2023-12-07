// builtins/ls.go

package builtins

import (
	"fmt"
	"io"
	"os"
)

func ListDirectory(w io.Writer, args ...string) error {
	// If no arguments provided, list the current directory
	if len(args) == 0 {
		args = append(args, ".")
	}

	for _, dir := range args {
		files, err := os.ReadDir(dir)
		if err != nil {
			return fmt.Errorf("Error reading directory %s: %v", dir, err)
		}

		// Print the list of files and directories
		for _, file := range files {
			fmt.Fprintln(w, file.Name())
		}
	}

	return nil
}
