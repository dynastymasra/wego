package wego_test

import (
	"net/http"
	"testing"
	"wego"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestNewAccessTokenSuccess(t *testing.T) {
	gock.New(wego.BaseURL).Get(wego.EndpointGetAccessToken).Reply(http.StatusOK).JSON(`{
		"access_token": "3ridcJbJ5LcsdCiA4vB3RTrP46rsKRkSsHCwQDITpo5EOk-PPxLbaY0ZVXHk3MWSVLcNec77gcIgwTsGTAMREDjzOU__sNWLkTQ1seqjnCTHhPJBVzPV-Q5fvcwgOJw1FXLhABAASG",
		"expires_in": 7200
	}`)
	defer gock.Off()

	model, res, err := wego.NewAccessToken("appId", "secret", wego.NewBackOff(10, -1)).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestNewAccessTokenError(t *testing.T) {
	gock.New(wego.BaseURL).Head(wego.EndpointGetAccessToken).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	model, res, err := wego.NewAccessToken("appId", "secret", wego.NewBackOff(10, -1)).Commit()

	assert.Nil(t, model)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Error(t, err)
}

func TestNewAccessTokenFailedUnmarshal(t *testing.T) {
	gock.New(wego.BaseURL).Get(wego.EndpointGetAccessToken).Reply(http.StatusOK).XML("")
	defer gock.Off()

	model, res, err := wego.NewAccessToken("appId", "secret", wego.NewBackOff(10, -1)).Commit()

	assert.Nil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Error(t, err)
}

func TestNewAccessTokenFailed(t *testing.T) {
	gock.New(wego.BaseURL).Get(wego.EndpointGetAccessToken).Reply(http.StatusOK).JSON(`{
		"errcode":40013,
	  	"errmsg":"invalid appid"
	}`)
	defer gock.Off()

	model, res, err := wego.NewAccessToken("appId", "secret", wego.NewBackOff(10, -1)).Commit()

	assert.Nil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Error(t, err)
}
