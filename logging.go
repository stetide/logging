package logging

import (
	"fmt"
	"io"
	"os"
	"sync"
)

var DefaultFormatter = SimpleFormatter{
	"$l [$t] :: $m",
	"2006-01-02 15:04:05",
	"\n",
}

var DefaultLogger = &SimpleLogger{
	DEBUG,
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
