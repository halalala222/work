package consts

import "path"

var (
	defaultLogDir           = "internal/logs"
	DefaultConfigFileType   = "json"
	DefaultConfigFileName   = "config"
	DefaultLogFilePath      = logFilePath("zap.log")
	DefaultSQLLogFilePath   = logFilePath("sql.log")
	DefaultFiberLogFilePath = logFilePath("fiber.log")
)

func logFilePath(filename string) string {
	return path.Join(defaultLogDir, filename)
}
