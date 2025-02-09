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

	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:     "search <query>",
	Short:   "Search for notes with the specified string - Example: goteplan search MikeS",
	Args:    cobra.MinimumNArgs(1),
	Example: "goteplan search QueryString \t# Search for QueryString using exact case\ngoteplan search -n QueryString \t# Search for querystring without case sensitivity",
	Run: func(cmd *cobra.Command, args []string) {
		query := args[0]
		counter := 0

		err := filepath.Walk(BaseDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.HasSuffix(info.Name(), ".md") {
				content, err := os.ReadFile(path)
				if err != nil {
					return nil
				}

				if NoCaseSearch {
					content = []byte(strings.ToLower(string(content)))
					query = strings.ToLower(query)
				}

				if strings.Contains(string(content), query) {
					fmt.Println("Match : ", strings.TrimPrefix(path, BaseDir+"/"))
					counter++
				}
			}
			return nil
		})

		if err != nil {
			fmt.Println("Error searching notes:", err)
		} else {
			fmt.Printf("%v notes found for query '%v' (Case Inensitive: %v)\n\n", counter, query, NoCaseSearch)
		}
	},
}

func init() {
	searchCmd.GroupID = "main"
	rootCmd.AddCommand(searchCmd)
}
