package leadlogger

const (
	LEVEL_EMERGENCY     = 0
	LEVEL_ALERT         = 1
	LEVEL_CRITICAL      = 2
	LEVEL_ERROR         = 3
	LEVEL_WARNING       = 4
	LEVEL_NOTICE        = 5
	LEVEL_INFORMATIONAL = 6
	LEVEL_DEBUG         = 7
)

type LogLevel int

func (l LogLevel) String() string {
	switch l {
	case LEVEL_EMERGENCY:
		return "EMERGENCY"
	case LEVEL_ALERT:
		return "ALERT"
	case LEVEL_CRITICAL:
		return "CRITICAL"
	case LEVEL_ERROR:
		return "ERROR"
	case LEVEL_WARNING:
		return "WARNING"
	case LEVEL_NOTICE:
		return "NOTICE"
	case LEVEL_INFORMATIONAL:
		return "INFORMATIONAL"
	case LEVEL_DEBUG:
		return "DEBUG"
	default:
		return "UNDEFINED LEVEL"
	}
}
