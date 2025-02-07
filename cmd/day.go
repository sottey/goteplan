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
	"time"

	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/spf13/cobra"
)

// dayCmd represents the day command
var dayCmd = &cobra.Command{
	Use: "day [YYYMMDD]",
	Short: `Display note from a specific date specified in the format YYYMMDD 
	    (if no date provided, today's note is shown)`,
	// Args:    cobra.MinimumNArgs(1),
	Example: "goteplan day 20250206",
	Run: func(cmd *cobra.Command, args []string) {
		var dateFile string

		// If date provided, use that, else use today's date
		if len(args) == 0 {
			year, month, day := time.Now().Date()
			dateFile = fmt.Sprintf("Calendar/%04d%02d%02d.md", year, int(month), day)
		} else {
			dateFile = fmt.Sprintf("Calendar/%v.md", args[0])
		}

		fmt.Printf("Displaying %v\n\n", dateFile)

		path := filepath.Join(BaseDir, dateFile)
		content, err := os.ReadFile(path)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		if RenderMarkdown {
			content = markdown.Render(string(content[:]), 80, 6)
		}

		fmt.Println(string(content[:]))
	},
}

func init() {
	dayCmd.GroupID = "main"
	rootCmd.AddCommand(dayCmd)
}
