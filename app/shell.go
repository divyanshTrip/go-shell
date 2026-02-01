package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Shell struct {
	Builtins map[string]BuiltinFunc
	Path     string
}

func NewShell() *Shell {

	s := &Shell{
		Path: os.Getenv("PATH"),
	}

	s.Builtins = map[string]BuiltinFunc{
		"echo": HandleEcho,
		"exit": HandleExit,
		"type": HandleType,
	}

	return s
}

func (s *Shell) Run() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprintf(os.Stdout, "$ ")
		input, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Fprintln(os.Stderr, "Error reading input :", err)
			continue
		}
		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}
		parts := strings.Fields(input)
		cmdName := parts[0]
		args := parts[1:]
		if fn, ok := s.Builtins[cmdName]; ok {
			err := fn(s, args, os.Stdout)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error :", err)
			}
		}
		continue
	}

}
