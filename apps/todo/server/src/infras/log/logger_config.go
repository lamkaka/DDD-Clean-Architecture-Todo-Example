package log

type Config struct {
	Level Level
}

type Level string

const (
	PanicLevel Level = "Panic"
	ErrorLevel Level = "Error"
	WarnLevel  Level = "Warn"
	InfoLevel  Level = "Info"
	DebugLevel Level = "Debug"
	TraceLevel Level = "Trace"
)
