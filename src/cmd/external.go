package cmd

import (
	"context"
	"fmt"
	"os/exec"
)

func Run_external(ctx context.Context, dir string, args ...string) {
	// fmt.Printf("Running command: %s with args: %v\n", cmd_str, args)
	cmd := exec.CommandContext(ctx, "powershell", args...)
	if val, err := cmd.CombinedOutput(); err != nil {
		fmt.Println("Could not run external command: ", err)
	} else {
		// time.Sleep(3 * time.Second)
		// fmt.Println("we made it" + string(val))
		fmt.Println(string(val))
	}
}
