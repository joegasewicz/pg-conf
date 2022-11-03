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

	// //------ Unhandled ----------- //
	// PGPassFile `PGPASSFILE` behaves the same as the passfile connection parameter.
	PGPassFile string
	// PGChannelBinding `PGCHANNELBINDING` behaves the same as the channel_binding connection parameter
	PGChannelBinding string
	// PGService `PGSERVICE` behaves the same as the service connection parameter
	PGService string
	// PGServiceFile `PGSERVICEFILE` specifies the name of the per-user connection service file.
	// Defaults to ~/.pg_service.conf, or %APPDATA%\postgresql\.pg_service.conf on Microsoft Windows
	PGServiceFile string
	// PGOptions `PGOPTIONS` behaves the same as the options connection parameter.
	PGOptions string
	// PGAppName `PGAPPNAME` behaves the same as the application_name connection parameter.
	PGAppName string
	// PGRequireSSL `PGREQUIRESSL` behaves the same as the requiressl connection parameter.
	// This environment variable is deprecated in favor of the `PGSSLMODE` variable;
	// setting both variables suppresses the effect of this one.
	PGRequireSSL string
	// PGSSLCompression `PGSSLCOMPRESSION` behaves the same as the sslcompression connection parameter.
	PGSSLCompression string
	// PGSSLCert `PGSSLCERT` behaves the same as the sslcert connection parameter.
	PGSSLCert string
	// PGSSLKey `PGSSLKEY` behaves the same as the sslkey connection parameter.
	PGSSLKey string
	// PGSSLRootCert `PGSSLROOTCERT` behaves the same as the sslrootcert connection parameter.
	PGSSLRootCert string
	// PGSSLCrl `PGSSLCRL` behaves the same as the sslcrl connection parameter.
	PGSSLCrl string
	// PGSSLCrlDir `PGSSLCRLDIR` `behaves the same as the sslcrldir connection parameter.
	PGSSLCrlDir string
	// PGSSLSni `PGSSLSNI` behaves the same as the sslsni connection parameter.
	PGSSLSni string
	// PGRequirePeer `PGREQUIREPEER` behaves the same as the requirepeer connection parameter.
	PGRequirePeer string
	// PGSSLMinProtocolVersion `PGSSLMINPROTOCOLVERSION` behaves the same as the ssl_min_protocol_version connection parameter.
	PGSSLMinProtocolVersion string
	// PGSSLMaxProtocolVersion `PGSSLMAXPROTOCOLVERSION` behaves the same as the ssl_max_protocol_version connection parameter.
	PGSSLMaxProtocolVersion string
	// PGGSSENCMode `PGGSSENCMODE` behaves the same as the gssencmode connection parameter.
	PGGSSENCMode string
	// PGKRBSRVName `PGKRBSRVNAME` behaves the same as the krbsrvname connection parameter.
	PGKRBSRVName string
	// PGGSSLib `PGGSSLIB` behaves the same as the gsslib connection parameter.
	PGGSSLib string
	// PGConnectTimeout `PGCONNECT_TIMEOUT` behaves the same as the connect_timeout connection parameter.
	PGConnectTimeout string
	// PGClientEncoding `PGCLIENTENCODING` behaves the same as the client_encoding connection parameter.
	PGClientEncoding string
	// PGTargetSessionAttrs `PGTARGETSESSIONATTRS` behaves the same as the target_session_attrs connection parameter.
	PGTargetSessionAttrs string
	// PGDataStyle `PGDATESTYLE` sets the default style of date/time representation. (Equivalent to SET datestyle TO ....)
	PGDataStyle string
	// PGTz `PGTZ` sets the default time zone. (Equivalent to SET timezone TO ....)
	PGTz string
	// PGGeqo `PGGEQO` sets the default mode for the genetic query optimizer. (Equivalent to SET geqo TO ....)
	PGGeqo string
	// PGSysConfDir `PGSYSCONFDIR` sets the directory containing the pg_service.conf file and in a future version possibly other system-wide configuration files.
	PGSysConfDir string
	// PGLocalEDir `PGLOCALEDIR` sets the directory containing the locale files for message localization.
	PGLocalEDir string
	// \\------ Unhandled ----------- \\

	// PGSSLMode `PGSSLMODE` behaves the same as the sslmode connection parameter.
	PGSSLMode string
}

// Update creates a new `PostgresConfig` type
func Update(p *PostgresConfig) error {
	p.PGHost = os.Getenv("PGHOST")
	p.PGHostAddr = os.Getenv("PGHOSTADDR")
	p.PGPort = os.Getenv("PGPORT")
	p.PGDatabase = os.Getenv("PGDATABASE")
	p.PGUser = os.Getenv("PGUSER")
	p.PGPassword = os.Getenv("PGPASSWORD")
	p.PGPassFile = os.Getenv("PGPASSFILE")
	p.PGChannelBinding = os.Getenv("PGCHANNELBINDING")
	p.PGService = os.Getenv("PGSERVICE")
	p.PGServiceFile = os.Getenv("PGSERVICEFILE")
	p.PGOptions = os.Getenv("PGOPTIONS")
	p.PGAppName = os.Getenv("PGAPPNAME")
	p.PGSSLMode = os.Getenv("PGSSLMODE")
	p.PGRequireSSL = os.Getenv("PGREQUIRESSL")
	p.PGSSLCompression = os.Getenv("PGSSLCOMPRESSION")
	p.PGSSLCert = os.Getenv("PGSSLCERT")
	p.PGSSLKey = os.Getenv("PGSSLKEY")
	p.PGSSLRootCert = os.Getenv("PGSSLROOTCERT")
	p.PGSSLCrl = os.Getenv("PGSSLCRL")
	p.PGSSLCrlDir = os.Getenv("PGSSLCRLDIR")
	p.PGSSLSni = os.Getenv("PGSSLSNI")
	p.PGRequirePeer = os.Getenv("PGREQUIREPEER")
	p.PGSSLMinProtocolVersion = os.Getenv("PGSSLMINPROTOCOLVERSION")
	p.PGSSLMaxProtocolVersion = os.Getenv("PGSSLMaxProtocolVersion")
	p.PGGSSENCMode = os.Getenv("PGGSSENCMODE")
	p.PGKRBSRVName = os.Getenv("PGKRBSRVNAME")
	p.PGGSSLib = os.Getenv("PGGSSLIB")
	p.PGConnectTimeout = os.Getenv("PGCONNECT_TIMEOUT")
	p.PGClientEncoding = os.Getenv("PGCLIENTENCODING")
	p.PGTargetSessionAttrs = os.Getenv("PGTARGETSESSIONATTRS")
	p.PGDataStyle = os.Getenv("PGTARGETSESSIONATTRS")
	p.PGTz = os.Getenv("PGTZ")
	p.PGGeqo = os.Getenv("PGGEQO")
	p.PGSysConfDir = os.Getenv("PGSYSCONFDIR")
	p.PGLocalEDir = os.Getenv("PGLOCALEDIR")

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
