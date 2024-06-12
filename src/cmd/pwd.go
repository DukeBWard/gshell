package cmd

import (
	"fmt"
)

// Don't forget in Go first letter needs to be capitalized for public
func Pwd(dir string, args ...string) (new_dir string) {
	// pwd, err := os.Getwd()
	// if err != nil {
	// 	fmt.Println("Could not get pwd: ", err)
	// }

	pwd := dir

	fmt.Println(pwd)
	return dir
}
