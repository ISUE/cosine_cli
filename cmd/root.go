package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var (
	rootCmd = &cobra.Command{
		Use:   "cosine_cli",
		Short: "A Command Line Interface to handle the cosine project",
	}

	cfgFile         string
	phoneTrialPaths []string
	viconTrialPath  string
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "",
		"config file (default is $HOME/.cosine_cli.yaml)")

	rootCmd.PersistentFlags().StringArrayVar(&phoneTrialPaths, "phone", nil,
		"An array of the path to the phone trial directories")

	rootCmd.PersistentFlags().StringVar(&viconTrialPath, "vicon", "",
		"The path to the Vicon captured trials folder")

	viper.BindPFlag("phone", rootCmd.Flags().Lookup("phone"))
	viper.BindPFlag("vicon", rootCmd.Flags().Lookup("vicon"))
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cosine_cli")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
