package lit

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"time"
)

// Error Levels that can be used to differentiate logged messages and also
// set the verbosity of logs to display.
const (

	// For unrecoverable errors where you would be unable to continue the
	// current scope of code.
	LogError int = iota

	// For non-critical errors that do not require you to abort/exit from the
	// current scope of code.
	LogWarning

	// For non-error "informational" logging.
	LogInformational

	// For any type of verbose debug specific logging.
	LogDebug
)

var (
	// ErrorPrefix is the custom prefix for errors
	ErrorPrefix string

	// WarningPrefix is the custom prefix for warnings
	WarningPrefix string

	// InformationalPrefix is the custom prefix for informational messages
	InformationalPrefix string

	// DebugPrefix is the custom prefix for debug messages
	DebugPrefix string

	// CustomPrefix enables or disables the use of custom prefixes
	CustomPrefix bool

	// Prefix is a string that is added to the start of any logged messages.
	Prefix string

	// LogLevel is the level of msgs to be logged.
	LogLevel int

	// Writer is the output io.Writer where messages are wrote to.
	Writer io.Writer

	// getPrefix returns the relevant prefix based off settings and log level
	getPrefix func(level int) string
)

// Set package defaults
func init() {
	Prefix = `LIT`
	LogLevel = 0
	Writer = os.Stderr
	getPrefix = func(level int) string {
		if !CustomPrefix {
			return fmt.Sprintf("%s%d", Prefix, level)
		} else if level == LogError {
			return ErrorPrefix
		} else if level == LogWarning {
			return WarningPrefix
		} else if level == LogInformational {
			return InformationalPrefix
		} else {
			return DebugPrefix
		}
	}
}

// Error logs a Error level message
func Error(format string, a ...interface{}) {
	if LogError > LogLevel {
		return
	}
	Custom(Writer, LogError, 2, format, a...)
}

// Warn logs a Warning level message
func Warn(format string, a ...interface{}) {
	if LogWarning > LogLevel {
		return
	}
	Custom(Writer, LogWarning, 2, format, a...)
}

// Info logs a Informational level message
func Info(format string, a ...interface{}) {
	if LogInformational > LogLevel {
		return
	}
	Custom(Writer, LogInformational, 2, format, a...)
}

// Debug logs a Debug level message
func Debug(format string, a ...interface{}) {
	if LogDebug > LogLevel {
		return
	}
	Custom(Writer, LogDebug, 2, format, a...)
}

// Custom formats and writes the provided message to the defined io.Writer output
// as long as the passed level is less than or equal to the Logger.LogLevel
//   out       : io.Writer to output message to
//   level     : Log Level of the message being logged.
//   calldepth : Distance from the caller
//   format    : Printf style message format
//   a ...     : comma separated list of values to pass (like Printf)
func Custom(out io.Writer, level int, calldepth int, format string, a ...interface{}) {

	if level > LogLevel {
		return
	}

	now := time.Now() // get this early.

	pc, file, line, _ := runtime.Caller(calldepth)

	files := strings.Split(file, "/")
	file = files[len(files)-1]

	name := runtime.FuncForPC(pc).Name()
	fns := strings.Split(name, ".")
	name = fns[len(fns)-1]

	msg := fmt.Sprintf(format, a...)

	fmt.Fprintf(out, "%s [%s] %s:%d:%s() %s\n", now.Format("2006-01-02 15:04:05"), getPrefix(level), file, line, name, msg)

}
