package color

import "fmt"

const escape = "\x1b"

type Color int

const (
	Reset Color = iota
	Bold
	Dim
	Italic
	Underline
)

const (
	FgBlack Color = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
)

const (
	FgDarkGray Color = iota + 90
	FgLightRed
	FgLightGreen
	FgLightYellow
	FgLightBlue
	FgLightMagenta
	FgLightCyan
	FgWhite
)

const (
	BgDefault Color = iota + 49
	BgBlack
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgLightGray
)

const (
	BgDarkGray Color = iota + 100
	BgLightRed
	BgLightGreen
	BgLightYellow
	BgLightBlue
	BgLightMagenta
	BgLightCyan
	BgWhite
)

func SprintColor(fg Color, bg Color, message string) string {
	return fmt.Sprintf("%s[%d;%dm%s%s[0;0m", escape, fg, bg, message, escape)
}

func PrintColor(fg Color, bg Color, message string) {
	fmt.Print(SprintColor(fg, bg, message))
}
