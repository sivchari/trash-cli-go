package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "trash",
	Short: "A trash CLI tool written in Go",
	Long:  "A cross-platform trash CLI tool that follows the FreeDesktop.org trash standard.",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(trashCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(restoreCmd)
	rootCmd.AddCommand(emptyCmd)
	rootCmd.AddCommand(rmCmd)
	rootCmd.AddCommand(versionCmd)
}
