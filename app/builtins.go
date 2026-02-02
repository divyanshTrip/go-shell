package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type BuiltinFunc func(s *Shell, args []string, stdout io.Writer) error

func HandleEcho(s *Shell, args []string, stdout io.Writer) error {
	_, err := fmt.Fprintf(stdout, strings.Join((args), " "))
	fmt.Println()
	return err
}

func HandleExit(s *Shell, args []string, stdout io.Writer) error {
	os.Exit(0)
	return nil
}

func HandleType(s *Shell, args []string, stdout io.Writer) error {
	if len(args) == 0 {
		return nil
	}
	//name of the command to be checked the type of
	cmdName := args[0]
	if _, ok := s.Builtins[cmdName]; ok {
		fmt.Fprintf(stdout, "%s is a shell builtin\n", cmdName)
		return nil
	}
	path, ok := searchPath(s.Path, cmdName)
	if ok {
		fmt.Fprintf(stdout, "%s is %s\n", cmdName, path)
		return nil
	}
	_, err := fmt.Fprintf(stdout, "%s: not found\n", cmdName)
	return err
}
