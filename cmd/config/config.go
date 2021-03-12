package config

import (
	"github.com/spf13/cobra"
	"github.com/ta9mi1shi1/toggl/cmdutil"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "config",
		Short:   "Manage toggl configs",
		PreRunE: cmdutil.InvokeHelpByDefault,
		Run:     func(cmd *cobra.Command, args []string) {},
	}
	return cmd
}
