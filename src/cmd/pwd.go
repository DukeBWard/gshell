package cmd

import (
	"fmt"
	"os"
)

// Don't forget in Go first letter needs to be capitalized for public
func Pwd(args ...string) {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Could not get pwd: ", err)
	}

	fmt.Println(pwd)
}
