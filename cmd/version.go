package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Version = "dev"
	Commit  = "none"
	Date    = "unknown"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  "Print the version number of trash-cli-go",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("trash-cli-go %s (commit: %s, built: %s)\n", Version, Commit, Date)
	},
}
