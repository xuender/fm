package cmd

import (
	"github.com/spf13/cobra"
	"github.com/youthlin/t"
)

func initCmd(cmd *cobra.Command) *cobra.Command {
	cmd.Short = t.T("Init FM config")
	cmd.Long = t.T("Init FM config.\n\nCreate config file .fm.toml.")
	cmd.Run = func(cmd *cobra.Command, args []string) {
		InitUI(cmd).Init()
	}

	return cmd
}
