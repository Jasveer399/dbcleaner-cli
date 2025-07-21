package database

import "context"

type Database interface {
	Connect(ctx context.Context) error
	Close() error
	TestConnection() error
	GetTables() ([]string, error)
	TruncateTable(table string) error
	DropTable(table string) error
	GetTableRowCount(table string) (int64, error)
	BackupDatabase(path string) error
}

type ConnectionConfig struct {
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}
