package cmd

import (
	"fmt"
	"log"
	"path/filepath"
	"sort"

	"github.com/HJyup/forge/config"
	"github.com/HJyup/forge/service/parser"
	tui "github.com/HJyup/forge/service/tui/root"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   config.AppName,
	Short: config.AppShortDesc,
	Long:  config.AppLongDesc,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.LoadConfig()
		basePath := ""

		if len(args) > 0 {
			basePath = args[0]
		} else {
			basePath = cfg.DefaultBasePath
		}

		absPath, err := filepath.Abs(basePath)
		if err != nil {
			log.Fatalf("Failed to get absolute path: %v", err)
		}

		projects, err := parser.ScanAllProjects(absPath)
		if err != nil {
			log.Fatalf("Scan failed: %v", err)
		}

		if len(projects) == 0 {
			fmt.Println("No projects found!")
			return
		}

		sort.Slice(projects, func(i, j int) bool {
			return projects[i].LastModified.After(projects[j].LastModified)
		})

		if err := tui.RunRootTUI(projects); err != nil {
			log.Fatalf("TUI error: %v", err)
		}
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Additional commands can be added here in the future.
}
