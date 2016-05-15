package leadlogger

import "fmt"

func log(format LogFormat, level LogLevel, tag string, message []interface{}) {
	_logChan <- &logEntry{format, level, tag, fmt.Sprintln(message...)}
}

func FormatEmergency(format LogFormat, tag string, message ...interface{}) {
	log(format, LEVEL_EMERGENCY, tag, message)
}

func Emergency(tag string, message ...interface{}) {
	FormatEmergency(_format, tag, message...)
}

func FormatAlert(format LogFormat, tag string, message ...interface{}) {
	log(format, LEVEL_ALERT, tag, message)
}

func Alert(tag string, message ...interface{}) {
	FormatAlert(_format, tag, message...)
}

func FormatCritical(format LogFormat, tag string, message ...interface{}) {
	log(format, LEVEL_CRITICAL, tag, message)
}

func Critical(tag string, message ...interface{}) {
	FormatCritical(_format, tag, message...)
}

func FormatError(format LogFormat, tag string, message ...interface{}) {
	log(format, LEVEL_ERROR, tag, message)
}

func Error(tag string, message ...interface{}) {
	FormatError(_format, tag, message...)
}

func FormatWarning(format LogFormat, tag string, message ...interface{}) {
	log(format, LEVEL_WARNING, tag, message)
}

func Warning(tag string, message ...interface{}) {
	FormatWarning(_format, tag, message...)
}

func FormatNotice(format LogFormat, tag string, message ...interface{}) {
	log(format, LEVEL_NOTICE, tag, message)
}

func Notice(tag string, message ...interface{}) {
	FormatNotice(_format, tag, message...)
}

func FormatInformational(format LogFormat, tag string, message ...interface{}) {
	log(format, LEVEL_INFORMATIONAL, tag, message)
}

func Informational(tag string, message ...interface{}) {
	FormatInformational(_format, tag, message...)
}

func FormatDebug(format LogFormat, tag string, message ...interface{}) {
	log(format, LEVEL_DEBUG, tag, message)
}

func Debug(tag string, message ...interface{}) {
	FormatDebug(_format, tag, message...)
}
