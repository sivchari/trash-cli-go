package cmd

import (
	"fmt"
	"strconv"

	"github.com/sivchari/trash-cli-go/internal/trash"
	"github.com/spf13/cobra"
)

var emptyCmd = &cobra.Command{
	Use:   "empty [days]",
	Short: "Empty the trash",
	Long:  "Empty the trash bin. If days is specified, only files older than the specified number of days will be removed.",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		days := 0

		if len(args) > 0 {
			var err error
			days, err = strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("invalid number of days: %s", args[0])
			}
		}

		if err := trash.EmptyTrash(days); err != nil {
			return fmt.Errorf("failed to empty trash: %w", err)
		}

		return nil
	},
}
