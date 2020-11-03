/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/quadrabee/maak/config"
	"github.com/quadrabee/maak/make"
	"github.com/spf13/cobra"
)

func find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// iniCmd represents the init command
var iniCmd = &cobra.Command{
	Use:       "init",
	ValidArgs: []string{"all", "maak", "makefile"},
	Short:     "Initialize maak for your project",
	Long:      `This will generate your first maak.yaml config and generate a couple of Makefile templates`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) <= 0 {
			args = append(args, "all")
		}

		_, genMaakConfig := find(args, "maak")
		_, genMakefile := find(args, "makefile")
		_, genAll := find(args, "all")

		if genAll || genMaakConfig {
			err := config.EnsureNotExists()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			err = config.Generate()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		if genAll || genMakefile {
			err := make.Generate()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(iniCmd)
}
