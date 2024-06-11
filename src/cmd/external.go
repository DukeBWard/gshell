package cmd

import (
	"fmt"
	"os/exec"
)

func Run_external(args ...string) {
	// fmt.Printf("Running command: %s with args: %v\n", cmd_str, args)
	cmd := exec.Command("powershell", args...)
	if val, err := cmd.CombinedOutput(); err != nil {
		fmt.Println("Could not run external command: ", err)
	} else {
		fmt.Println(string(val))
	}
}
