package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/user"
	"strings"

	"github.com/HO4678/CSCE4600/Project2/builtins"
)

func main() {
	exit := make(chan struct{}, 2) // buffer this so there's no deadlock.
	runLoop(os.Stdin, os.Stdout, os.Stderr, exit)
}

func runLoop(r io.Reader, w, errW io.Writer, exit chan struct{}) {
	var (
		input    string
		err      error
		readLoop = bufio.NewReader(r)
	)
	for {
		select {
		case <-exit:
			_, _ = fmt.Fprintln(w, "exiting gracefully...")
			return
		default:
			if err := printPrompt(w); err != nil {
				_, _ = fmt.Fprintln(errW, err)
				continue
			}
			if input, err = readLoop.ReadString('\n'); err != nil {
				_, _ = fmt.Fprintln(errW, err)
				continue
			}
			if err = handleInput(w, input, exit); err != nil {
				_, _ = fmt.Fprintln(errW, err)
			}
		}
	}
}

func printPrompt(w io.Writer) error {
	u, err := user.Current()
	if err != nil {
		return err
	}
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(w, "%v [%v] $ ", wd, u.Username)

	return err
}

func handleInput(w io.Writer, input string, exit chan<- struct{}) error {
	input = strings.TrimSpace(input)
	args := strings.Split(input, " ")
	name, args := args[0], args[1:]

	switch name {
	case "cd":
		return builtins.ChangeDirectory(args...)
	case "env":
		return builtins.EnvironmentVariables(w, args...)
	case "exit":
		exit <- struct{}{}
		return nil
	case "echo":
		builtins.Echo(w, args...)
	case "pwd":
		builtins.PrintWorkingDirectory(w)
	case "mkdir":
		return builtins.MakeDirectory(args...)
	case "rmdir":
		return builtins.RemoveDirectory(args...)
	case "touch":
		return builtins.TouchFile(args...)
	case "ls":
		return builtins.ListDirectory(w, args...)
	case "cat":
		return builtins.ConcatenateFiles(w, args...)
	case "grep":
		return builtins.Grep(w, args...)
	}

	return executeCommand(name, args...)
}

func executeCommand(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
