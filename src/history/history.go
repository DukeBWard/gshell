package history

import (
	"bufio"
	"fmt"
	"os"

	"github.com/rivo/tview"
)

var HIST_NAME = os.Getenv("userprofile") + "\\" + ".gshell_history"

func List_history(dir string, args ...string) (new_dir string) {

	file, err := os.Open(HIST_NAME)
	if err != nil {
		fmt.Errorf("Cannot open file: %w", err)
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Errorf("Scanner had an error: %w", err)
	}

	app := tview.NewApplication()
	textView := tview.NewTextView().SetDynamicColors(true).SetRegions(true).SetScrollable(true)

	for _, line := range lines {
		fmt.Fprintln(textView, line)
	}

	textView.SetChangedFunc(func() {
		app.Draw()
	})

	if err := app.SetRoot(textView, true).Run(); err != nil {
		fmt.Errorf("Error running application: %v", err)
	}

	return dir
}

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
}
