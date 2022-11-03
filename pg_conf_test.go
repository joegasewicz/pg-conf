package pg_conf

import (
	"os"
	"testing"
)

func TestDefaultConnectionString(t *testing.T) {

	type Config struct {
		*PostgresConfig
	}

	cfg := Config{
		PostgresConfig: &PostgresConfig{
			PGHost:     "",
			PGUser:     "",
			PGPassword: "",
			PGDatabase: "",
			PGPort:     "",
			PGSSLMode:  "",
		},
	}

	Update(cfg.PostgresConfig)

	connStr := cfg.PostgresConfig.GetPostgresConnStr()

	expected := "host=localhost user=postgres dbname=postgres port=5432 sslmode=disable"
	if connStr != expected {
		t.Logf("\nexpected: %s\nactual:   %s\n", expected, connStr)
		t.Fail()
	}
}

func TestCallerConnectionString(t *testing.T) {

	os.Setenv("PGHOST", "www.google.com")
	os.Setenv("PGUSER", "John")
	os.Setenv("PGPASSWORD", "wizard")
	os.Setenv("PGDATABASE", "happy_db")
	os.Setenv("PGPORT", "1234")
	os.Setenv("PGSSLMODE", "verify-ca")

	type Config struct {
		*PostgresConfig
	}

	cfg := Config{
		PostgresConfig: &PostgresConfig{
			PGHost:     "",
			PGUser:     "",
			PGPassword: "",
			PGDatabase: "",
			PGPort:     "",
			PGSSLMode:  "",
		},
	}

	Update(cfg.PostgresConfig)

	connStr := cfg.PostgresConfig.GetPostgresConnStr()

	expected := "host=www.google.com user=John dbname=happy_db password=wizard port=1234 sslmode=verify-ca"
	if connStr != expected {
		t.Logf("\nexpected: %s\nactual:   %s\n", expected, connStr)
		t.Fail()
	}
}
