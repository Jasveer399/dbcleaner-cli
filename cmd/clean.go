package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/Jasveer399/dbcleaner-cli/internal/cleaner"
	"github.com/Jasveer399/dbcleaner-cli/internal/config"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean database tables",
	Long:  "Clean database tables by truncating or dropping them based on configuration",
	Run:   runClean,
}

var (
	dryRun     bool
	backup     bool
	truncate   bool
	configFile string
)

func init() {
	rootCmd.AddCommand(cleanCmd)
	cleanCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Show what would be cleaned without actually doing it")
	cleanCmd.Flags().BoolVar(&backup, "backup", false, "Create backup before cleaning")
	cleanCmd.Flags().BoolVar(&truncate, "truncate", false, "Truncate tables instead of dropping")
	cleanCmd.Flags().StringVar(&configFile, "config", "dbcleaner.yaml", "Configuration file path")
}

func runClean(cmd *cobra.Command, args []string) {
	cfg, err := config.LoadConfig(configFile)

	if err != nil {
		color.Red("❌ Failed to load configuration: %v", err)
		color.Yellow("💡 Run 'dbcleaner init' to create a configuration file")
		return
	}

	if dryRun {
		color.Yellow("🔍 Running in dry-run mode...")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cleanerInstance := cleaner.New(cfg)

	if err := cleanerInstance.Connect(ctx); err != nil {
		color.Red("❌ Failed to connect to database: %v", err)
		return
	}
	defer cleanerInstance.Close()

	// Get tables
	tables, err := cleanerInstance.GetTables()
	if err != nil {
		color.Red("❌ Failed to get tables: %v", err)
		return
	}

	if len(tables) == 0 {
		color.Yellow("ℹ️  No tables found in database")
		return
	}

	// Show tables and confirm
	color.Cyan("📋 Found %d tables:", len(tables))
	for _, table := range tables {
		fmt.Printf("  - %s\n", table)
	}

	if !dryRun {
		prompt := promptui.Prompt{
			Label:     "Are you sure you want to clean these tables? (yes/no)",
			IsConfirm: true,
		}

		if _, err := prompt.Run(); err != nil {
			color.Yellow("❌ Operation cancelled")
			return
		}
	}

	// Perform cleaning
	opts := cleaner.CleanOptions{
		DryRun:   dryRun,
		Backup:   backup,
		Truncate: truncate,
	}

	if err := cleanerInstance.CleanTables(tables, opts); err != nil {
		color.Red("❌ Cleaning failed: %v", err)
		return
	}

	if dryRun {
		color.Green("✅ Dry run completed successfully")
	} else {
		color.Green("✅ Database cleaned successfully")
	}
}
