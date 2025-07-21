package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database DatabaseConfig `yaml:"database"`
	Cleaner  CleanerConfig  `yaml:"cleaner,omitempty"`
}

type DatabaseConfig struct {
	Driver   string `yaml:"driver"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode,omitempty"`
}

type CleanerConfig struct {
	DryRun        bool     `yaml:"dry_run"`
	BackupBefore  bool     `yaml:"backup_before"`
	ExcludeTables []string `yaml:"exclude_tables,omitempty"`
	IncludeTables []string `yaml:"include_tables,omitempty"`
	TruncateOnly  bool     `yaml:"truncate_only"`
}

func (c *Config) Save(filename string) error {
	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func LoadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	return &cfg, err
}
