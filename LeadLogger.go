package leadlogger

import "time"

var (
	_channelBufferSize = 100
	_format            = DEFAULT_FORMAT
	_logChan           chan *logEntry
	_syncChan          chan bool
	_timeOut           = time.Second * 5
)

type status int

func Duration(duration time.Duration) {
	_timeOut = duration
}

func Format(format LogFormat) {
	_format = format
}

func Start(targets ...LogTarget) {
	_logChan = make(chan *logEntry, _channelBufferSize)
	_syncChan = make(chan bool)
	go handleLogging(targets)
}

func Stop() {
	close(_logChan)
	<-_syncChan
	close(_syncChan)
}

func Running() bool {
	return _running
}
