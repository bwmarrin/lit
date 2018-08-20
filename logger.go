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
	Error int = iota

	// For non-critical errors that do not require you to abort/exit from the
	// current scope of code.
	Warning

	// For non-error "informational" logging.
	Informational

	// For any type of verbose debug specific logging.
	Debug
)

// Logger is the main struct of the lit package.  It holds all logging config
// options and methods.
type Logger struct {
	// Prefix is a string that is added to the start of any logged messages.
	Prefix string

	// LogLevel is the level of msgs to be logged.
	LogLevel int

	// out is the output io.Writer where messages are wrote to.
	out io.Writer
}

// New returns a new *Logger struct with default options.
// You can (eventually) set options using the various currently non-existant
// option functions :)
func New(options ...func(*Logger)) *Logger {
	l := Logger{}
	l.Prefix = `LIT`
	l.LogLevel = 0
	l.out = os.Stderr
	return &l
}

// Error logs a Error level message
func (l Logger) Error(format string, a ...interface{}) {
	l.Log(Error, 2, format, a...)
}

// Warn logs a Warning level message
func (l Logger) Warn(format string, a ...interface{}) {
	l.Log(Warning, 2, format, a...)
}

// Info logs a Informational level message
func (l Logger) Info(format string, a ...interface{}) {
	l.Log(Informational, 2, format, a...)
}

// Debug logs a Debug level message
func (l Logger) Debug(format string, a ...interface{}) {
	l.Log(Debug, 2, format, a...)
}

// Log formats and writes the provided message to the defined io.Writer output
// as long as the passed level is less than or equal to the Logger.LogLevel
//   level     : Log Level of the message being logged.
//   calldepth : Distance from the caller
//   format    : Printf style message format
//   a ...     : comma separated list of values to pass (like Printf)
func (l Logger) Log(level int, calldepth int, format string, a ...interface{}) {

	if level > l.LogLevel {
		return
	}

	pc, file, line, _ := runtime.Caller(calldepth)

	files := strings.Split(file, "/")
	file = files[len(files)-1]

	name := runtime.FuncForPC(pc).Name()
	fns := strings.Split(name, ".")
	name = fns[len(fns)-1]

	msg := fmt.Sprintf(format, a...)

	fmt.Fprintf(l.out, "[%s%d] %s:%d:%s() %s\n", l.Prefix, level, file, line, name, msg)
}
