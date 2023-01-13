package logging

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
