package logseal

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type F map[string]interface{}

func (lg Logseal) conv(itf []interface{}) (msg string, fields logrus.Fields) {
	if len(itf) > 0 {
		msg = fmt.Sprintf("%s", itf[0])
	}

	if len(itf) > 1 {
		switch val := itf[1].(type) {
		case F:
			fields = logrus.Fields(val)
		case logrus.Fields:
			fields = val
		}
	}
	return
}
