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

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use: "build",
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		_, err := config.Load()
		if err != nil {
			return []string{}, cobra.ShellCompDirectiveNoFileComp
		}
		comps := config.ComponentNames()
		return comps, cobra.ShellCompDirectiveNoFileComp
	},
	Short: "Build a component",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		config.EnsureLoaded()

		buildAll, _ := cmd.Flags().GetBool("all")
		buildForce, _ := cmd.Flags().GetBool("force")

		if !buildAll && len(args) <= 0 {
			fmt.Println("Component name expected, or use --all to build all components")
			os.Exit(1)
		}

		var components []string
		if buildAll {
			components = config.ComponentNames()
		} else {
			components = args
		}

		err := make.BuildComponents(components, buildForce)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

	buildCmd.PersistentFlags().BoolP("all", "a", false, "Build all components")
	buildCmd.PersistentFlags().BoolP("force", "f", false, "Force the rebuild")
}
