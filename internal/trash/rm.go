package trash

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func RemoveFromTrash(patterns []string) error {
	items, err := ListTrash()
	if err != nil {
		return err
	}

	if len(items) == 0 {
		return nil
	}

	var toRemove []TrashItem

	if len(patterns) == 0 {
		return fmt.Errorf("no patterns specified")
	}

	for _, item := range items {
		for _, pattern := range patterns {
			matched, err := filepath.Match(pattern, filepath.Base(item.OriginalPath))
			if err != nil {
				continue
			}
			if matched || strings.Contains(item.OriginalPath, pattern) {
				toRemove = append(toRemove, item)
				break
			}
		}
	}

	if len(toRemove) == 0 {
		return nil
	}

	for _, item := range toRemove {
		if err := removeTrashItem(item); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to remove %s: %v\n", item.OriginalPath, err)
			continue
		}
	}

	return nil
}

func removeTrashItem(item TrashItem) error {
	if err := os.Remove(item.TrashPath); err != nil && !os.IsNotExist(err) {
		return err
	}

	if err := os.Remove(item.InfoPath); err != nil && !os.IsNotExist(err) {
		return err
	}

	return nil
}
