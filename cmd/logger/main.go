package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
)

var (
	LOG_VERBOSE byte = 0
	LOG_INFO    byte = 1
	LOG_WARNING byte = 2
	LOG_ERROR   byte = 3
	LOG_FATAL   byte = 4
)

var LOGLEVEL byte = LOG_INFO
var EXIT_ON_FATAL = true

func logMessage(level byte, format string, a ...any) string {
	if level < LOGLEVEL {
		return ""
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
	output := fmt.Sprintf(str, a...)

	fmt.Print(output)
	return output
}

func Verbose(format string, a ...any) string {
	return logMessage(LOG_VERBOSE, format, a...)
}

func Info(format string, a ...any) string {
	return logMessage(LOG_INFO, format, a...)
}

func Warning(format string, a ...any) string {
	return logMessage(LOG_WARNING, format, a...)
}

func Error(format string, a ...any) string {
	return logMessage(LOG_ERROR, format, a...)
}

func Fatal(format string, a ...any) string {
	output := logMessage(LOG_FATAL, format, a...)

	if EXIT_ON_FATAL {
		os.Exit(1)
	}

	return output
}

func SetLogLevel(level byte) {
	LOGLEVEL = level
}

func SetExitOnFatal(value bool) {
	EXIT_ON_FATAL = value
}
