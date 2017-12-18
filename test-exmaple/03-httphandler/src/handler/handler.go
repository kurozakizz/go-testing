package handler

import (
	"net/http"
)

type CustomHandler interface {
	Get(url string) (resp *http.Response, err error)
}

// GetHTTPStatus return http status after call url
func GetHTTPStatus(c CustomHandler, EndPointURL string) int {
	response, err := c.Get(EndPointURL)
	if err != nil {
		return http.StatusServiceUnavailable
	}
	defer response.Body.Close()

	return response.StatusCode
}
