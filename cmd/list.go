package cmd

import (
	"fmt"

	"github.com/sivchari/trash-cli-go/internal/trash"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List trashed files",
	Long:  "List all files currently in the trash bin.",
	RunE: func(cmd *cobra.Command, args []string) error {
		items, err := trash.ListTrash()
		if err != nil {
			return fmt.Errorf("failed to list trash: %w", err)
		}

		if len(items) == 0 {
			return nil
		}

		trash.PrintTrashList(items)
		return nil
	},
}
