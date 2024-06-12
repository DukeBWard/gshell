package cmd

import (
	"fmt"
	"os"
)

func Cd(dir string, args ...string) {
	fmt.Println(args)

	if len(args) < 2 {
		fmt.Println("cd usage: cd <directory>")
		return
	}

	curr_directory, err := os.Getwd()
	if err != nil {
		fmt.Println("Could not change directory: ", err)
	}

	err = os.Chdir(curr_directory + "/" + args[1])
	if err != nil {
		fmt.Println("Could not change directory: ", err)
	}

	curr_directory, err = os.Getwd()
	if err != nil {
		fmt.Println("Could not change directory: ", err)
	}
	fmt.Println(curr_directory)
	return
}
