package bingcontent

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/oauth2"
)

type AuthHeader struct {
	Action              string
	ApplicationToken    string
	AuthenticationToken string
	CustomerAccountId   int64
	CustomerId          int64
	DeveloperToken      string
	UserName            string
	Password            string
}

type RequestHeader struct {
	AuthenticationToken string
	CustomerAccountId   string
	CustomerId          string
	DeveloperToken      string
	Password            string
	Username            string
}

type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

var debug = os.Getenv("BING_SDK_DEBUG")
var baseUrl = "https://content.api.bingads.microsoft.com/shopping/v9.1"

func (b *Session) SendRequest(body interface{}, endpoint string) ([]byte, error) {
	header := RequestHeader{
		CustomerAccountId: b.AccountId,
		CustomerId:        b.CustomerId,
		DeveloperToken:    b.DeveloperToken,
	}
	if b.TokenSource != nil {
		token, err := b.TokenSource.Token()
		if err != nil {
			return nil, err
		}
		header.AuthenticationToken = token.AccessToken
	} else {
		header.Username = b.Username
		header.Password = b.Password
	}

	envelope := RequestEnvelope{
		EnvNS:  "http://www.w3.org/2001/XMLSchema-instance",
		EnvSS:  "http://schemas.xmlsoap.org/soap/envelope/",
		Header: header,
		Body: RequestBody{
			Body: body,
		},
	}

	req, err := xml.MarshalIndent(envelope, "", "  ")

	if err != nil {
		return nil, err
	}

	httpRequest, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(req))

	if err != nil {
		return nil, err
	}

	httpRequest.Header.Add("Content-Type", "text/xml; charset=utf-8")

	response, err := b.HTTPClient.Do(httpRequest)

	if err != nil {
		return nil, err
	}

	raw, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if debug != "" {
		fmt.Println(string(req))
	}
	//fmt.Println(string(raw))

	res := SoapResponseEnvelope{}

	err = xml.Unmarshal(raw, &res)
	if err != nil {
		return nil, err
	}

	if debug != "" {
		fmt.Println(">>>")
		fmt.Println(string(res.Body.OperationResponse))
		fmt.Println(">>>")
	}

	switch response.StatusCode {
	case 400, 401, 403, 405, 500:
		fault := Fault{}
		err = xml.Unmarshal(res.Body.OperationResponse, &fault)
		if err != nil {
			return res.Body.OperationResponse, err
		}
		for _, e := range fault.Detail.Errors.AdApiErrors {
			switch e.ErrorCode {
			case "AuthenticationTokenExpired", "InvalidCredentials", "InternalError", "CallRateExceeded":
				return res.Body.OperationResponse, &baseError{
					code:    e.ErrorCode,
					origErr: &fault.Detail.Errors,
				}
			}
		}
		return res.Body.OperationResponse, &fault.Detail.Errors
	}

	return res.Body.OperationResponse, err
}


