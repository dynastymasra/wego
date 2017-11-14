package wego

import (
	"net/http"

	"github.com/parnurzeal/gorequest"
)

type (
	// ErrorMessage to handle error response from api
	ErrorMessage struct {
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
	}
)

// MakeHTTPResponse create mock http response if request to API is error internal
func MakeHTTPResponse(agent *gorequest.SuperAgent) *http.Response {
	request, err := agent.MakeRequest()
	if err != nil {
		return &http.Response{StatusCode: http.StatusInternalServerError}
	}

	return &http.Response{
		StatusCode: http.StatusInternalServerError,
		Header:     request.Header,
		Request:    request,
	}
}
