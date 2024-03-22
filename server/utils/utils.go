package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

const UPLOAD_DIR = "/Users/zulfiqarahmed/apps/vault/upload"

func CreateUploadDirectoryFor(name string) (string, error) {
	dir := filepath.Join(UPLOAD_DIR, name)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return "", fmt.Errorf("unable to create upload directory for %s", name)
		}
	}

	return dir, nil
}
