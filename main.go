package main

import (
	"bufio"
	"fmt"
	"os"
)

// Cannot use := at the package level

func main() {
	exit := false
	cusor := bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))

	input, err := cusor.ReadString('\n')
	if err != nil {
		fmt.Println("Error with input:", err)
		return
	}

	for !exit {

	}

}
