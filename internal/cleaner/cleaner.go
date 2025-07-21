package cleaner

import (
	"context"
	"fmt"

	"github.com/Jasveer399/dbcleaner-cli/internal/config"
	"github.com/Jasveer399/dbcleaner-cli/internal/database"
	"github.com/fatih/color"
)

type Cleaner struct {
	db     database.Database
	config *config.Config
}

type CleanOptions struct {
	DryRun   bool
	Backup   bool
	Truncate bool
}

func New(cfg *config.Config) *Cleaner {
	dbConfig := &database.ConnectionConfig{
		Driver:   cfg.Database.Driver,
		Host:     cfg.Database.Host,
		Port:     cfg.Database.Port,
		User:     cfg.Database.User,
		Password: cfg.Database.Password,
		DBName:   cfg.Database.DBName,
		SSLMode:  cfg.Database.SSLMode,
	}

	var db database.Database

	switch cfg.Database.Driver {
	case "postgres":
		db = database.NewPostgresDB(dbConfig)
	case "mysql":
		// db = database.NewMySQLDB(dbConfig)
		panic("MySQL support not implemented yet")
	case "mongodb":
		// db = database.NewMongoDB(dbConfig)
		panic("MongoDB support not implemented yet")
	default:
		panic(fmt.Sprintf("unsupported database driver: %s", cfg.Database.Driver))
	}

	return &Cleaner{
		db:     db,
		config: cfg,
	}
}

func (c *Cleaner) Connect(ctx context.Context) error {
	return c.db.Connect(ctx)
}

func (c *Cleaner) Close() error {
	return c.db.Close()
}

func (c *Cleaner) GetTables() ([]string, error) {
	return c.db.GetTables()
}

func (c *Cleaner) CleanTables(tables []string, opts CleanOptions) error {
	for _, table := range tables {
		if opts.DryRun {
			count, err := c.db.GetTableRowCount(table)
			if err != nil {
				color.Yellow("⚠️  Could not get row count for table %s: %v", table, err)
				continue
			}

			action := "truncate"
			if !opts.Truncate {
				action = "drop"
			}
			color.Yellow("🔍 Would %s table '%s' (%d rows)", action, table, count)
			continue
		}

		if opts.Backup {
			color.Blue("💾 Creating backup for table: %s", table)
			// Implement backup logic here
		}

		var err error
		if opts.Truncate {
			color.Blue("🗑️  Truncating table: %s", table)
			err = c.db.TruncateTable(table)
		} else {
			color.Blue("🗑️  Dropping table: %s", table)
			err = c.db.DropTable(table)
		}

		if err != nil {
			color.Red("❌ Failed to clean table %s: %v", table, err)
			return err
		}

		color.Green("✅ Successfully cleaned table: %s", table)
	}

	return nil
}
