package logseal

func (lg Logseal) Debug(msg string, fields interface{}) {
	lg.Logrus.WithFields(lg.conv(fields)).Debug(msg)
}

func (lg Logseal) Info(msg string, fields interface{}) {
	lg.Logrus.WithFields(lg.conv(fields)).Info(msg)
}

func (lg Logseal) Warn(msg string, fields interface{}) {
	lg.Logrus.WithFields(lg.conv(fields)).Warn(msg)
}

func (lg Logseal) Error(msg interface{}, fields interface{}) {
	lg.Logrus.WithFields(lg.conv(fields)).Error(msg)
}

func (lg Logseal) Fatal(msg string, fields interface{}) {
	lg.Logrus.WithFields(lg.conv(fields)).Fatal(msg)
}
