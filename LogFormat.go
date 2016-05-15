package leadlogger

import "fmt"
import "time"

var (
	DEFAULT_FORMAT = func(level LogLevel, tag, message string) string {
		return fmt.Sprintf("%v - %v [%v]: %v",
			time.Now().Format(time.UnixDate),
			level,
			tag,
			message)
	}
)

type LogFormat func(level LogLevel, tag, message string) string
