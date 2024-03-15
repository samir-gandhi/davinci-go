package test

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func ReadJsonFile(path string) (string, error) {
	_, currentPath, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("failed to get current path")
	}
	flowByte, err := os.ReadFile(filepath.Join(filepath.Dir(currentPath), filepath.Clean(path)))
	if err != nil {
		return "", err
	}

	return string(flowByte), nil
}
