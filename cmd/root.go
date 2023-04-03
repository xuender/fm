package cmd

import (
	"os"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/xuender/kit/logs"
	"github.com/youthlin/t"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	root := &cobra.Command{
		Use:   "fm",
		Short: t.T("File Meta"),
		Long:  t.T("File Meta.\n\n  Move file based on meta."),
		Run: func(cmd *cobra.Command, args []string) {
			defer func() {
				if err := recover(); err != nil {
					logs.E.Println(err)
				}
			}()

			InitMove(cmd).Move(lo.Map(lo.Must1(os.ReadDir(".")), func(entry os.DirEntry, _ int) string {
				return entry.Name()
			}))
		},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if lo.Must1(cmd.Flags().GetBool("debug")) {
				logs.SetLevel(logs.Debug)
			} else {
				logs.SetLevel(logs.Info)
			}
		},
	}
	root.AddCommand(initCmd(&cobra.Command{Use: "init", Aliases: []string{"i"}}))
	root.AddCommand(metaCmd(&cobra.Command{Use: "meta", Aliases: []string{"m"}}))
	root.PersistentFlags().BoolP("debug", "d", false, t.T("Debug Mode"))

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
