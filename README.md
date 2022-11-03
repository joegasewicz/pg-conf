# PG Conf
Sets Postgres environment variables to a config type

### Example
```go
os.Setenv("PGHOST", "www.google.com")
os.Setenv("PGUSER", "John")
os.Setenv("PGPASSWORD", "wizard")
os.Setenv("PGDATABASE", "happy_db")
os.Setenv("PGPORT", "1234")
os.Setenv("PGSSLMODE", "verify-ca")

type Config struct {
    *pg_conf.PostgresConfig
}

cfg := Config{
    PostgresConfig: &PostgresConfig{
      // Add defaults here if requires...
    },
}

pg_conf.Update(cfg.PostgresConfig)

connStr := cfg.PostgresConfig.GetPostgresConnStr()

expected := "host=www.google.com user=John dbname=happy_db password=wizard port=1234 sslmode=verify-ca"
```