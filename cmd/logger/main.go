package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
)

var (
	LOG_VERBOSE uint8 = 0
	LOG_INFO    uint8 = 1
	LOG_WARNING uint8 = 2
	LOG_ERROR   uint8 = 3
	LOG_FATAL   uint8 = 4
)

var LOGLEVEL uint8 = LOG_INFO

func log(level byte, format string, a ...any) {
	if level < LOGLEVEL {
		return
	}

	timestamp := time.Now().Format(time.DateTime)
	var str string

	switch level {
	case LOG_VERBOSE:
		str = fmt.Sprintf(color.HiBlackString("  %s"), format)
		break
	case LOG_INFO:
		str = fmt.Sprintf(color.HiCyanString("ℹ %s"), format)
		break
	case LOG_WARNING:
		str = fmt.Sprintf(color.YellowString("⚠ %s"), format)
		break
	case LOG_ERROR:
		str = fmt.Sprintf(color.HiRedString("‼ %s"), format)
		break
	case LOG_FATAL:
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
	log(LOG_VERBOSE, format, a...)
}

func Info(format string, a ...any) {
	log(LOG_INFO, format, a...)
}

func Warning(format string, a ...any) {
	log(LOG_WARNING, format, a...)
}

func Error(format string, a ...any) {
	log(LOG_ERROR, format, a...)
}

func Fatal(format string, a ...any) {
	log(LOG_FATAL, format, a...)
	os.Exit(1)
}

func SetLogLevel(level uint8) {
	LOGLEVEL = level
}
