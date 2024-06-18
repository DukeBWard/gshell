package main

import (
	"bufio"
	"context"
	"fmt"
	"gshell/src/cmd"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

// Cannot use := at the package level
var COMMANDS = map[string]func(dir string, args ...string) (new_dir string){
	"ls":  cmd.Ls,
	"pwd": cmd.Pwd,
	"cd":  cmd.Cd,
}

func main() {
	var wg sync.WaitGroup

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)

	cursor := bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))
	curr_directory, err := os.Getwd()
	if err != nil {
		fmt.Println("Error with startup:", err)
		return
	}

	for {
		ctx, cancel := context.WithCancel(context.Background())
		cursor.WriteString("gshell> ")
		// WriteString writes to a buffer so you need to flush it to display it
		cursor.Flush()

		input, err := cursor.ReadString('\n')
		if err != nil {
			fmt.Println("Error with input:", err)
			return
		}

		input = strings.TrimSpace(input)
		input_slice := strings.Split(input, " ")
		command, exists := COMMANDS[input_slice[0]]

		if input == "exit" {
			break
		}

		wg.Add(1)
		if exists {
			go func() {
				defer wg.Done()
				curr_directory = command(curr_directory, input_slice[1:]...)
			}()
		} else if input_slice[0] == "" {
			wg.Done()
			continue
		} else {
			go func() {
				defer wg.Done()
				cmd.Run_external(ctx, curr_directory, input_slice...)
			}()
		}

		select {
		case <-sigChan:
			fmt.Println("\nCommand interrupted")
			cancel()
		default:
			wg.Wait()
		}

	}
}
