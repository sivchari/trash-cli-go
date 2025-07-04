package trash

import (
	"os"
	"time"
)

func EmptyTrash(olderThanDays int) error {
	items, err := ListTrash()
	if err != nil {
		return err
	}

	cutoffTime := time.Now().AddDate(0, 0, -olderThanDays)

	for _, item := range items {
		if olderThanDays > 0 && item.DeletionDate.After(cutoffTime) {
			continue
		}

		if err := os.Remove(item.TrashPath); err != nil && !os.IsNotExist(err) {
			continue
		}

		if err := os.Remove(item.InfoPath); err != nil && !os.IsNotExist(err) {
			continue
		}
	}

	return nil
}
