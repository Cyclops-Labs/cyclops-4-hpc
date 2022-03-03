//
// A logging framework for use within the Cyclops Accounting
// and Billing framework.
//
// (C) Cyclops Labs 2019
//

package logging

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
)

// higher log level means more logging...
const (
	ERROR = iota
	WARNING
	INFO
	DEBUG
	TRACE
)

var logLevelStringMap = map[string]int{
	"error":   ERROR,
	"warning": WARNING,
	"info":    INFO,
	"debug":   DEBUG,
	"trace":   TRACE,
}

type logOutput interface {
	Printf(format string, a ...interface{})
	// Println(format string, a ...interface{})
	init(m *io.Writer, s string, ll int)
}

type logStream struct {
	s        *log.Logger
	logLevel int
}

var (
	Trace         logStream
	Debug         logStream
	Info          logStream
	Warning       logStream
	Error         logStream
	DefaultLogger logStream
	LogLevel      = WARNING
)

func (l *logStream) Printf(format string, a ...interface{}) {

	// describe(a)
	if LogLevel < l.logLevel {

		return

	}

	// if no specific logger is configured for this logstrea, write to stdout...
	logMessage := fmt.Sprintf(format, a...)

	// Caller(1) is one level up the call stack - 0 identifies the caller of
	// Caller which is this function...
	callingStr := "[Unknown calling function]"

	_, fn, lineNum, callerValid := runtime.Caller(1)

	if callerValid {

		callingStr = fn + ":" + strconv.Itoa(lineNum)

	}

	if l.s == nil {

		fmt.Print(callingStr, " ", logMessage)

		return

	}

	l.s.Print(callingStr, " ", logMessage)
	// Printf(format, a...)

}

func (l *logStream) init(m io.Writer, s string, ll int) {

	l.s = log.New(m, s, log.Ldate|log.Ltime)

	l.logLevel = ll

}

// function to parse string representation of log level into integer - if there
// is no match of the string, it will return 0 which maps to error only
// logging...perhaps this could be made more explicit
func determineLogLevel(lstr string) (returnVal int, err error) {

	var validLogLevel bool

	lstrLowerCase := strings.ToLower(lstr)

	if returnVal, validLogLevel = logLevelStringMap[lstrLowerCase]; validLogLevel == false {

		err = errors.New("Invalid Log Level")

	}

	return

}

// InitLogger initializes the logging system
func InitLogger(logFile string, logLevelStr string, logToConsole bool) (err error) {

	var file *os.File
	var writters []io.Writer

	if LogLevel, err = determineLogLevel(logLevelStr); err != nil {

		log.Println("Unable to determine the log level:", logLevelStr, ":", err)

		return

	}

	if logFile != "" {

		if file, err = os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666); err != nil {

			log.Println("Failed to open log file:", logFile, ":", err)

			return

		}

		writters = append(writters, file)

	}

	if logToConsole {

		writters = append(writters, os.Stdout)

	}

	multi := io.MultiWriter(writters...)

	Trace.init(multi, "   TRACE: ", TRACE)
	Debug.init(multi, "  DEBUG: ", DEBUG)
	Info.init(multi, "   INFO: ", INFO)
	Warning.init(multi, "WARNING: ", WARNING)
	Error.init(multi, "  ERROR: ", ERROR)

	return

}
