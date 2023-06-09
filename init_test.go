package logseal

import "testing"

func TestInit(t *testing.T) {
	_ = Init()
	_ = Init("info")
	_ = Init("info", "/dev/stdout")
	_ = Init("info", "/tmp/tempfile")
	_ = Init("info", "/dev/stdout", false)
	_ = Init("info", "/dev/stdout", false, false)
}
