package pg_conf

import (
	"fmt"
	"os"
)

// PostgresConfig Postgres 34.15. Environment Variables
type PostgresConfig struct {
	// PGHost `PGHOST` behaves the same as the host connection parameter
	PGHost string
	// PGHostAddr `PGHOSTADDR` behaves the same as the hostaddr connection parameter. This can be set instead of or in addition to `PGHOST` to avoid DNS lookup overhead.
	PGHostAddr string
	// PGPort `PGPORT`
	PGPort string
	// PGDatabase `PGDATABASE` behaves the same as the dbname connection parameter.
	PGDatabase string
	// PGUser `PGUSER` behaves the same as the user connection parameter.
	PGUser string
	// PGPassword `PGPASSWORD` behaves the same as the password connection parameter.
	// Use of this environment variable is not recommended for security reasons,
	// as some operating systems allow non-root users to see process environment variables via ps;
	// instead consider using a password file.
	PGPassword string
	// PGSSLMode `PGSSLMODE` behaves the same as the sslmode connection parameter.
	PGSSLMode string
}

// Update creates a new `PostgresConfig` type
func Update(p *PostgresConfig) error {
	if p.PGHost == "" {
		p.PGHost = os.Getenv("PGHOST")
	}
	if p.PGPort == "" {
		p.PGPort = os.Getenv("PGPORT")
	}
	if p.PGDatabase == "" {
		p.PGDatabase = os.Getenv("PGDATABASE")
	}
	if p.PGUser == "" {
		p.PGUser = os.Getenv("PGUSER")
	}
	if p.PGPassword == "" {
		p.PGPassword = os.Getenv("PGPASSWORD")
	}
	if p.PGSSLMode == "" {
		p.PGSSLMode = os.Getenv("PGSSLMODE")
	}
	return nil
}

// GetPostgresConnStr creates a Postgres connection string
func (c *PostgresConfig) GetPostgresConnStr() string {
	var connStr string
	// set defaults
	if c.PGHost != "" {
		connStr += fmt.Sprintf("host=%s", c.PGHost)
	} else {
		connStr += fmt.Sprintf("host=%s", "localhost")
	}
	if c.PGUser != "" {
		connStr += fmt.Sprintf(" user=%s", c.PGUser)
	} else {
		connStr += fmt.Sprintf(" user=%s", "postgres")
	}
	if c.PGDatabase != "" {
		connStr += fmt.Sprintf(" dbname=%s", c.PGDatabase)
	} else {
		connStr += fmt.Sprintf(" dbname=%s", "postgres")
	}
	if c.PGPassword != "" {
		connStr += fmt.Sprintf(" password=%s", c.PGPassword)
	}
	if c.PGPort != "" {
		connStr += fmt.Sprintf(" port=%s", c.PGPort)
	} else {
		connStr += fmt.Sprintf(" port=%s", "5432")
	}
	if c.PGSSLMode != "" {
		connStr += fmt.Sprintf(" sslmode=%s", c.PGSSLMode)
	} else {
		connStr += fmt.Sprintf(" sslmode=%s", "disable")
	}

	return connStr
}
