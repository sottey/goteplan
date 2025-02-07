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
	"github.com/spf13/viper"
)

var BaseDir string
var RenderMarkdown bool

var rootCmd = &cobra.Command{
	Use:   "goteplan",
	Short: "A CLI for NotePlan notes management",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	var mainGroup cobra.Group
	mainGroup.Title = "Commands:"
	mainGroup.ID = "main"
	rootCmd.AddGroup(&mainGroup)
	SetupViper()
}

func SetupViper() {
	dataPath, _ := expandPath("~/Library/Containers/co.noteplan.NotePlan-setapp/Data/Library/Application Support/co.noteplan.NotePlan-setapp")
	configPath, _ := expandPath("~/")
	rootCmd.PersistentFlags().StringVarP(&BaseDir, "basedir", "b", "", "Root location of the NotePlan data")
	rootCmd.PersistentFlags().BoolVarP(&RenderMarkdown, "render", "r", false, "If present, display will attempt to render markdown. If not, source will be shown.")
	viper.BindPFlag("basedir", rootCmd.PersistentFlags().Lookup("basedir"))
	viper.BindPFlag("render", rootCmd.PersistentFlags().Lookup("render"))

	viper.SetConfigName(".goteplan")
	viper.SetConfigType("json")
	viper.AddConfigPath(configPath)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Config file not found. Creating...")
			viper.SetDefault("basedir", dataPath)
			if errTwo := viper.SafeWriteConfig(); errTwo != nil {
				fmt.Printf("Error creating config file: '%v'\n", errTwo)
			}
		} else {
			fmt.Printf("Error opening config file: %v\n", err)
			return
		}
	}

	BaseDir = viper.GetString("basedir")
	RenderMarkdown = viper.GetBool("render")
}

func expandPath(path string) (string, error) {
	if strings.HasPrefix(path, "~") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(homeDir, path[1:]), nil
	}
	return filepath.Abs(path)
}
