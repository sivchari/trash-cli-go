package cmd

import (
	"fmt"

	"github.com/sivchari/trash-cli-go/internal/trash"
	"github.com/sivchari/trash-cli-go/internal/ui"
	"github.com/spf13/cobra"
)

var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore a trashed file",
	Long:  "Restore a file from the trash bin to its original location using an interactive interface.",
	RunE: func(cmd *cobra.Command, args []string) error {
		items, err := trash.ListTrash()
		if err != nil {
			return err
		}

		if len(items) == 0 {
			return nil
		}

		if err := ui.RunRestoreUI(items); err != nil {
			return fmt.Errorf("failed to restore file: %w", err)
		}

		return nil
	},
}
