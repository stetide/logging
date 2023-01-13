package logging

import (
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"
)

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
