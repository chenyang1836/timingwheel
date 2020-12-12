package timingwheel

import (
	"net/http"
)

func GetHttpCode(u string) (int, error) {
	r, err := http.Head(u)
	defer r.Body.Close()
	return r.StatusCode, err
}
