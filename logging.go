package logging

// logging package
// simple python like logging module
// It's recommended to call BasicConfig first,
// otherwise the standarts will be used.

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// LogLevel int from 0 - 4
type LogLevel int

// LogLevel constants
const (
	DEBUG    LogLevel = 0
	INFO     LogLevel = 1
	WARNING  LogLevel = 2
	ERROR    LogLevel = 3
	CRITICAL LogLevel = 4
)

// logging prefixes
const (
	pDebug    = "[+]"
	pInfo     = "[*]"
	pWarning  = "[-]"
	pError    = "[!]"
	pCritical = "[!!]"
)

var (
	logLevel    LogLevel
	logFormat   string = "[{prefix}] {message}"
	timeFormat  string = "15:24:05"
	logFile     *os.File
	printOutput bool
)

// BasicConfig sets the minimum LogLevel and the format to print
// it initializes the logging module
// level can be DEBUG, INFO, WARNING, ERROR, CRITICAL
// format -> see doc for examples
// file an *os.File the output will be written to
// fileOnly whether the output should not be printed but only written to the file
func BasicConfig(level LogLevel, format string, file *os.File, fileOnly bool) {
	logLevel = level
	logFormat = format
	logFile = file
	printOutput = (fileOnly == false)
}

// SetTimeFormat sets a custom time format to output if the logging format contains "{asctime}"
// the format is the same as go uses in the time package
func SetTimeFormat(format string) {
	timeFormat = format
}

// format the format string appropriately
// {prefix}
// {levelname}
// {asctime}
// {message}
func format(prefix, levelname, asctime, message string) string {
	r := strings.NewReplacer(
		"{prefix}", prefix,
		"{levelname}", levelname,
		"{asctime}", asctime,
		"{message}", message,
	)
	return r.Replace(logFormat)
}

// Debug logs debug message
func Debug(a interface{}) {
	if logLevel == 0 {
		str := format(pDebug, "DEBUG", time.Now().Format(timeFormat), fmt.Sprintf("%v", a))

		if printOutput || logFile == nil {
			fmt.Println(str)
		}
		if logFile != nil {
			logFile.WriteString(str)
		}
	}
}

// Info logs info message
func Info(a interface{}) {
	if logLevel <= 1 {
		str := format(pInfo, "INFO", time.Now().Format(timeFormat), fmt.Sprintf("%v", a))

		if printOutput || logFile == nil {
			fmt.Println(str)
		}
		if logFile != nil {
			logFile.WriteString(str)
		}
	}
}

// Warning logs warning message
func Warning(a interface{}) {
	if logLevel <= 2 {
		str := format(pWarning, "WARNING", time.Now().Format(timeFormat), fmt.Sprintf("%v", a))

		if printOutput || logFile == nil {
			fmt.Println(str)
		}
		if logFile != nil {
			logFile.WriteString(str)
		}
	}
}

// Error logs error message
func Error(a interface{}) {
	if logLevel <= 3 {
		str := format(pError, "ERROR", time.Now().Format(timeFormat), fmt.Sprintf("%v", a))

		if printOutput || logFile == nil {
			fmt.Println(str)
		}
		if logFile != nil {
			logFile.WriteString(str)
		}
	}
}

// Critical logs critical message and exits program
func Critical(a interface{}) {
	str := format(pCritical, "CRITICAL", time.Now().Format(timeFormat), fmt.Sprintf("%v", a))

	if printOutput || logFile == nil {
		fmt.Println(str)
	}
	if logFile != nil {
		logFile.WriteString(str)
		logFile.Close()
	}
	os.Exit(1)
}
