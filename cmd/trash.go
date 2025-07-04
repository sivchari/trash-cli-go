package cmd

import (
	"fmt"
	"os"

	"github.com/sivchari/trash-cli-go/internal/trash"
	"github.com/spf13/cobra"
)

var trashCmd = &cobra.Command{
	Use:   "put [files...]",
	Short: "Move files to trash",
	Long:  "Move files and directories to the trash bin.",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		for _, filePath := range args {
			if err := trash.MoveToTrash(filePath); err != nil {
				fmt.Fprintf(os.Stderr, "Error trashing %s: %v\n", filePath, err)
				continue
			}
		}
		return nil
	},
}
