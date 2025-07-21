package cmd

import (
	"fmt"

	"github.com/Jasveer399/dbcleaner-cli/internal/config"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize DBCleaner configuration",
	Long:  "Initialize DBCleaner configuration by setting up database connection details",
	Run:   initConfig,
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func initConfig(cmd *cobra.Command, args []string) {
	color.Green("🚀 Welcome to DBCleaner CLI Setup!")
	fmt.Println()

	dbPrompt := promptui.Select{
		Label: "Select Database Type",
		Items: []string{"PostgreSQL", "MySQL", "MongoDB"},
	}

	_, dbType, err := dbPrompt.Run()
	if err != nil {
		fmt.Printf("❌ Database selection failed: %v", err)
		return
	}

	var cfg *config.Config
	switch dbType {
	case "PostgreSQL":
		cfg = setupPostgreSQL()
	case "MySQL":
		cfg = setupMySQL()
	case "MongoDB":
		cfg = setupMongoDB()
	}

	if cfg == nil {
		color.Red("❌ Configuration setup failed")
		return
	}

	if err := cfg.Save("dbcleaner.yaml"); err != nil {
		color.Red("❌ Failed to save configuration: %v", err)
		return
	}

	color.Green("✅ Configuration saved successfully to dbcleaner.yaml")
}

func setupPostgreSQL() *config.Config {
	cfg := &config.Config{}
	cfg.Database.Driver = "postgres"

	prompts := []struct {
		label    string
		field    *string
		validate func(string) error
	}{
		{"Host", &cfg.Database.Host, nil},
		{"Port", &cfg.Database.Port, nil},
		{"Username", &cfg.Database.User, nil},
		{"Password", &cfg.Database.Password, nil},
		{"Database Name", &cfg.Database.DBName, nil},
	}

	for _, p := range prompts {
		prompt := promptui.Prompt{
			Label: p.label,
		}
		if p.field == &cfg.Database.Password {
			prompt.Mask = '*'
		}

		value, err := prompt.Run()
		if err != nil {
			color.Red("❌ Input failed: %v", err)
			return nil
		}
		*p.field = value
	}

	return cfg
}

func setupMySQL() *config.Config {
	cfg := &config.Config{}
	cfg.Database.Driver = "mysql"

	prompts := []struct {
		label string
		field *string
	}{
		{"Host", &cfg.Database.Host},
		{"Port", &cfg.Database.Port},
		{"Username", &cfg.Database.User},
		{"Password", &cfg.Database.Password},
		{"Database Name", &cfg.Database.DBName},
	}

	for _, p := range prompts {
		prompt := promptui.Prompt{
			Label: p.label,
		}
		if p.field == &cfg.Database.Password {
			prompt.Mask = '*'
		}

		value, err := prompt.Run()
		if err != nil {
			color.Red("❌ Input failed: %v", err)
			return nil
		}
		*p.field = value
	}

	return cfg
}

func setupMongoDB() *config.Config {
	cfg := &config.Config{}
	cfg.Database.Driver = "mongodb"

	prompts := []struct {
		label string
		field *string
	}{
		{"Host", &cfg.Database.Host},
		{"Port", &cfg.Database.Port},
		{"Username", &cfg.Database.User},
		{"Password", &cfg.Database.Password},
		{"Database Name", &cfg.Database.DBName},
	}

	for _, p := range prompts {
		prompt := promptui.Prompt{
			Label: p.label,
		}
		if p.field == &cfg.Database.Password {
			prompt.Mask = '*'
		}

		value, err := prompt.Run()
		if err != nil {
			color.Red("❌ Input failed: %v", err)
			return nil
		}
		*p.field = value
	}

	return cfg
}
