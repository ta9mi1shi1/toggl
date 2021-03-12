package root

import (
	"github.com/spf13/cobra"
	"github.com/ta9mi1shi1/toggl/cmd/config"
	"github.com/ta9mi1shi1/toggl/cmdutil"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "toggl",
		Short:   "The unofficial CLI client for Toggl Track",
		Version: "0.1.0",
		PreRunE: cmdutil.InvokeHelpByDefault,
	}

	cmd.AddCommand(config.NewCmd())

	return cmd
}
