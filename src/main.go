package main

import (
	"bufio"
	"fmt"
	"gshell/src/cmd"
	"os"
	"strings"
)

// Cannot use := at the package level

var COMMANDS = map[string]func(dir string, args ...string){
	"ls":  cmd.Ls,
	"pwd": cmd.Pwd,
	"cd":  cmd.Cd,
}

func main() {
	exit := false
	cursor := bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))
	curr_directory, err := os.Getwd()
	if err != nil {
		fmt.Println("Error with startup: ", err)
	}

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
		input_slice := strings.Split(input, " ")

		if input == "exit" {
			exit = true
			return
		}

		// Silly go formatting, you need else to be on the same line as end bracket of if
		if exists {
			//fmt.Println(input_slice[1:])
			command(curr_directory, input_slice[1:]...)
		} else {
			cmd.Run_external(curr_directory, input_slice...)
			//fmt.Println("Command does not exist.")
		}

	}
}
