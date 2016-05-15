package leadlogger

import "time"

var (
	_running = false
)

func handleLogging(targets []LogTarget) {
	_running = true

	for _, target := range targets {
		target.Init()
	}
	listen(targets)
	_running = false
	_syncChan <- true
}

func listen(targets []LogTarget) {
	for {
		select {
		case message, ok := <-_logChan:
			pushToTargets(message, targets)
			if !ok {
				for len(_logChan) > 0 {
					pushToTargets(<-_logChan, targets)
				}
				logAllTargets(targets)
				return
			}
		case <-time.After(_timeOut):
			logAllTargets(targets)
		}
	}
}

func logAllTargets(targets []LogTarget) {
	for _, target := range targets {
		target.Log()
	}
}

func pushToTargets(message *logEntry, targets []LogTarget) {
	if message == nil {
		return
	}
	for _, target := range targets {
		if target.MaximumLevel() >= message.level {
			if target.Push(message) {
				target.Log()
			}
		}
	}
}
