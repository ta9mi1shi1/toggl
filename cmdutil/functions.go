package cmdutil

import (
	"github.com/spf13/cobra"
)

func InvokeHelpByDefault(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return cmd.Help()
	}
	return nil
}
