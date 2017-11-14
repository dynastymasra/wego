package wego

import (
	"fmt"

	"net/http"

	"encoding/json"

	"github.com/cenkalti/backoff"
	"github.com/parnurzeal/gorequest"
)

type (
	// AccessToken An access token is a globally unique token that each Official Account must obtain before calling APIs.
	// Developers should save an access token once obtained. A minimum of 512 bytes should be reserved per access token.
	// Normally, an access token is valid for 7,200 seconds. Getting a new access token will invalidate the previous one.
	AccessToken struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int64  `json:"expires_in"`
	}

	// AccessTokenResponse handle response access token
	AccessTokenResponse struct {
		BackOff *backoff.ExponentialBackOff
		Request *gorequest.SuperAgent
	}
)

// NewAccessToken request generate new wechat access token
func NewAccessToken(appID, secret string, backOff *backoff.ExponentialBackOff) *AccessTokenResponse {
	url := BaseURL + EndpointGetAccessToken
	request := gorequest.New().Get(url).Type(gorequest.TypeJSON).Set(UserAgentHeader, UserAgent+"/"+Version).
		Query(fmt.Sprintf("grant_type=%v&appid=%v&secret=%v", "client_credential", appID, secret))

	return &AccessTokenResponse{
		BackOff: backOff,
		Request: request,
	}
}

// Commit request to wechat api
func (token *AccessTokenResponse) Commit() (*AccessToken, *http.Response, error) {
	var errs []error
	var body []byte
	res := &http.Response{}

	operation := func() error {
		res, body, errs = token.Request.EndBytes()
		if len(errs) > 0 {
			return errs[0]
		}
		return nil
	}

	if err := backoff.Retry(operation, token.BackOff); err != nil {
		return nil, MakeHTTPResponse(token.Request), err
	}
	return parseAccessToken(res, body)
}

func parseAccessToken(res *http.Response, body []byte) (*AccessToken, *http.Response, error) {
	model := struct {
		*ErrorMessage
		*AccessToken
	}{}
	if err := json.Unmarshal(body, &model); err != nil {
		return nil, res, err
	}
	if model.ErrorMessage != nil {
		return nil, res, fmt.Errorf("%v %v", model.ErrCode, model.ErrMsg)
	}
	return model.AccessToken, res, nil
}
