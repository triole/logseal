package logseal

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

type Logseal struct {
	Logrus    *logrus.Logger
	LogToFile bool
}

func Init(itf ...interface{}) (lg Logseal) {
	logLevel := "info"
	if len(itf) > 0 {
		logLevel = itf[0].(string)
	}

	logFile := "/dev/stdout"
	if len(itf) > 1 {
		logFile = itf[1].(string)
	}

	noColours := false
	if len(itf) > 2 {
		noColours = itf[2].(bool)
	}

	JSONLog := false
	if len(itf) > 3 {
		JSONLog = itf[3].(bool)
	}

	timeStampFormat := "2006-01-02 15:04:05.000 MST"
	lg.Logrus = logrus.New()

	if JSONLog {
		lg.Logrus.SetFormatter(&logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime:  "date",
				logrus.FieldKeyLevel: "level",
				logrus.FieldKeyMsg:   "msg",
			},
			TimestampFormat:   timeStampFormat,
			PrettyPrint:       false,
			DisableHTMLEscape: false,
		})
	} else {
		form := &logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: timeStampFormat,
			DisableQuote:    true,
			PadLevelText:    true,
			ForceColors:     true,
		}
		if noColours {
			form.ForceColors = false
			form.DisableColors = true
		}
		lg.Logrus.SetFormatter(form)
	}

	if logFile != "/dev/stdout" {
		lg.LogToFile = true
		openLogFile, err := os.OpenFile(
			logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644,
		)
		lg.IfErrFatal(
			"Can not open log file",
			F{
				"logfile": logFile,
				"error":   err.Error(),
			},
		)

		lg.Logrus.SetOutput(openLogFile)
	}
	lg.setLevel(logLevel)

	return lg
}

func (lg *Logseal) setLevel(level string) {
	logLevels := map[string]logrus.Level{
		"trace": logrus.TraceLevel,
		"debug": logrus.DebugLevel,
		"info":  logrus.InfoLevel,
		"warn":  logrus.WarnLevel,
		"error": logrus.ErrorLevel,
	}

	if val, ok := logLevels[strings.ToLower(level)]; ok {
		lg.Logrus.SetLevel(val)
	} else {
		lg.setLevel("info")
	}
}
