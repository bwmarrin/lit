package lit

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
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
	// PrefixError is the custom prefix for errors
	PrefixError string

	// PrefixWarning is the custom prefix for warnings
	PrefixWarning string

	// PrefixInformational is the custom prefix for informational messages
	PrefixInformational string

	// PrefixDebug is the custom prefix for debug messages
	PrefixDebug string

	// Prefix is a string that is added to the start of any logged messages.
	Prefix string

	// LogLevel is the level of msgs to be logged.
	LogLevel int

	// Writer is the output io.Writer where messages are wrote to.
	Writer io.Writer
)

// Set package defaults
func init() {
	Prefix = `LIT`
	LogLevel = 0
	Writer = os.Stderr

	PrefixError = strconv.Itoa(LogError)
	PrefixDebug = strconv.Itoa(LogDebug)
	PrefixWarning = strconv.Itoa(LogWarning)
	PrefixInformational = strconv.Itoa(LogInformational)
}

func getPrefix(level int) string {
	if level == LogError {
		return PrefixError
	} else if level == LogWarning {
		return PrefixWarning
	} else if level == LogInformational {
		return PrefixInformational
	} else {
		return PrefixDebug
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

	fmt.Fprintf(out, "%s [%s%s] %s:%d:%s() %s\n", now.Format("2006-01-02 15:04:05"), Prefix, getPrefix(level), file, line, name, msg)

}
