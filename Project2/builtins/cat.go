// builtins/cat.go

package builtins

import (
	"fmt"
	"io"
	"os"
)

// ConcatenateFiles concatenates and prints the content of one or more files.
func ConcatenateFiles(w io.Writer, files ...string) error {
	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read file %s: %w", file, err)
		}
		fmt.Fprint(w, string(content))
	}
	return nil
}
