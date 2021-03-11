/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/SECCDC/flexo/flexo"
)

var (
	dbUser    string
	dbPass    string
	dbAddr    string
	dbName    string
	dbSSLMode string
	secret    string
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start flexo webserver",
	Run: func(cmd *cobra.Command, args []string) {
		if viper.GetString("dbPass") == "" {
			fmt.Println("Password not provided. Exiting!")
			os.Exit(2)
		}
		c := flexo.Config{
			DBUser: viper.GetString("dbUser"),
			DBPass: viper.GetString("dbPass"),
			DBAddr: viper.GetString("dbAddr"),
			DBName: viper.GetString("dbName"),
			DBssl:  viper.GetString("dbSSL"),
			Secret: viper.GetString("secret"),
		}
		flexo.Run(c)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
