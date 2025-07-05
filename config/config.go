package config

import (
	"os"
	"path/filepath"
)

type Config struct {
	// Default base path for scanning projects
	DefaultBasePath string

	// Directories to ignore during filesystem scanning
	IgnoredDirectories []string

	// File extensions to consider when determining last modified time
	TrackedExtensions []string

	// Maximum number of projects to display
	MaxProjects int

	// Enable debug logging
	Debug bool
}

func DefaultConfig() *Config {
	home, _ := os.UserHomeDir()
	defaultPath := filepath.Join(home, DefaultDevPath)

	return &Config{
		DefaultBasePath: defaultPath,
		IgnoredDirectories: []string{
			NodeModulesDir,
			GitDir,
			".DS_Store",
			".vscode",
			".idea",
			"dist",
			"build",
			"target",
			"vendor",
		},
		TrackedExtensions: []string{
			".js", ".ts", ".jsx", ".tsx",
			".go", ".py", ".java", ".cpp",
			".c", ".h", ".css", ".scss",
			".html", ".md", ".json", ".yaml",
			".yml", ".toml", ".xml",
		},
		MaxProjects: 100,
		Debug:       false,
	}
}

func LoadConfig() *Config {
	cfg := DefaultConfig()

	if path := os.Getenv("FORGE_BASE_PATH"); path != "" {
		cfg.DefaultBasePath = path
	}

	if os.Getenv("FORGE_DEBUG") == "true" {
		cfg.Debug = true
	}

	return cfg
}
