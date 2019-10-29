package cmd

import (
	//. "../refactor"
	. "../refactor/unused"
	"github.com/spf13/cobra"
)

var refactorCmd *cobra.Command = &cobra.Command{
	Use:   "refactor",
	Short: "auto refactor code",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		moveConfig := cmd.Flag("move").Value.String()
		path := cmd.Flag("path").Value.String()

		if moveConfig != "" && path != "" {
			//app := NewMoveClassApp(moveConfig, path)
			//app.Analysis()

			app2 := NewRemoveUnusedImportApp(moveConfig, path)
			app2.Analysis()
		}
	},
}

func init() {
	rootCmd.AddCommand(refactorCmd)

	refactorCmd.PersistentFlags().StringP("path", "p", "", "path")
	refactorCmd.PersistentFlags().StringP("move", "m", "", "with config example -m config.files")
}
