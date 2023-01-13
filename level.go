package logging

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

	return [...]string{"[+]", "[*]", "[~]", "[!]", "[x]"}[l]
}

func (l Level) Name() string {
	if l < DEBUG || l > FATAL {
		return ""
	}

	return [...]string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}[l]
}

func (l Level) Color() Color {
	if l < DEBUG || l > FATAL {
		return ColorWhite
	}

	return [...]Color{ColorCyan, ColorWhite, ColorYellow, ColorRed, ColorBrightRed}[l]
}
