/*
Copyright © 2021 Bren 'fraq' Briggs (code@fraq.io)

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
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "flexo",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.flexo.yaml)")

	// All your args are belong to Viper
	rootCmd.PersistentFlags().StringVarP(&dbUser, "dbUser", "", dbUser, "database username")
	err := viper.BindPFlag("dbUser", rootCmd.PersistentFlags().Lookup("dbUser"))
	if err != nil {
		fmt.Printf("Couldn't get flag: %s\n", err)
	}

	rootCmd.PersistentFlags().StringVarP(&dbPass, "dbPass", "", dbPass, "database password")
	err = viper.BindPFlag("dbPass", rootCmd.PersistentFlags().Lookup("dbPass"))
	if err != nil {
		fmt.Printf("Couldn't get flag: %s\n", err)
	}

	rootCmd.PersistentFlags().StringVarP(&dbAddr, "dbAddr", "", dbAddr, "database address")
	err = viper.BindPFlag("dbAddr", rootCmd.PersistentFlags().Lookup("dbAddr"))
	if err != nil {
		fmt.Printf("Couldn't get flag: %s\n", err)
	}

	rootCmd.PersistentFlags().StringVarP(&dbName, "dbName", "", dbName, "database to use")
	err = viper.BindPFlag("dbName", rootCmd.PersistentFlags().Lookup("dbName"))
	if err != nil {
		fmt.Printf("Couldn't get flag: %s\n", err)
	}

	rootCmd.PersistentFlags().StringVarP(&dbSSLMode, "dbSSL", "", dbSSLMode, "use sslmode to connect to the database")
	err = viper.BindPFlag("dbSSL", rootCmd.PersistentFlags().Lookup("dbSSL"))
	if err != nil {
		fmt.Printf("Couldn't get flag: %s\n", err)
	}

	rootCmd.PersistentFlags().StringVarP(&secret, "secret", "", secret, "secret shared with the front-end")
	err = viper.BindPFlag("secret", rootCmd.PersistentFlags().Lookup("secret"))
	if err != nil {
		fmt.Printf("Couldn't get flag: %s\n", err)
	}

	rootCmd.PersistentFlags().StringVarP(&BytebotRedis, "bytebotRedis", "", BytebotRedis, "Address of a redis db used by bytebot to transmit events to discord")
	err = viper.BindPFlag("bytebotRedis", rootCmd.PersistentFlags().Lookup("bytebotRedis"))
	if err != nil {
		fmt.Printf("Couldn't get flag: %s\n", err)
	}

	rootCmd.PersistentFlags().StringVarP(&BytebotPubSub, "bytebotPubSub", "", BytebotPubSub, "Address of the bytebot channel used for outbound discord messages")
	err = viper.BindPFlag("bytebotPubSub", rootCmd.PersistentFlags().Lookup("bytebotPubSub"))
	if err != nil {
		fmt.Printf("Couldn't get flag: %s\n", err)
	}

	rootCmd.PersistentFlags().StringVarP(&BytebotChannel, "bytebotChannel", "", BytebotChannel, "ID of the discord channel messages should be pushed to")
	err = viper.BindPFlag("bytebotChannel", rootCmd.PersistentFlags().Lookup("bytebotChannel"))
	if err != nil {
		fmt.Printf("Couldn't get flag: %s\n", err)
	}

	viper.SetDefault("dbUser", "root")
	viper.SetDefault("dbAddr", "127.0.0.1:5432")
	viper.SetDefault("dbName", "flexo")
	viper.SetDefault("dbSSL", "disable")
	viper.SetDefault("secret", "shared_secret")
	viper.SetDefault("bytebotRedis", "redis:6379")
	viper.SetDefault("bytebotPubSub", "discord-outbound")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".flexo" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".flexo")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
