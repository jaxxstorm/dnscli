// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile    string
	recordType string
	ttl        int
	address    string
	record     string
	dnsserver  string
	dnsport    string
	dnskey     string
	dnskeyalgo string
	dnskeyname string
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "dnscli",
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

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	//RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.dnscli.yaml)")
	RootCmd.PersistentFlags().StringVarP(&recordType, "type", "t", "A", "DNS Record Type")
	RootCmd.PersistentFlags().IntVar(&ttl, "ttl", 8600, "Record TTL")
	RootCmd.PersistentFlags().StringVarP(&address, "address", "a", "", "Destination address for record")
	RootCmd.PersistentFlags().StringVarP(&record, "record", "r", "", "Fully qualified record name")
	RootCmd.PersistentFlags().StringVarP(&dnsserver, "server", "s", "127.0.0.1", "Address of DNS server")
	RootCmd.PersistentFlags().StringVarP(&dnsport, "port", "p", "53", "Port DNS server is listening on")
	RootCmd.PersistentFlags().StringVarP(&dnskey, "key", "k", "", "DNS Key to use")
	RootCmd.PersistentFlags().StringVar(&dnskeyalgo, "algo", "", "DNS Key Algo")
	RootCmd.PersistentFlags().StringVar(&dnskeyname, "key-name", "", "DNS Key name")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".dnscli") // name of config file (without extension)
	viper.AddConfigPath("$HOME")   // adding home directory as first search path
	viper.AutomaticEnv()           // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
