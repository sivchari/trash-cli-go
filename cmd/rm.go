package cmd

import (
	"github.com/sivchari/trash-cli-go/internal/trash"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm [patterns...]",
	Short: "Remove files from trash",
	Long:  "Remove files from trash matching the given patterns.",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return trash.RemoveFromTrash(args)
	},
}
