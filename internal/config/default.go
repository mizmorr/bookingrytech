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
}

type option struct {
	name        string
	typing      string
	value       interface{}
	description string
}
