package cmd

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func NewCmdRoot() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     "toggl",
		Short:   "The unofficial CLI client for Toggl Track",
		Version: "0.1.0",
		PreRunE: invokeHelpByDefault,
		Run:     func(cmd *cobra.Command, args []string) {},
	}
	return rootCmd
}

func invokeHelpByDefault(cmd *cobra.Command, args []string) error {
	hasChangedFlag := false
	cmd.Flags().Visit(func(_ *pflag.Flag) { hasChangedFlag = true })
	if !hasChangedFlag && len(args) == 0 {
		return cmd.Help()
	}
	return nil
}

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
