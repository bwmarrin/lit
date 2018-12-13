package lit

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
)

// Error Levels that can be used to differentiate logged messages and also
// set the verbosity of logs to display.
const (

	// For unrecoverable errors where you would be unable to continue to
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
	// Prefix is a string that is added to the start of any logged messages.
	Prefix string

	// LogLevel is the level of msgs to be logged.
	LogLevel int

	// out is the output io.Writer where messages are wrote to.
	Writer io.Writer
)

// New returns a new *Logger struct with default options.
// You can (eventually) set options using the various currently non-existant
// option functions :)
func init() {
	Prefix = `LIT`
	LogLevel = 0
	Writer = os.Stderr
}

// Error logs a Error level message
func Error(format string, a ...interface{}) {
	Log(Writer, ErrorLevel, 2, format, a...)
}

// Warn logs a Warning level message
func Warn(format string, a ...interface{}) {
	Log(Writer, WarningLevel, 2, format, a...)
}

// Info logs a Informational level message
func Info(format string, a ...interface{}) {
	Log(Writer, InformationalLevel, 2, format, a...)
}

// Debug logs a Debug level message
func Debug(format string, a ...interface{}) {
	Log(Writer, DebugLevel, 2, format, a...)
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

	pc, file, line, _ := runtime.Caller(calldepth)

	files := strings.Split(file, "/")
	file = files[len(files)-1]

	name := runtime.FuncForPC(pc).Name()
	fns := strings.Split(name, ".")
	name = fns[len(fns)-1]

	msg := fmt.Sprintf(format, a...)

	fmt.Fprintf(out, "[%s%d] %s:%d:%s() %s\n", Prefix, level, file, line, name, msg)
}
