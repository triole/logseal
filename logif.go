package logseal

func (lg Logseal) IfErrError(itf ...interface{}) {
	msg, fields := lg.conv(itf)
	for key, val := range fields {
		if (key == "error" || key == "err") && val != nil {
			lg.Error(msg, fields)
		}
	}
}

func (lg Logseal) IfErrWarn(itf ...interface{}) {
	msg, fields := lg.conv(itf)
	for key, val := range fields {
		if (key == "error" || key == "err") && val != nil {
			lg.Warn(msg, fields)
		}
	}
}

func (lg Logseal) IfErrFatal(itf ...interface{}) {
	msg, fields := lg.conv(itf)
	for key, val := range fields {
		if (key == "error" || key == "err") && val != nil {
			lg.Fatal(msg, fields)
		}
	}
}
