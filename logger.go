package logging

import (
	"fmt"
	"io"
	"os"
	"sync"
)

type Logger interface {
	Log(lvl Level, msg string)
	Debug(msg string)
	Debugf(msg string, a ...any)
	Info(msg string)
	Infof(msg string, a ...any)
	Print(msg string)
	Printf(msg string, a ...any)
	Warn(msg string)
	Warnf(msg string, a ...any)
	Error(msg string)
	Errorf(msg string, a ...any)
	Fatal(msg string)
	Fatalf(msg string, a ...any)
}

type SimpleLogger struct {
	Lvl  Level
	Form Formatter
	Out  io.Writer
	mu   sync.Mutex
}

func New(out io.Writer) *SimpleLogger {
	return &SimpleLogger{
		INFO,
		DefaultFormatter,
		out,
		sync.Mutex{},
	}
}

func NewSimpleLogger(level Level, formatter Formatter, out io.Writer) *SimpleLogger {
	return &SimpleLogger{
		level,
		formatter,
		out,
		sync.Mutex{},
	}
}

func (l *SimpleLogger) Log(lvl Level, msg string) {
	if lvl < l.Lvl {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	fmt.Fprint(l.Out, l.Form.Format(lvl, msg))
}

func (l *SimpleLogger) Debug(msg string)            { l.Log(DEBUG, msg) }
func (l *SimpleLogger) Debugf(msg string, a ...any) { l.Log(DEBUG, fmt.Sprintf(msg, a...)) }

func (l *SimpleLogger) Info(msg string)            { l.Log(INFO, msg) }
func (l *SimpleLogger) Infof(msg string, a ...any) { l.Log(INFO, fmt.Sprintf(msg, a...)) }

func (l *SimpleLogger) Print(msg string)            { l.Log(INFO, msg) }
func (l *SimpleLogger) Printf(msg string, a ...any) { l.Log(INFO, fmt.Sprintf(msg, a...)) }

func (l *SimpleLogger) Warn(msg string)            { l.Log(WARN, msg) }
func (l *SimpleLogger) Warnf(msg string, a ...any) { l.Log(WARN, fmt.Sprintf(msg, a...)) }

func (l *SimpleLogger) Error(msg string)            { l.Log(ERROR, msg) }
func (l *SimpleLogger) Errorf(msg string, a ...any) { l.Log(ERROR, fmt.Sprintf(msg, a...)) }

func (l *SimpleLogger) Fatal(msg string)            { l.Log(FATAL, msg); os.Exit(1) }
func (l *SimpleLogger) Fatalf(msg string, a ...any) { l.Log(FATAL, fmt.Sprintf(msg, a...)); os.Exit(1) }

type MultiLogger struct {
	Lvl     Level
	Loggers []Logger
}

func NewMultiLogger(level Level, loggers ...Logger) *MultiLogger {
	return &MultiLogger{
		level,
		loggers,
	}
}

func (m *MultiLogger) Log(lvl Level, msg string) {
	if lvl < m.Lvl {
		return
	}

	for _, l := range m.Loggers {
		l.Log(lvl, msg)
	}
}

func (l *MultiLogger) Debug(msg string)            { l.Log(DEBUG, msg) }
func (l *MultiLogger) Debugf(msg string, a ...any) { l.Log(DEBUG, fmt.Sprintf(msg, a...)) }

func (l *MultiLogger) Info(msg string)            { l.Log(INFO, msg) }
func (l *MultiLogger) Infof(msg string, a ...any) { l.Log(INFO, fmt.Sprintf(msg, a...)) }

func (l *MultiLogger) Print(msg string)            { l.Log(INFO, msg) }
func (l *MultiLogger) Printf(msg string, a ...any) { l.Log(INFO, fmt.Sprintf(msg, a...)) }

func (l *MultiLogger) Warn(msg string)            { l.Log(WARN, msg) }
func (l *MultiLogger) Warnf(msg string, a ...any) { l.Log(WARN, fmt.Sprintf(msg, a...)) }

func (l *MultiLogger) Error(msg string)            { l.Log(ERROR, msg) }
func (l *MultiLogger) Errorf(msg string, a ...any) { l.Log(ERROR, fmt.Sprintf(msg, a...)) }

func (l *MultiLogger) Fatal(msg string)            { l.Log(FATAL, msg); os.Exit(1) }
func (l *MultiLogger) Fatalf(msg string, a ...any) { l.Log(FATAL, fmt.Sprintf(msg, a...)); os.Exit(1) }
