package config

var defaults = []option{
	{
		name:        "logger.level",
		typing:      "string",
		value:       "debug",
		description: "Level of logging",
	},
	{
		name:        "logger.pathFile",
		typing:      "string",
		value:       "./logs/app.log",
		description: "Path to the log file",
	},
	{
		name:        "listen.httpHost",
		typing:      "string",
		value:       "localhost",
		description: "Server host",
	},
	{
		name:        "listen.httpPort",
		typing:      "string",
		value:       "8080",
		description: "Server port",
	},
	{
		name:        "listen.shutdowntimeout",
		typing:      "duration",
		value:       "5s",
		description: "Timeout for graceful shutdown",
	},

	{
		name:        "worker.keepAliveTimeout",
		typing:      "duration",
		value:       "5s",
		description: "Timeout for worker keep-alive",
	},
	{
		name:        "postgres.URL",
		typing:      "string",
		value:       "postgres://postgres:post@localhost:5432/books?sslmode=disable",
		description: "Postgres database URL",
	},
	{
		name:        "postgres.Timeout",
		typing:      "duration",
		value:       "2s",
		description: "Timeout for database connection",
	},
	{
		name:        "postgres.ConnectAttempts",
		typing:      "int",
		value:       10,
		description: "Number of database connection attempts",
	},
	{
		name:        "postgres.MaxIdleTime",
		typing:      "duration",
		value:       "5m",
		description: "Maximum idle time for database connections",
	},
	{
		name:        "postgres.MaxOpenConns",
		typing:      "int",
		value:       100,
		description: "Maximum number of open database connections",
	},
	{
		name:        "postgres.HealthCheckPeriod",
		typing:      "duration",
		value:       "5m",
		description: "Period for database health check",
	},
	{
		name:        "worker.keepAliveTimeout",
		typing:      "duration",
		value:       "5s",
		description: "Timeout for worker keep-alive",
	},
}

type option struct {
	name        string
	typing      string
	value       interface{}
	description string
}
