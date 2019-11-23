package kryptono

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewOrder(t *testing.T) {
	pseudoAPIKey := uuid.NewV4()
	pseudoAPISecret := "4a894c5c-8a7e-4337-bb6b-9fde16e3dddd"
	body := `{
		"order_id": "02140bef-0c98-4997-9412-9e7ca6f1cc0e",
		"account_id": "bzbf4991-ad06-44e5-908c-691fdd55da14",
		"order_symbol": "KNOW_ETH",
		"order_side": "BUY",
		"status": "open",
		"createTime": 1528277973947,
		"type": "limit",
		"order_price": "0.00001230",
		"order_size": "7777",
		"executed": "0",
		"stop_price": "0.00000000",
		"avg": "0.00001230",
		"total": "0.09565710 ETH"
	  }`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/api/v2/order/add", r.URL.String())
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assert.Equal(t, "XMLHttpRequest", r.Header.Get(HeaderXRequestedWith))

		assert.Equal(t, pseudoAPIKey.String(), r.Header.Get(HeaderAuthorization))
		assert.Equal(t, "a36b5ab0fcf4a203101f9df0dcb149077e6b6215c2495a43f2c8a977dbdf0e85", r.Header.Get(HeaderSignature))

		expectedReqBody := `{
			"order_symbol" : "KNOW_ETH",
			"order_side" : "BUY",
			"order_price" : "0.0000123",
			"order_size" : "7777",
			"type" : "LIMIT",
			"timestamp" : 1507725176599,
			"recvWindow" : 5000
		  }`
		reqBody, _ := ioutil.ReadAll(r.Body)
		equal, err := isEqualJSON(expectedReqBody, string(reqBody))
		assert.Nil(t, err, err)
		assert.True(t, equal, fmt.Sprintf("%s is not equal to %s", expectedReqBody, string(reqBody)))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(body))
	}))
	defer ts.Close()

	client, err := newClientWithURL(ts.URL, pseudoAPIKey.String(), pseudoAPISecret)
	if err != nil {
		t.Error(err.Error())
	}
	request := &NewOrderRequest{
		OrderSymbol: "KNOW_ETH",
		OrderSide:   "BUY",
		OrderPrice:  0.0000123,
		OrderSize:   7777,
		Type:        OrderTypeLimit,
		Timestamp:   1507725176599,
		RecvWindow:  5000,
	}
	resp, err := client.NewOrder(request)

	assert.NotNil(t, resp, fmt.Sprintf("error: %v", err))

	respBytes, _ := json.Marshal(resp)
	equal, _ := isEqualJSON(body, string(respBytes))
	assert.True(t, equal, fmt.Sprintf("%s is not equal to %s", body, string(respBytes)))
}

func TestTestNewOrder(t *testing.T) {
	pseudoAPIKey := uuid.NewV4()
	pseudoAPISecret := "4a894c5c-8a7e-4337-bb6b-9fde16e3dddd"
	body := `{
		"result": true
	}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/api/v2/order/test", r.URL.String())
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assert.Equal(t, "XMLHttpRequest", r.Header.Get(HeaderXRequestedWith))

		assert.Equal(t, pseudoAPIKey.String(), r.Header.Get(HeaderAuthorization))
		assert.Equal(t, "a36b5ab0fcf4a203101f9df0dcb149077e6b6215c2495a43f2c8a977dbdf0e85", r.Header.Get(HeaderSignature))

		expectedReqBody := `{
			"order_symbol" : "KNOW_ETH",
			"order_side" : "BUY",
			"order_price" : "0.0000123",
			"order_size" : "7777",
			"type" : "LIMIT",
			"timestamp" : 1507725176599,
			"recvWindow" : 5000
		  }`
		reqBody, _ := ioutil.ReadAll(r.Body)
		equal, err := isEqualJSON(expectedReqBody, string(reqBody))
		assert.Nil(t, err, err)
		assert.True(t, equal, fmt.Sprintf("%s is not equal to %s", expectedReqBody, string(reqBody)))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(body))
	}))
	defer ts.Close()

	client, err := newClientWithURL(ts.URL, pseudoAPIKey.String(), pseudoAPISecret)
	if err != nil {
		t.Error(err.Error())
	}
	request := &NewOrderRequest{
		OrderSymbol: "KNOW_ETH",
		OrderSide:   "BUY",
		OrderPrice:  0.0000123,
		OrderSize:   7777,
		Type:        OrderTypeLimit,
		Timestamp:   1507725176599,
		RecvWindow:  5000,
	}
	resp, err := client.TestNewOrder(request)

	assert.NotNil(t, resp, fmt.Sprintf("error: %v", err))

	respBytes, _ := json.Marshal(resp)
	equal, _ := isEqualJSON(body, string(respBytes))
	assert.True(t, equal, fmt.Sprintf("%s is not equal to %s", body, string(respBytes)))
}
