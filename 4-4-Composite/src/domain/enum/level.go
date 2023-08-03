package enum

type Level int

const (
	TRACE Level = iota
	INFO
	DEBUG
	WARN
	ERROR
)

func (l Level) String() string {
	return map[Level]string{
		TRACE: "TRACE",
		INFO:  "INFO",
		DEBUG: "DEBUG",
		WARN:  "WARN",
		ERROR: "ERROR",
	}[l]
}

func ParseLevel(key string) Level {
	return map[string]Level{
		"TRACE": TRACE,
		"INFO":  INFO,
		"DEBUG": DEBUG,
		"WARN":  WARN,
		"ERROR": ERROR,
	}[key]
}
