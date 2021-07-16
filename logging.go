package logging

import (
	"fmt"
	"io"
	"strings"
	"time"
)

/*
LOGGING FORMATTING
	%(TIME) -> time
	%(LEVEL) -> LogLevel Symbol
	%(LEVELNAME) -> LogLevel Name
	%(MSG) -> message
*/

type LogLevel int

func (l LogLevel) valid() bool {
	return l >= 0 && l <= 5
}

func (l LogLevel) String() string {
	if !l.valid() {
		return ""
	}
	return []string{"DEBUG", "INFO", "WARNING", "ERROR", "CRITICAL"}[l]
}

func (l LogLevel) Symbol() string {
	if !l.valid() {
		return ""
	}
	return []string{"[+]", "[*]", "[~]", "[!]", "[x]"}[l]
}

func (l LogLevel) Name() string {
	return l.String()
}

const (
	DEBUG    LogLevel = iota // [+] additional debug information
	INFO                     // [*] info for the user
	WARNING                  // [~] unexpected behaviour that might cause problems
	ERROR                    // [!] error but the program can still run
	CRITICAL                 // [x] critical error -> program is terminated
)

// Standard Logger
// contains unexported fields
type Logger struct {
	format     string
	timeformat string
	writer     io.Writer
	level      LogLevel
	nextLine   string
}

// NewLogger returns a new Logger with format as default format.
// The format could be eg. "%(TIME): %(MSG)" - %(TIME) will be replaced with the current time
// and %(MSG) will be replaced with the message
// writer if the writer is not nil, the message will not only be printed to the stdout but will
// also be written to the writer (which could be a file)
// level is the minimum level that the logger will log. If the level is set to INFO logger.Debug
// will do nothing
func NewLogger(format string, writer io.Writer, level LogLevel) *Logger {
	return &Logger{
		format,
		"",
		writer,
		level,
		"\n",
	}
}

// Sets the logging level of the Logger
func (l *Logger) SetLevel(level LogLevel) {
	if level.valid() {
		l.level = level
	}
}

func (l *Logger) SetWriter(writer io.Writer) {
	l.writer = writer
}

// golangs time Layout eg: "2021-08-18 15:03"
func (l *Logger) SetTimeFormat(layout string) {
	l.timeformat = layout
}

func (l *Logger) SetFormat(format string) {
	l.format = format
}

// SetNextLine is defaulted to a new line at the end of every log call.
func (l *Logger) SetNextLine(endl string) {
	l.nextLine = endl
}

// LOGGING FORMATTING
// 	   %(TIME) -> time
// 	   %(LEVEL) -> LogLevel Symbol
// 	   %(LEVELNAME) -> LogLevel Name
//     %(MSG) -> message
func (l *Logger) log(lvl LogLevel, msg string) {
	if lvl < l.level {
		return
	}

	m := l.format
	m += l.nextLine
	m = strings.ReplaceAll(m, "%(TIME)", time.Now().Format(l.timeformat))
	m = strings.ReplaceAll(m, "%(LEVEL)", l.level.Symbol())
	m = strings.ReplaceAll(m, "%(LEVELNAME)", l.level.Name())
	m = strings.ReplaceAll(m, "%(MSG)", msg)

	if l.writer != nil {
		l.writer.Write([]byte(m))
	}
	fmt.Print(m)
}

func (l *Logger) Log(level LogLevel, msg string) {
	l.log(level, msg)
}

func (l *Logger) Debug(msg string) {
	l.log(DEBUG, msg)
}

func (l *Logger) Info(msg string) {
	l.log(INFO, msg)
}

func (l *Logger) Warning(msg string) {
	l.log(WARNING, msg)
}

func (l *Logger) Error(msg string) {
	l.log(ERROR, msg)
}

func (l *Logger) Critical(msg string) {
	l.log(CRITICAL, msg)
}

var defaultLogger = &Logger{
	format:     "%(LEVEL) %(TIME): %(MSG)",
	timeformat: "2021-08-18 15:06:12",
	writer:     nil,
	level:      DEBUG,
	nextLine:   "\n",
}

func SetLevel(level LogLevel) {
	defaultLogger.SetLevel(level)
}

func SetWriter(writer io.Writer) {
	defaultLogger.SetWriter(writer)
}

func Log(level LogLevel, msg string) {
	defaultLogger.Log(level, msg)
}

func Debug(msg string) {
	defaultLogger.Debug(msg)
}

func Info(msg string) {
	defaultLogger.Info(msg)
}

func Warning(msg string) {
	defaultLogger.Warning(msg)
}

func Error(msg string) {
	defaultLogger.Error(msg)
}

func Critical(msg string) {
	defaultLogger.Critical(msg)
}
