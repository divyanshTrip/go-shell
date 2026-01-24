package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var commands = map[string]bool{
	"type": true,
	"echo": true,
	"exit": true,
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "error reading input:", err)
			break
		}
		command = strings.TrimSpace(command)
		parts := strings.Fields(command)
		if len(parts) == 0 {
			continue
		}
		cmd := parts[0]

		if cmd == "exit" {
			break
		}

		if cmd == "echo" {
			fmt.Println(strings.Join(parts[1:], " "))
			continue
		}
		if cmd == "type" {
			if len(parts) == 1 {
				fmt.Println()
			} else {
				value, ok := commands[parts[1]]
				if ok == true {
					if value == true {
						fmt.Println(parts[1], "is a shell builtin")
						continue
					} else {
						fmt.Println(parts[1], "is not a shell builtin")
						continue
					}
				} else {
					fmt.Println(parts[1] + ": not found")
					continue
				}
			}
		}
		fmt.Println(cmd + ": not found")
	}
}
