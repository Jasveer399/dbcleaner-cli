package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "dbcleaner",
	Short: "A powerful database cleaner CLI tool",
	Long: `DBCleaner is a CLI tool for cleaning and managing databases.
It supports PostgreSQL, MySQL, and MongoDB with flexible configuration options.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfigLoad)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./dbcleaner.yaml)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
}

func initConfigLoad() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME/.dbcleaner")
		viper.SetConfigName("dbcleaner")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Printf("Using config file: %s\n", viper.ConfigFileUsed())
	}
}
