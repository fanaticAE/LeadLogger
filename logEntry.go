package leadlogger

type logEntry struct {
	format  LogFormat
	level   LogLevel
	tag     string
	message string
}

func (l logEntry) String() string {
	return l.format(l.level, l.tag, l.message)
}
