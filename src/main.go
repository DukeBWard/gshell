package main

import (
	"bufio"
	"fmt"
	"gshell/src/cmd"
	"os"
	"strings"
)

// Cannot use := at the package level

var COMMANDS = map[string]func(){
	"ls":  cmd.Ls,
	"pwd": cmd.Pwd,
}

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

		command, exists := COMMANDS[input]

		if input == "exit" {
			exit = true
			return
		}

		// Silly go formatting, you need else to be on the same line as end bracket of if
		if exists {
			command()
		} else {
			fmt.Println("Command does not exist.")
		}

	}
}
