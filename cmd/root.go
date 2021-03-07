package cmd

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	userConfigDir, err := os.UserConfigDir()
	cobra.CheckErr(err)

	viper.AddConfigPath(userConfigDir)
	viper.SetConfigName("toggl")
	viper.SetConfigType("json")

	err = viper.ReadInConfig()
	if !errors.As(err, &viper.ConfigFileNotFoundError{}) {
		cobra.CheckErr(err)
	}
}

var rootCmd = &cobra.Command{
	Use:     "toggl",
	Short:   "The unofficial CLI client for Toggl Track",
	Version: "0.1.0",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	err := rootCmd.Execute()
	cobra.CheckErr(err)
}
