package cmdutil

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
	if errors.As(err, &viper.ConfigFileNotFoundError{}) {
		err = viper.SafeWriteConfig() // Create empty config file
	}
	cobra.CheckErr(err)
}
