package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgresDB struct {
	db     *sql.DB
	config *ConnectionConfig
}

func NewPostgresDB(config *ConnectionConfig) *PostgresDB {
	return &PostgresDB{
		config: config,
	}
}

func (p *PostgresDB) Connect(ctx context.Context) error {

	sslmode := p.config.SSLMode
	if sslmode == "" {
		sslmode = "disable"
	}
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", p.config.Host, p.config.Port, p.config.User, p.config.Password, p.config.DBName, sslmode)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}

	p.db = db
	return p.db.PingContext(ctx)
}

func (p *PostgresDB) Close() error {
	if p.db != nil {
		return p.db.Close()
	}
	return nil
}

func (p *PostgresDB) TestConnection() error {
	return p.db.Ping()
}

func (p *PostgresDB) GetTables() ([]string, error) {
	query := `SELECT table_name FROM information_schema.tables 
              WHERE table_schema = 'public' AND table_type = 'BASE TABLE'`

	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tables []string
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			return nil, err
		}
		tables = append(tables, table)
	}

	return tables, nil
}

func (p *PostgresDB) TruncateTable(table string) error {
	query := fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", table)
	_, err := p.db.Exec(query)
	return err
}

func (p *PostgresDB) DropTable(table string) error {
	query := fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE", table)
	_, err := p.db.Exec(query)
	return err
}

func (p *PostgresDB) GetTableRowCount(table string) (int64, error) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", table)
	var count int64
	err := p.db.QueryRow(query).Scan(&count)
	return count, err
}

func (p *PostgresDB) BackupDatabase(path string) error {
	// Implementation for pg_dump backup
	return fmt.Errorf("backup not implemented yet")
}
