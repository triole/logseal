package logseal

func (lg Logseal) Trace(itf ...interface{}) {
	msg, fields := lg.conv(itf)
	lg.Logrus.WithFields(fields).Trace(msg)
}

func (lg Logseal) Debug(itf ...interface{}) {
	msg, fields := lg.conv(itf)
	lg.Logrus.WithFields(fields).Debug(msg)
}

func (lg Logseal) Info(itf ...interface{}) {
	msg, fields := lg.conv(itf)
	lg.Logrus.WithFields(fields).Info(msg)
}

func (lg Logseal) Warn(itf ...interface{}) {
	msg, fields := lg.conv(itf)
	lg.Logrus.WithFields(fields).Warn(msg)
}

func (lg Logseal) Error(itf ...interface{}) {
	msg, fields := lg.conv(itf)
	lg.Logrus.WithFields(fields).Error(msg)
}

func (lg Logseal) Fatal(itf ...interface{}) {
	msg, fields := lg.conv(itf)
	lg.Logrus.WithFields(fields).Fatal(msg)
}
