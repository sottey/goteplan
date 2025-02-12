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
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:     "create <filename>",
	Short:   "Create a new note - Example: goteplan create Notes/Home/NewNote.md",
	Args:    cobra.MinimumNArgs(1),
	Example: "goteplan create Notes/Home/mynote.md",
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		path := filepath.Join(BaseDir, filename)
		fmt.Println("Enter your note content. Press Ctrl+D when done:")
		scanner := bufio.NewScanner(os.Stdin)
		var content string
		for scanner.Scan() {
			content += scanner.Text() + "\n"
		}
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			fmt.Println("Error creating file:", err)
		}
		fmt.Printf("Note saved to '%v'\n", path)
	},
}

func init() {
	createCmd.GroupID = "main"
	rootCmd.AddCommand(createCmd)
}
