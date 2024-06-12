package cmd

import (
	"fmt"
	"os"
)

func Cd(dir string, args ...string) (new_dir string) {

	// error checking
	if len(args) < 1 {
		fmt.Println("cd usage: cd <directory>")
		return dir
	}

	if args[0] == ".." {
		err := os.Chdir("..")
		if err != nil {
			fmt.Println("Could not change directory: ", err)
		}
	} else {
		err := os.Chdir(dir + "\\" + args[0])
		if err != nil {
			fmt.Println("Could not change directory: ", err)
		}
	}

	new_dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Could not change directory: ", err)
	}

	fmt.Println(new_dir)
	return new_dir
}
