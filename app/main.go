package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
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
			_, ok := commands[parts[1]]
			if ok == true {
				fmt.Println(parts[1], "is a shell builtin")
				continue
			} else {
				path, found := searchPath(os.Getenv("PATH"), parts[1])
				if found == true {
					fmt.Println(parts[1] + " is " + path)
					continue
				} else {
					fmt.Println(parts[1] + ": not found")
					continue
				}
			}
		} else {
			fmt.Println(parts[0] + ": not found")
		}
	}
}

func searchPath(path string, cmd string) (string, bool) {
	dirs := strings.Split(path, string(os.PathListSeparator))
	for _, dir := range dirs {

		if dir == "" {
			dir = "."
		}

		candidate := filepath.Join(dir, cmd)

		info, err := os.Stat(candidate)
		if err != nil {
			continue
		}
		if info.IsDir() {
			continue
		}
		if info.Mode()&0111 != 0 {
			return candidate, true
		}

	}
	return "", false
}
