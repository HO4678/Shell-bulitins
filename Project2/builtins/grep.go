// builtins/grep.go

package builtins

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Grep searches for a pattern in a file and prints matching lines.
func Grep(w io.Writer, args ...string) error {
	if len(args) < 2 {
		return fmt.Errorf("Usage: grep <pattern> <file>")
	}

	pattern := args[0]
	file := args[1]

	f, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", file, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, pattern) {
			fmt.Fprintln(w, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file %s: %w", file, err)
	}

	return nil
}
