package kryptono

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	// Authorization Header: API-keys are passed into the Rest API via the Authorization header.
	HeaderAuthorization = "Authorization"
	// Signature Header: Signature are passed into the Rest API via the Signature header.
	HeaderSignature = "Signature"
	// XRequestedWith Header:
	HeaderXRequestedWith = "X-Requested-With"
)

type response struct {
	Header     http.Header
	Body       io.ReadCloser
	StatusCode int
	Status     string
}

func checkHTTPStatus(resp response, expected ...int) error {
	for _, e := range expected {
		if resp.StatusCode == e {
			return nil
		}
	}
	return fmt.Errorf("http response status != %+v, got %d", expected, resp.StatusCode)
}

func mergeHeaders(firstHeaders map[string]string, secondHeaders map[string]string) map[string]string {
	if secondHeaders == nil {
		return firstHeaders
	}
	if firstHeaders == nil {
		return secondHeaders
	}
	for k, v := range secondHeaders {
		if firstHeaders[k] == "" {
			firstHeaders[k] = v
		}
	}
	return firstHeaders
}

func (c *client) sendPost(url string, additionalHeaders map[string]string, body io.Reader) (*response, error) {
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(bodyBytes))
	if err != nil {
		return &response{}, fmt.Errorf("error creating POST request, %v", err)
	}

	if additionalHeaders == nil {
		additionalHeaders = make(map[string]string)
	}

	if c.auth != nil {
		h := hmac.New(sha256.New, []byte(c.auth.APISecret))
		h.Write(bodyBytes)
		signature := hex.EncodeToString(h.Sum(nil))
		additionalHeaders[HeaderSignature] = signature
	}

	return c.sendRequest(req, additionalHeaders)
}

func (c *client) sendGet(url string, additionalHeaders map[string]string, body io.Reader) (*response, error) {
	var err error
	var req *http.Request
	var bodyBytes []byte
	if body != nil {
		bodyBytes, err = ioutil.ReadAll(body)
		if err != nil {
			return nil, err
		}
		req, err = http.NewRequest("GET", url, bytes.NewReader(bodyBytes))
	} else {
		req, err = http.NewRequest("GET", url, nil)
	}

	if err != nil {
		return &response{}, fmt.Errorf("error creating GET request, %v", err)
	}

	if additionalHeaders == nil {
		additionalHeaders = make(map[string]string)
	}

	if bodyBytes != nil && c.auth != nil {
		h := hmac.New(sha256.New, []byte(c.auth.APISecret))
		h.Write(bodyBytes)
		signature := hex.EncodeToString(h.Sum(nil))
		additionalHeaders[HeaderSignature] = signature
	}

	return c.sendRequest(req, additionalHeaders)
}

func (c *client) sendDelete(url string, additionalHeaders map[string]string, body io.Reader) (*response, error) {
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", url, bytes.NewReader(bodyBytes))
	if err != nil {
		return &response{}, fmt.Errorf("error creating DELETE request, %v", err)
	}

	if additionalHeaders == nil {
		additionalHeaders = make(map[string]string)
	}

	if c.auth != nil {
		h := hmac.New(sha256.New, []byte(c.auth.APISecret))
		h.Write(bodyBytes)
		signature := hex.EncodeToString(h.Sum(nil))
		additionalHeaders[HeaderSignature] = signature
	}

	return c.sendRequest(req, additionalHeaders)
}

func (c *client) sendRequest(request *http.Request, additionalHeaders map[string]string) (*response, error) {

	for k, v := range additionalHeaders {
		request.Header.Add(k, v)
	}

	thisHeaders := map[string]string{}
	thisHeaders["Content-type"] = "application/json"
	thisHeaders[HeaderXRequestedWith] = "XMLHttpRequest"
	if c.auth != nil {
		thisHeaders[HeaderAuthorization] = c.auth.APIKey
	}
	headers := mergeHeaders(additionalHeaders, thisHeaders)
	for k, v := range headers {
		request.Header.Add(k, v)
	}
	resp, err := c.http.Do(request)
	if err != nil {
		fmt.Println(fmt.Sprintf("erro: %v", err))
		return nil, err
	}
	return &response{
		StatusCode: resp.StatusCode,
		Status:     resp.Status,
		Header:     resp.Header,
		Body:       resp.Body,
	}, nil
}
