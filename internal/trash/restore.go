package trash

import (
	"fmt"
	"os"
	"path/filepath"
)

func RestoreTrash() error {
	items, err := ListTrash()
	if err != nil {
		return err
	}

	if len(items) == 0 {
		fmt.Println("Trash is empty")
		return nil
	}

	return fmt.Errorf("use RestoreTrashUI for interactive restore")
}

func RestoreFile(item TrashItem) error {
	if _, err := os.Stat(item.TrashPath); os.IsNotExist(err) {
		return fmt.Errorf("trash file not found: %s", item.TrashPath)
	}

	if _, err := os.Stat(item.OriginalPath); err == nil {
		return fmt.Errorf("cannot restore: file already exists at %s", item.OriginalPath)
	}

	originalDir := filepath.Dir(item.OriginalPath)
	if err := os.MkdirAll(originalDir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", originalDir, err)
	}

	if err := os.Rename(item.TrashPath, item.OriginalPath); err != nil {
		return fmt.Errorf("failed to restore file to %s: %w", item.OriginalPath, err)
	}

	if err := os.Remove(item.InfoPath); err != nil {
		return fmt.Errorf("failed to remove trash info file: %w", err)
	}

	return nil
}
