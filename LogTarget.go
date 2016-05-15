package leadlogger

import "fmt"
import "os"

var (
	TargetStdOut = &StdOutTarget{MaxLevel: LEVEL_DEBUG, BufferSize: 0}
)

type LogTarget interface {
	Init()
	MaximumLevel() LogLevel
	//If push always returns true, there is immediate logging
	Push(*logEntry) bool
	Log()
}

type StdOutTarget struct {
	MaxLevel   LogLevel
	BufferSize int
	buffer     []*logEntry
}

func (t *StdOutTarget) Init() {
	t.buffer = make([]*logEntry, 0, t.BufferSize+1)
}
func (t *StdOutTarget) MaximumLevel() LogLevel { return t.MaxLevel }
func (t *StdOutTarget) Push(l *logEntry) bool {
	t.buffer = append(t.buffer, l)
	if len(t.buffer) >= t.BufferSize {
		return true
	}
	return false
}

func (t *StdOutTarget) Log() {
	for i, l := range t.buffer {
		fmt.Println(l)
		t.buffer[i] = nil
	}
	t.buffer = t.buffer[:0]
}

type FileTarget struct {
	Path       string
	MaxLevel   LogLevel
	BufferSize int
	buffer     []*logEntry
}

func (t *FileTarget) Init() {
	fp, err := os.OpenFile(t.Path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	fp.Close()
	t.buffer = make([]*logEntry, 0, t.BufferSize+1)
}
func (t *FileTarget) MaximumLevel() LogLevel { return t.MaxLevel }
func (t *FileTarget) Push(l *logEntry) bool {
	t.buffer = append(t.buffer, l)
	if len(t.buffer) >= t.BufferSize {
		return true
	}
	return false
}

func (t *FileTarget) Log() {
	fp, err := os.OpenFile(t.Path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fp.Close(); err != nil {
			panic(err)
		}
	}()
	defer func() {
		if err := fp.Sync(); err != nil {
			panic(err)
		}
	}()
	for i, l := range t.buffer {
		_, err := fp.WriteString(l.String())
		if err != nil {
			panic(err)
		}
		t.buffer[i] = nil
	}
	t.buffer = t.buffer[:0]
}
