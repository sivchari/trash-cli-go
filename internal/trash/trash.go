package trash

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type TrashItem struct {
	OriginalPath string
	DeletionDate time.Time
	TrashPath    string
	InfoPath     string
}

func GetTrashDir() (string, error) {
	trashDir, err := getTrashDir()
	if err != nil {
		return "", err
	}

	if err := os.MkdirAll(filepath.Join(trashDir, "files"), 0755); err != nil {
		return "", err
	}

	if err := os.MkdirAll(filepath.Join(trashDir, "info"), 0755); err != nil {
		return "", err
	}

	return trashDir, nil
}

func MoveToTrash(filePath string) error {
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return err
	}

	fileInfo, err := os.Stat(absPath)
	if os.IsNotExist(err) {
		return fmt.Errorf("file does not exist: %s", absPath)
	}
	if err != nil {
		return err
	}

	trashDir, err := GetTrashDir()
	if err != nil {
		return err
	}

	fileName := filepath.Base(absPath)
	uniqueName, err := generateUniqueName(trashDir, fileName)
	if err != nil {
		return err
	}

	trashFilePath := filepath.Join(trashDir, "files", uniqueName)
	infoFilePath := filepath.Join(trashDir, "info", uniqueName+".trashinfo")

	if err := moveFile(absPath, trashFilePath); err != nil {
		return fmt.Errorf("failed to move file to trash: %w", err)
	}

	if err := preserveFileMode(trashFilePath, fileInfo.Mode()); err != nil {
		moveFile(trashFilePath, absPath)
		return fmt.Errorf("failed to preserve file permissions: %w", err)
	}

	infoContent := fmt.Sprintf(`[Trash Info]
Path=%s
DeletionDate=%s
`, absPath, time.Now().Format("2006-01-02T15:04:05"))

	if err := os.WriteFile(infoFilePath, []byte(infoContent), 0644); err != nil {
		os.Remove(trashFilePath)
		return fmt.Errorf("failed to create trash info file: %w", err)
	}

	return nil
}

func generateUniqueName(trashDir, fileName string) (string, error) {
	baseName := fileName
	counter := 0

	for {
		var candidateName string
		if counter == 0 {
			candidateName = baseName
		} else {
			timestamp := time.Now().Format("20060102_150405")
			candidateName = fmt.Sprintf("%s_%s_%d", baseName, timestamp, counter)
		}

		filesPath := filepath.Join(trashDir, "files", candidateName)
		infoPath := filepath.Join(trashDir, "info", candidateName+".trashinfo")

		if _, err := os.Stat(filesPath); os.IsNotExist(err) {
			if _, err := os.Stat(infoPath); os.IsNotExist(err) {
				return candidateName, nil
			}
		}

		counter++
		if counter > 1000 {
			return "", fmt.Errorf("unable to generate unique name for %s", fileName)
		}
	}
}

func preserveFileMode(filePath string, mode os.FileMode) error {
	return os.Chmod(filePath, mode)
}

