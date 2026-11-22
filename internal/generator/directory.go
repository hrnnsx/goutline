package generator

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func PrepDir(projectName string) (string, error) {

	path, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current directory: %w", err)
	}

	projectPath := filepath.Join(path, projectName)
	entries, err := os.ReadDir(projectPath)

	if err == nil {
		if len(entries) == 0 {
			return projectPath, nil
		}

		return "", errors.New("Project directory is not empty, please check it first.")
	}

	if errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(projectPath, 0755)
		if err != nil {
			return "", fmt.Errorf("Failed to create project directory: %w", err)
		}
	} else {
		return "", fmt.Errorf("failed to read project directory: %w", err)
	}

	return projectPath, nil
}
