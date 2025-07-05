package filesystem

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/HJyup/forge/config"
)

func GetLastModifiedTime(path string) (time.Time, error) {
	var latestTime time.Time

	err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if info.IsDir() && (info.Name() == config.NodeModulesDir || info.Name() == config.GitDir || strings.HasPrefix(info.Name(), config.HiddenDirPrefix)) {
			return filepath.SkipDir
		}

		if !info.IsDir() && info.ModTime().After(latestTime) {
			latestTime = info.ModTime()
		}

		return nil
	})

	return latestTime, err
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func DirExists(path string) bool {
	info, err := os.Stat(path)
	return !os.IsNotExist(err) && info.IsDir()
}
