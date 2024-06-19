package history

import (
	"fmt"
	"os"
)

const HIST_NAME = ".gshell_history"

func Append_history(cmd string) {

	data := []byte(cmd)

	err := os.WriteFile(HIST_NAME, data, 0777)
	if err != nil {
		fmt.Println(err)
	}

	// data, err = os.ReadFile(HIST_NAME)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(string(data))
}
