package cmd

import (
	"fmt"
	"os"
)

// Don't forget in Go first letter needs to be capitalized for public
func Ls(dir string, args ...string) (new_dir string) {
	entries, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	maxLength := 0
	for _, entry := range entries {
		if len(entry.Name()) > maxLength {
			maxLength = len(entry.Name())
		}
	}

	for _, entry := range entries {
		fmt.Printf("%-*s", maxLength+2, entry.Name())
	}

	fmt.Println()
	return dir
}
