/*
Copyright © 2025 sottey

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// tasksCmd represents the tasks command
var tasksCmd = &cobra.Command{
	Use: "tasks [YYYMMDD]",
	Short: `Display tasks from a specific date specified in the format YYYMMDD 
	        (if no date provided, today's note is used)`,
	Example: "goteplan tasks 20250206",
	Run: func(cmd *cobra.Command, args []string) {
		var dateFile string

		// If date provided, use that, else use today's date
		if len(args) == 0 {
			year, month, day := time.Now().Date()
			dateFile = fmt.Sprintf("Calendar/%04d%02d%02d.md", year, int(month), day)
		} else {
			dateFile = fmt.Sprintf("Calendar/%v.md", args[0])
		}

		fmt.Printf("Displaying tasks for %v\n\n", dateFile)

		path := filepath.Join(BaseDir, dateFile)
		content, err := os.ReadFile(path)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		tasks := getTasks(string(content[:]))

		fmt.Println(tasks)
	},
}

func init() {
	tasksCmd.GroupID = "main"
	rootCmd.AddCommand(tasksCmd)
}

func getTasks(note string) string {
	var ret string
	noteLines := strings.Split(note, "\n")

	for _, currLine := range noteLines {
		currLine = strings.TrimSpace(currLine)
		if strings.HasPrefix(currLine, TodoSymbol+" ") {
			task := strings.Replace(currLine, TodoSymbol+" ", "", -1)
			if !strings.HasPrefix(task, "[x]") {
				task = "[ ] " + task
			}

			ret += task + "\n"
		}
	}

	if len(ret) == 0 {
		ret = "No tasks found"
	}

	return ret
}
