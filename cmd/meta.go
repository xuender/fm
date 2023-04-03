package cmd

import (
	"github.com/spf13/cobra"
	"github.com/youthlin/t"
)

func metaCmd(cmd *cobra.Command) *cobra.Command {
	cmd.Short = t.T("Show file meta")
	// nolint: lll
	cmd.Long = t.T("Show file meta.\n\n  Image, Video, Audio, Archive, Documents, Font, Application, Java, Golang, JavaScript")
	cmd.Run = func(cmd *cobra.Command, args []string) {
		service := InitMeta(cmd)

		for _, arg := range args {
			service.Info(arg).Output(cmd.OutOrStderr())
		}
	}

	return cmd
}
