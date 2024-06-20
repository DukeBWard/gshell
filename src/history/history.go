package history

import (
	"fmt"
	"os"
)

const HIST_NAME = ".gshell_history"

func Append_history(cmd string) error {

	// check for file
	_, err := os.Stat(HIST_NAME)
	if os.IsNotExist(err) {
		file, err := os.Create(HIST_NAME)
		if err != nil {
			return fmt.Errorf("cannot create file: %w", err)

		}
		defer file.Close()
	}

	file, err := os.OpenFile(HIST_NAME, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("cannot open file: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(cmd)
	if err != nil {
		return fmt.Errorf("cannot write to file: %w", err)
	}

	return nil

	// data := []byte(cmd)

	// err := os.WriteFile(HIST_NAME, data, 0777)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// data, err = os.ReadFile(HIST_NAME)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(string(data))
}
