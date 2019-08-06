// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"strings"

	defaults "github.com/mcuadros/go-defaults"
	"github.com/scraly/hello-world/cli/hello-world/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.zenithar.org/pkg/flags"
	"go.zenithar.org/pkg/log"
)

var (
	cfgFile string
	conf    = &config.Configuration{}
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hello-world",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.hello-world.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(clientCmd)
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(sayCmd)
}

// initConfig reads in config file and ENV variables if set.
// func initConfig() {
// 	if cfgFile != "" {
// 		// Use config file from the flag.
// 		viper.SetConfigFile(cfgFile)
// 	} else {
// 		// Find home directory.
// 		home, err := homedir.Dir()
// 		if err != nil {
// 			fmt.Println(err)
// 			os.Exit(1)
// 		}

// 		// Search config in home directory with name ".hello-world" (without extension).
// 		viper.AddConfigPath(home)
// 		viper.SetConfigName(".hello-world")
// 	}

// 	viper.AutomaticEnv() // read in environment variables that match

// 	// If a config file is found, read it in.
// 	if err := viper.ReadInConfig(); err == nil {
// 		fmt.Println("Using config file:", viper.ConfigFileUsed())
// 	}
// }
func initConfig() {
	for k := range flags.AsEnvVariables(conf, "", false) {
		log.CheckErr("Unable to bind environment variable", viper.BindEnv(strings.ToLower(strings.Replace(k, "_", ".", -1)), "SCHEMA_"+k), zap.String("var", "SCHEMA_"+k))
	}

	switch {
	case cfgFile != "":
		// If the config file doesn't exists, let's exit
		if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
			log.Bg().Fatal("File doesn't exists", zap.Error(err))
		}
		fmt.Println("Reading configuration file", cfgFile)

		viper.SetConfigFile(cfgFile)
		if err := viper.ReadInConfig(); err != nil {
			log.Bg().Fatal("Unable to read config", zap.Error(err))
		}
	default:
		defaults.SetDefaults(conf)
	}

	if err := viper.Unmarshal(conf); err != nil {
		log.Bg().Fatal("Unable to parse config", zap.Error(err))
	}
}
