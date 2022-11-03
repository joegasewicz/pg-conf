package pg_conf

import "fmt"

type PostgresConfig struct {
	Host     string `pg_conf:""`
	User     string
	Password string
	Name     string
	Port     uint16
	SSLMode  string
}

// New creates a new `PostgresConfig` type
func New() *PostgresConfig {
	return &PostgresConfig{
		Host:     "localhost",
		User:     "admin",
		Password: "admin",
		Name:     "identity_db",
		Port:     5433,
		SSLMode:  "disable",
	}
}

// GetPostgresConnStr creates a Postgres connection string
func (c *PostgresConfig) GetPostgresConnStr() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		c.Host,
		c.User,
		c.Password,
		c.Name,
		c.Port,
		c.SSLMode,
	)
}
