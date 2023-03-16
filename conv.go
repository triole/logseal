package logseal

import (
	"github.com/sirupsen/logrus"
)

type F map[string]interface{}

func (lg Logseal) conv(itf interface{}) logrus.Fields {
	switch val := itf.(type) {
	case logrus.Fields:
		return val
	case F:
		return lg.toLogrusFields(val)
	default:
		return logrus.Fields{}
	}
}

func (lg Logseal) toLogrusFields(fields F) logrus.Fields {
	if fields != nil {
		return logrus.Fields(fields)
	}
	return logrus.Fields{}
}
