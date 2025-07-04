//go:build !windows

package trash

import (
	"os"
	"path/filepath"
)

func getTrashDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, ".local", "share", "Trash"), nil
}

func moveFile(src, dst string) error {
	return os.Rename(src, dst)
}