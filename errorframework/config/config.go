package config

type Config struct {
	Logger LoggerConfig

	Trace TraceConfig

	StackTrace StackTraceConfig

	Database DatabaseConfig

	Validator ValidatorConfig
}

type LoggerConfig struct {
	ConsoleEnabled bool
	FileEnabled    bool
	FilePath       string
	Level          string
	Encoding       string // json or console
}

type TraceConfig struct {
	Enabled bool
}

type StackTraceConfig struct {
	Enabled bool
}

type DatabaseConfig struct {
	Type string // pgx, mysql, mongo

	IncludeConstraintDetails bool
	IncludeTableDetails      bool
}

type ValidatorConfig struct {
	Enabled bool
}
