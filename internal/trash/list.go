package trash

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func ListTrash() ([]TrashItem, error) {
	trashDir, err := GetTrashDir()
	if err != nil {
		return nil, err
	}

	infoDir := filepath.Join(trashDir, "info")

	entries, err := os.ReadDir(infoDir)
	if err != nil {
		return nil, err
	}

	var items []TrashItem

	for _, entry := range entries {
		if !strings.HasSuffix(entry.Name(), ".trashinfo") {
			continue
		}

		infoPath := filepath.Join(infoDir, entry.Name())
		item, err := parseTrashInfo(infoPath)
		if err != nil {
			continue
		}

		items = append(items, item)
	}

	return items, nil
}

func parseTrashInfo(infoPath string) (TrashItem, error) {
	file, err := os.Open(infoPath)
	if err != nil {
		return TrashItem{}, err
	}
	defer file.Close()

	var item TrashItem
	item.InfoPath = infoPath

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if after, ok := strings.CutPrefix(line, "Path="); ok {
			item.OriginalPath = after
		} else if after, ok := strings.CutPrefix(line, "DeletionDate="); ok {
			dateStr := after
			if parsedTime, err := time.Parse("2006-01-02T15:04:05", dateStr); err == nil {
				item.DeletionDate = parsedTime
			}
		}
	}

	baseName := strings.TrimSuffix(filepath.Base(infoPath), ".trashinfo")
	trashDir := filepath.Dir(filepath.Dir(infoPath))
	item.TrashPath = filepath.Join(trashDir, "files", baseName)

	return item, nil
}

func PrintTrashList(items []TrashItem) {
	for _, item := range items {
		fmt.Printf("%s %s\n", item.DeletionDate.Format("2006-01-02 15:04:05"), item.OriginalPath)
	}
}
