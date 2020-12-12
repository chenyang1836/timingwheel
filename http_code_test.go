package timingwheel

import "testing"

func TestGetHttpCode(t *testing.T) {
	t.Log(GetHttpCode("http://httpbin.org/get"))
}
