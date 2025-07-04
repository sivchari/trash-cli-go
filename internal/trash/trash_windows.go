//go:build windows

package trash

import (
	"io"
	"os"
	"path/filepath"
)

func getTrashDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, "AppData", "Local", "Trash"), nil
}

func moveFile(src, dst string) error {
	err := os.Rename(src, dst)
	if err != nil {
		return copyFile(src, dst)
	}
	return nil
}

func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}

	dstFile, err := os.Create(dst)
	if err != nil {
		srcFile.Close()
		return err
	}

	_, err = io.Copy(dstFile, srcFile)
	srcFile.Close()
	
	if err != nil {
		dstFile.Close()
		os.Remove(dst)
		return err
	}

	if err := dstFile.Sync(); err != nil {
		dstFile.Close()
		os.Remove(dst)
		return err
	}
	
	dstFile.Close()

	return os.Remove(src)
}