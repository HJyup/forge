package parser

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/HJyup/forge/config"
	"github.com/HJyup/forge/service/filesystem"
	"github.com/HJyup/forge/service/git"
)

type ProjectSummary struct {
	Name         string
	Path         string
	LastModified time.Time
}

type PackageJSON struct {
	Name string `json:"name"`
}

func readPackageJSON(path string) (PackageJSON, error) {
	var pkg PackageJSON

	packagePath := filepath.Join(path, config.PackageJSONFile)
	data, err := os.ReadFile(packagePath)
	if err != nil {
		return pkg, err
	}

	err = json.Unmarshal(data, &pkg)
	return pkg, err
}

func ScanProjectSummary(path string) (ProjectSummary, error) {
	pkg, _ := readPackageJSON(path)

	name := filepath.Base(path)
	if pkg.Name != "" {
		name = pkg.Name
	}

	lastMod, _ := git.GetLastCommitTime(path)
	if lastMod.IsZero() {
		lastMod, _ = filesystem.GetLastModifiedTime(path)
	}

	return ProjectSummary{
		Name:         name,
		Path:         path,
		LastModified: lastMod,
	}, nil
}

func ScanAllProjects(basePath string) ([]ProjectSummary, error) {
	var projects []ProjectSummary

	entries, err := os.ReadDir(basePath)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		projectPath := filepath.Join(basePath, entry.Name())

		if !isValidProject(projectPath) {
			fmt.Printf(config.SkipNoPackageJSON, projectPath)
			continue
		}

		summary, err := ScanProjectSummary(projectPath)
		if err != nil {
			fmt.Printf(config.SkipGeneralError, projectPath, err)
			continue
		}

		projects = append(projects, summary)
	}

	return projects, nil
}

func isValidProject(path string) bool {
	packagePath := filepath.Join(path, config.PackageJSONFile)
	return filesystem.FileExists(packagePath)
}
