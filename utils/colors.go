package utils

import "regexp"

// ANSI color codes for terminal output
const (
	Reset  = "\033[0m"
	Bold   = "\033[1m"
	Under  = "\033[4m"
	Italic = "\033[3m"

	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"

	HighlightDest   = Red + Bold
	HighlightDate   = Green + Bold
	HighlightTime   = Blue
	HighlightOffset = Cyan + Italic
)

// RemoveANSI removes color codes for nofancy mode
func RemoveANSI(input string) string {
	// Match all escape codes like "\x1b[32m"
	re := regexp.MustCompile(`\x1b\[[0-9]*m`)
	return re.ReplaceAllString(input, "")
}
