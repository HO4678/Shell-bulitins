// builtins/pwd.go

package builtins

import (
	"fmt"
	"io"
	"os"
)

// PrintWorkingDirectory prints the current working directory to the provided writer.
func PrintWorkingDirectory(w io.Writer) {
	// Get the current working directory.
	dir, err := os.Getwd()
	if err != nil {
		// Print the error if unable to get the working directory.
		fmt.Fprintln(w, "Error:", err)
		return
	}

	// Print the current working directory.
	fmt.Fprintln(w, dir)
}
