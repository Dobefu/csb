package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
)

func log(level byte, format string, a ...any) {
	timestamp := time.Now().Format(time.DateTime)
	var str string

	switch level {
	case 0:
		str = fmt.Sprintf(color.HiBlackString("  %s"), format)
		break
	case 1:
		str = fmt.Sprintf(color.HiCyanString("ℹ %s"), format)
		break
	case 2:
		str = fmt.Sprintf(color.YellowString("⚠ %s"), format)
		break
	case 3:
		str = fmt.Sprintf(color.HiRedString("‼ %s"), format)
		break
	case 4:
		str = fmt.Sprintf(color.RedString("‼ %s"), format)
		break
	default:
		str = fmt.Sprintf("%s", format)
		break
	}

	str = fmt.Sprintf("[%s] %s\n", timestamp, str)

	fmt.Printf(str, a...)
}

func Verbose(format string, a ...any) {
	log(0, format, a...)
}

func Info(format string, a ...any) {
	log(1, format, a...)
}

func Warning(format string, a ...any) {
	log(2, format, a...)
}

func Error(format string, a ...any) {
	log(3, format, a...)
}

func Fatal(format string, a ...any) {
	log(4, format, a...)
	os.Exit(1)
}
