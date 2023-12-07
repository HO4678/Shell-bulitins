// echo.go
package builtins

import (
	"fmt"
	"io"
	"strings"
)

// Echo writes the provided arguments to the specified writer.
func Echo(w io.Writer, args ...string) {
	fmt.Fprintln(w, strings.Join(args, " "))
}
