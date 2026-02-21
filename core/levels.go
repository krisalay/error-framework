package core

type ErrorLevel int

const (
	LevelDebug ErrorLevel = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

func (l ErrorLevel) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}
