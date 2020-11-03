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
	"github.com/spf13/cobra"
)

// depsCmd represents the deps command
var depsCmd = &cobra.Command{
	Use: "deps",
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		_, err := config.Load()
		if err != nil {
			return []string{}, cobra.ShellCompDirectiveNoFileComp
		}
		conts := config.ContainerNames()
		return conts, cobra.ShellCompDirectiveNoFileComp
	},
	Short: "Lists the dependencies of a container",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		config.EnsureLoaded()

		if len(args) > 1 {
			fmt.Println("Only one container can be provided")
			os.Exit(1)
		}

		if len(args) < 1 {
			fmt.Println("Please specify the name of a container")
			os.Exit(1)
		}

		container := args[0]
		cont := config.GetContainer(container)

		if cont == nil {
			fmt.Println("Unknown container:", container)
			os.Exit(1)
		}

		deps, err := cont.GlobDependencies()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for _, file := range deps {
			fmt.Println(file)
		}
	},
}

func init() {
	rootCmd.AddCommand(depsCmd)
}
