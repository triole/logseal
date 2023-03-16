package logseal

import (
	"errors"
	"testing"
)

func TestLogging(t *testing.T) {
	lg := Init("trace")
	fireLogs(lg)

	lg = Init("trace", "/dev/null")
	fireLogs(lg)

	lg = Init("trace", "/dev/stdout")
	fireLogs(lg)

	lg = Init("trace", "/tmp/tempfile")
	fireLogs(lg)
}

func fireLogs(lg Logseal) {
	lg.Trace("trace message", F{"lev": "trace"})
	lg.Debug("debug message", F{"lev": "debug", "status": "important"})
	lg.Info(
		"info message",
		F{"lev": "info", "status": "not important", "advise": "don't read"},
	)

	err := errors.New("this is an error")

	// note that iferr methods look up field names
	// if there is either an 'error' or an 'err' field containing
	// a value, the logging gets triggered
	lg.IfErrWarn("warning because of an error", F{"error": err})
	lg.IfErrWarn("not shown, because there is no error", F{"noerror": err})
	lg.IfErrError("error log because of an error", F{"error": err})
	lg.IfErrFatal("not shown, because there is no error", F{"noerror": err})
}
