package logging

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"
	"time"
)

type Level int

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
	FATAL
)

func (l Level) Symbol() string {
	if l < DEBUG || l > FATAL {
		return ""
	}

	return []string{"[+]", "[*]", "[~]", "[!]", "[x]"}[l]
}

func (l Level) Name() string {
	if l < DEBUG || l > FATAL {
		return ""
	}

	return []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}[l]
}

func (l Level) Color() Color {
	if l < DEBUG || l > FATAL {
		return ColorWhite
	}

	return []Color{ColorCyan, ColorWhite, ColorYellow, ColorRed, ColorBrightRed}[l]
}

type Color string

const (
	ColorBlack   Color = "\u001b[30m"
	ColorRed     Color = "\u001b[31m"
	ColorGreen   Color = "\u001b[32m"
	ColorYellow  Color = "\u001b[33m"
	ColorBlue    Color = "\u001b[34m"
	ColorMagenta Color = "\u001b[35m"
	ColorCyan    Color = "\u001b[36m"
	ColorWhite   Color = "\u001b[37m"

	ColorBrightBlack   Color = "\u001b[90m"
	ColorBrightRed     Color = "\u001b[91m"
	ColorBrightGreen   Color = "\u001b[92m"
	ColorBrightYellow  Color = "\u001b[93m"
	ColorBrightBlue    Color = "\u001b[94m"
	ColorBrightMagenta Color = "\u001b[95m"
	ColorBrightCyan    Color = "\u001b[96m"
	ColorBrightWhite   Color = "\u001b[97m"

	ColorReset Color = "\u001b[0m"
)

func colorString(c Color, s string) string {
	return string(c) + s + string(ColorReset)
}

type Formatter interface {
	Format(lvl Level, msg string) string
}

// format
// TIME: $t
// LEVEL SYMBOL $l
// LEVEL NAME: $L
// FILE SHORT: $f
// FILE LONG: $F
// MESSAGE: $m
type SimpleFormatter struct {
	MsgFormat  string
	TimeFormat string
	LineEnd    string
}

func (f SimpleFormatter) Format(lvl Level, msg string) string {
	str := f.MsgFormat + f.LineEnd
	str = strings.ReplaceAll(str, "$t", time.Now().Format(f.TimeFormat))
	str = strings.ReplaceAll(str, "$l", lvl.Symbol())
	str = strings.ReplaceAll(str, "$L", lvl.Name())
	str = strings.ReplaceAll(str, "$m", msg)

	_, file, line, ok := runtime.Caller(3)
	if !ok {
		file = "???"
		line = 0
	}

	str = strings.ReplaceAll(str, "$f", fmt.Sprintf("%s:%d", path.Base(file), line))
	str = strings.ReplaceAll(str, "$F", fmt.Sprintf("%s:%d", file, line))

	return str
}

type ColorFormatter struct {
	SimpleFormatter
	DefaultColor Color // TODO: implement
}

func (f ColorFormatter) color(c Color, s string) string {
	return fmt.Sprintf("%s%s%s", c, s, f.DefaultColor)
}

func (f ColorFormatter) Format(lvl Level, msg string) string {
	str := colorString(f.DefaultColor, f.MsgFormat+f.LineEnd)
	str = strings.ReplaceAll(str, "$t", f.color(ColorGreen, time.Now().Format(f.TimeFormat)))
	str = strings.ReplaceAll(str, "$l", f.color(lvl.Color(), lvl.Symbol()))
	str = strings.ReplaceAll(str, "$L", f.color(lvl.Color(), lvl.Name()))
	str = strings.ReplaceAll(str, "$m", f.color(lvl.Color(), msg))

	_, file, line, ok := runtime.Caller(3)
	if !ok {
		file = "???"
		line = 0
	}

	str = strings.ReplaceAll(str, "$f", f.color(ColorBrightBlack, fmt.Sprintf("%s:%d", path.Base(file), line)))
	str = strings.ReplaceAll(str, "$F", f.color(ColorBrightBlack, fmt.Sprintf("%s:%d", file, line)))

	return str
}

var DefaultFormatter = SimpleFormatter{
	"$l [$t] :: $m",
	"2006-01-02 15:04:05",
	"\n",
}

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

var DefaultLogger = &SimpleLogger{
	INFO,
	DefaultFormatter,
	os.Stdout,
	sync.Mutex{},
}

func SetFormatter(format Formatter) { DefaultLogger.Form = format }
func SetLevel(level Level)          { DefaultLogger.Lvl = level }
func SetOutput(out io.Writer)       { DefaultLogger.Out = out }

func Debug(msg string)            { DefaultLogger.Log(DEBUG, msg) }
func Debugf(msg string, a ...any) { DefaultLogger.Log(DEBUG, fmt.Sprintf(msg, a...)) }

func Info(msg string)            { DefaultLogger.Log(INFO, msg) }
func Infof(msg string, a ...any) { DefaultLogger.Log(INFO, fmt.Sprintf(msg, a...)) }

func Print(msg string)            { DefaultLogger.Log(INFO, msg) }
func Printf(msg string, a ...any) { DefaultLogger.Log(INFO, fmt.Sprintf(msg, a...)) }

func Warn(msg string)            { DefaultLogger.Log(WARN, msg) }
func Warnf(msg string, a ...any) { DefaultLogger.Log(WARN, fmt.Sprintf(msg, a...)) }

func Error(msg string)            { DefaultLogger.Log(ERROR, msg) }
func Errorf(msg string, a ...any) { DefaultLogger.Log(ERROR, fmt.Sprintf(msg, a...)) }

func Fatal(msg string)            { DefaultLogger.Log(FATAL, msg); os.Exit(1) }
func Fatalf(msg string, a ...any) { DefaultLogger.Log(FATAL, fmt.Sprintf(msg, a...)); os.Exit(1) }
