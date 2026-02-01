package main

import (
	"os"
	"path/filepath"
	"strings"
)

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
