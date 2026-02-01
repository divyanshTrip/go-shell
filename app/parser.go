package main

import (
	"strings"
)

type Command struct {
	Name string
	Args []string
}

func parse(line string) Command {

	parts := strings.Fields(line)
	if len(parts) == 0 {
		return Command{}
	}

	return Command{
		Name: parts[0],
		Args: parts[1:],
	}
}
