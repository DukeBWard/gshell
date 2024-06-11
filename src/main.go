package main

import (
	"bufio"
	"fmt"
	"gshell/src/cmd"
	"os"
	"strings"
)

// Cannot use := at the package level

func main() {
	exit := false
	cursor := bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))

	for !exit {
		cursor.WriteString("gshell> ")
		// WriteString writes to a buffer so you need to flush it to display it
		cursor.Flush()

		input, err := cursor.ReadString('\n')
		if err != nil {
			fmt.Println("Error with input:", err)
			return
		}

		input = strings.TrimSpace(input)

		if input == "exit" {
			exit = true
		}

		if input == "ls" {
			cmd.Ls()
		}

	}
}
