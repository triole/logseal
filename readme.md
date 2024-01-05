# Logseal

Logseal is a wrapper around [logrus](https://github.com/sirupsen/logrus) that provides a few functions that make life a little easier.

## Usage

Four init parameters are supported.

| parameter             | type   | default     |
|-----------------------|--------|-------------|
| log level             | string | info        |
| log file              | string | /dev/stdout |
| no colours, bool      | bool   | false       |
| enable JSON log, bool | bool   | false       |

```go
import (
	"github.com/triole/logseal"
)

lg = logseal.Init("debug", nil, true, false)
```
