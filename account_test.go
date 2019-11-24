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
		Type:        "LIMIT",
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
		Type:        "LIMIT",
		Timestamp:   1507725176599,
		RecvWindow:  5000,
	}
	resp, err := client.TestNewOrder(request)

	assert.NotNil(t, resp, fmt.Sprintf("error: %v", err))

	respBytes, _ := json.Marshal(resp)
	equal, _ := isEqualJSON(body, string(respBytes))
	assert.True(t, equal, fmt.Sprintf("%s is not equal to %s", body, string(respBytes)))
}

func TestOrderDetail(t *testing.T) {
	pseudoAPIKey := uuid.NewV4()
	pseudoAPISecret := "4a894c5c-8a7e-4337-bb6b-9fde16e3dddd"
	body := `{
		"order_id": "0e3f05e0-912c-4957-9322-d1a34ef6e312",
		"account_id": "14ce3690-4e86-4f69-8412-b9fd88535f8z",
		"order_symbol": "KNOW_BTC",
		"order_side": "SELL",
		"status": "open",
		"createTime": 1429514463266,
		"type": "limit",
		"order_price": "0.00001234",
		"order_size": "1000",
		"executed": "0",
		"stop_price": "0.00456",
		"avg": "0.00001234",
		"total": "0.1234 BTC"
	  }`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/api/v2/order/details", r.URL.String())
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assert.Equal(t, "XMLHttpRequest", r.Header.Get(HeaderXRequestedWith))

		assert.Equal(t, pseudoAPIKey.String(), r.Header.Get(HeaderAuthorization))
		assert.Equal(t, "a111fb0a7c6260451da568116506cb8639e6c9de08b791bb27904c4894faf39a", r.Header.Get(HeaderSignature))

		expectedReqBody := `{
			"order_id" : "0e3f05e0-912c-4957-9322-d1a34ef6e312",
			"timestamp" : 1429514463299,
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
	request := &OrderDetailRequest{
		OrderID:    "0e3f05e0-912c-4957-9322-d1a34ef6e312",
		Timestamp:  1429514463299,
		RecvWindow: 5000,
	}
	resp, err := client.OrderDetail(request)

	assert.NotNil(t, resp, fmt.Sprintf("error: %v", err))

	respBytes, _ := json.Marshal(resp)
	equal, _ := isEqualJSON(body, string(respBytes))
	assert.True(t, equal, fmt.Sprintf("%s is not equal to %s", body, string(respBytes)))
}

func TestCancelOrder(t *testing.T) {
	pseudoAPIKey := uuid.NewV4()
	pseudoAPISecret := "4a894c5c-8a7e-4337-bb6b-9fde16e3dddd"
	body := `{
		"order_id" : "02140bef-0c98-4997-9412-9e7ca6f1cc0e",
		"order_symbol" : "KNOW_ETH"
	  }`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "DELETE", r.Method)
		assert.Equal(t, "/api/v2/order/cancel", r.URL.String())
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assert.Equal(t, "XMLHttpRequest", r.Header.Get(HeaderXRequestedWith))

		assert.Equal(t, pseudoAPIKey.String(), r.Header.Get(HeaderAuthorization))
		assert.Equal(t, "213235fefba2a1c791acff50ad0ad0322b8bfc1ef93efcc4fe1dd6a4504f700f", r.Header.Get(HeaderSignature))

		expectedReqBody := `{
			"order_id" : "02140bef-0c98-4997-9412-9e7ca6f1cc0e",
			"order_symbol" : "KNOW_ETH",
			"timestamp" : 1429514463299,
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
	request := &CancelOrderRequest{
		OrderID:     "02140bef-0c98-4997-9412-9e7ca6f1cc0e",
		OrderSymbol: "KNOW_ETH",
		Timestamp:   1429514463299,
		RecvWindow:  5000,
	}
	resp, err := client.CancelOrder(request)

	assert.NotNil(t, resp, fmt.Sprintf("error: %v", err))

	respBytes, _ := json.Marshal(resp)
	equal, _ := isEqualJSON(body, string(respBytes))
	assert.True(t, equal, fmt.Sprintf("%s is not equal to %s", body, string(respBytes)))
}

func TestTradeDetails(t *testing.T) {
	pseudoAPIKey := uuid.NewV4()
	pseudoAPISecret := "4a894c5c-8a7e-4337-bb6b-9fde16e3dddd"
	body := `[
		{
		  "hex_id": "5b31eb38892faf4c3529ba89",
		  "symbol": "KNOW_BTC",
		  "order_id": "08098511-ae65-452b-9a84-5b79a5160b5f",
		  "order_side": "SELL",
		  "price": "0.00007677",
		  "quantity": "749",
		  "fee": "0.37449524 KNOW",
		  "total": "0.05750073 BTC",
		  "timestamp": 1529998122350
		}
	  ]`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/api/v2/order/trade-detail", r.URL.String())
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assert.Equal(t, "XMLHttpRequest", r.Header.Get(HeaderXRequestedWith))

		assert.Equal(t, pseudoAPIKey.String(), r.Header.Get(HeaderAuthorization))
		assert.Equal(t, "f07baa4c1763d5258bb9e0287c855fdef4b5cb44d83ab7cc37742190da805499", r.Header.Get(HeaderSignature))

		expectedReqBody := `{
			"order_id" : "08098511-ae65-452b-9a84-5b79a5160b5f",
			"timestamp" : 1429514463299,
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
	request := &TradeDetailsRequest{
		OrderID:    "08098511-ae65-452b-9a84-5b79a5160b5f",
		Timestamp:  1429514463299,
		RecvWindow: 5000,
	}
	resp, err := client.TradeDetails(request)

	assert.NotNil(t, resp, fmt.Sprintf("error: %v", err))

	respBytes, _ := json.Marshal(resp)
	equal, _ := isEqualJSON(body, string(respBytes))
	assert.True(t, equal, fmt.Sprintf("%s is not equal to %s", body, string(respBytes)))
}

func TestOpenOrders(t *testing.T) {
	pseudoAPIKey := uuid.NewV4()
	pseudoAPISecret := "4a894c5c-8a7e-4337-bb6b-9fde16e3dddd"
	body := `{
		"total": 5,
		"list": [
		  {
			"order_id": "02140bef-0c98-4997-9412-9e7ca6f1cc0e",
			"account_id": "bzbf4991-ad06-44e5-908c-691fdd55da14",
			"order_symbol": "KNOW_BTC",
			"order_side": "BUY",
			"status": "open",
			"createTime": 1528277973947,
			"type": "limit",
			"order_price": "0.0000123",
			"order_size": "7777",
			"executed": "0",
			"stop_price": "0",
			"avg": "0.0000123",
			"total": "0.09565710 BTC"
		  }
		]
	  }`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/api/v2/order/list/open", r.URL.String())
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assert.Equal(t, "XMLHttpRequest", r.Header.Get(HeaderXRequestedWith))

		assert.Equal(t, pseudoAPIKey.String(), r.Header.Get(HeaderAuthorization))
		assert.Equal(t, "67f6ef3576c78eae3059ceba1edb825e751b3d48205e30f8c6bac4760fc216b6", r.Header.Get(HeaderSignature))

		expectedReqBody := `{
			"limit" : 10,
			"page" : 0,
			"symbol" : "KNOW_BTC",
			"timestamp" : 1429514463299,
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
	request := &OpenOrdersRequest{
		Limit:      10,
		Page:       0,
		Symbol:     "KNOW_BTC",
		Timestamp:  1429514463299,
		RecvWindow: 5000,
	}
	resp, err := client.OpenOrders(request)

	assert.NotNil(t, resp, fmt.Sprintf("error: %v", err))

	respBytes, _ := json.Marshal(resp)
	equal, _ := isEqualJSON(body, string(respBytes))
	assert.True(t, equal, fmt.Sprintf("%s is not equal to %s", body, string(respBytes)))
}

func TestCompletedOrders(t *testing.T) {
	pseudoAPIKey := uuid.NewV4()
	pseudoAPISecret := "4a894c5c-8a7e-4337-bb6b-9fde16e3dddd"
	body := `{
		"total": 5,
		"list": [
		  {
			"order_id": "02140bef-0c98-4997-9412-9e7ca6f1cc0e",
			"account_id": "bzbf4991-ad06-44e5-908c-691fdd55da14",
			"order_symbol": "KNOW_BTC",
			"order_side": "BUY",
			"status": "open",
			"createTime": 1528277973947,
			"type": "limit",
			"order_price": "0.0000123",
			"order_size": "7777",
			"executed": "0",
			"stop_price": "0",
			"avg": "0.0000123",
			"total": "0.09565710 BTC"
		  }
		]
	  }`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/api/v2/order/list/completed", r.URL.String())
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assert.Equal(t, "XMLHttpRequest", r.Header.Get(HeaderXRequestedWith))

		assert.Equal(t, pseudoAPIKey.String(), r.Header.Get(HeaderAuthorization))
		assert.Equal(t, "67f6ef3576c78eae3059ceba1edb825e751b3d48205e30f8c6bac4760fc216b6", r.Header.Get(HeaderSignature))

		expectedReqBody := `{
			"limit" : 10,
			"page" : 0,
			"symbol" : "KNOW_BTC",
			"timestamp" : 1429514463299,
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
	request := &CompletedOrdersRequest{
		Limit:      10,
		Page:       0,
		Symbol:     "KNOW_BTC",
		Timestamp:  1429514463299,
		RecvWindow: 5000,
	}
	resp, err := client.CompletedOrders(request)

	assert.NotNil(t, resp, fmt.Sprintf("error: %v", err))

	respBytes, _ := json.Marshal(resp)
	equal, _ := isEqualJSON(body, string(respBytes))
	assert.True(t, equal, fmt.Sprintf("%s is not equal to %s", body, string(respBytes)))
}

func TestAllOrders(t *testing.T) {
	pseudoAPIKey := uuid.NewV4()
	pseudoAPISecret := "4a894c5c-8a7e-4337-bb6b-9fde16e3dddd"
	body := `[
		{
		  "order_id": "02140bef-0c98-4997-9412-9e7ca6f1cc0e",
		  "account_id": "bzbf4991-ad06-44e5-908c-691fdd55da14",
		  "order_symbol": "KNOW_BTC",
		  "order_side": "BUY",
		  "status": "open",
		  "createTime": 1528277973947,
		  "type": "limit",
		  "order_price": "0.00001230",
		  "order_size": "7777",
		  "executed": "0",
		  "stop_price": "0.00000000",
		  "avg": "0.00001230",
		  "total": "0.09565710 BTC"
		}
	  ]`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/api/v2/order/list/all", r.URL.String())
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assert.Equal(t, "XMLHttpRequest", r.Header.Get(HeaderXRequestedWith))

		assert.Equal(t, pseudoAPIKey.String(), r.Header.Get(HeaderAuthorization))
		assert.Equal(t, "1d51840d4372ad6833da2205b17a74899a976a9799cec61dd6bf67d8b6b070ce", r.Header.Get(HeaderSignature))

		expectedReqBody := `{
			"symbol" : "KNOW_BTC",
			"limit" : 50,
			"timestamp" : 1530682938651,
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
	request := &AllOrdersRequest{
		Symbol:     "KNOW_BTC",
		Limit:      50,
		Timestamp:  1530682938651,
		RecvWindow: 5000,
	}
	resp, err := client.AllOrders(request)

	assert.NotNil(t, resp, fmt.Sprintf("error: %v", err))

	respBytes, _ := json.Marshal(resp)
	equal, _ := isEqualJSON(body, string(respBytes))
	assert.True(t, equal, fmt.Sprintf("%s is not equal to %s", body, string(respBytes)))
}

func TestTradeList(t *testing.T) {
	pseudoAPIKey := uuid.NewV4()
	pseudoAPISecret := "4a894c5c-8a7e-4337-bb6b-9fde16e3dddd"
	body := `[
		{
		  "hex_id": "5b31eb38892faf4c3529ba89",
		  "symbol": "KNOW_BTC",
		  "order_id": "08098511-ae65-452b-9a84-5b79a5160b5f",
		  "order_side": "SELL",
		  "price": "0.00007677",
		  "quantity": "749",
		  "fee": "0.37449524 KNOW",
		  "total": "0.05750073 BTC",
		  "timestamp": 1529998122350
		}
	  ]`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/api/v2/order/list/trades", r.URL.String())
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assert.Equal(t, "XMLHttpRequest", r.Header.Get(HeaderXRequestedWith))

		assert.Equal(t, pseudoAPIKey.String(), r.Header.Get(HeaderAuthorization))
		assert.Equal(t, "1d51840d4372ad6833da2205b17a74899a976a9799cec61dd6bf67d8b6b070ce", r.Header.Get(HeaderSignature))

		expectedReqBody := `{
			"symbol" : "KNOW_BTC",
			"limit" : 50,
			"timestamp" : 1530682938651,
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
	request := &TradeListRequest{
		Symbol:     "KNOW_BTC",
		Limit:      50,
		Timestamp:  1530682938651,
		RecvWindow: 5000,
	}
	resp, err := client.TradeList(request)

	assert.NotNil(t, resp, fmt.Sprintf("error: %v", err))

	respBytes, _ := json.Marshal(resp)
	equal, _ := isEqualJSON(body, string(respBytes))
	assert.True(t, equal, fmt.Sprintf("%s is not equal to %s", body, string(respBytes)))
}

func TestAccountInformation(t *testing.T) {
	pseudoAPIKey := uuid.NewV4()
	pseudoAPISecret := "4a894c5c-8a7e-4337-bb6b-9fde16e3dddd"
	body := `{
		"account_id": "5377f2e2-4b0e-4b15-be17-28092ae0c346",
		"email": "email@email.com",
		"phone": null,
		"enable_google_2fa": true,
		"status": "offline",
		"create_at": 1524567654822,
		"nick_name": "Nickname 1",
		"chat_id": "xxxx@kryptono.exchange",
		"chat_password": "xxxxxxxxxxx",
		"banks": [],
		"country": "US",
		"language": "en",
		"kyc_status": null,
		"kyc_level": "level1",
		"last_login_history": {
		  "id": {
			"timestamp": 1528199468,
			"machineIdentifier": 8990639,
			"processIdentifier": 20156,
			"counter": 7772354,
			"time": 1528199468000,
			"date": 1528199468000,
			"timeSecond": 1528199468
		  },
		  "account_id": "5377f2e2-4b0e-4b15-be17-28092ae0c346",
		  "nick_name": "Nickname 1",
		  "email": "email@email.com",
		  "ip_address": "xxx.xxx.xxx.xxx",
		  "login_at": 1528199468073,
		  "os_name": "Mac OS X",
		  "browser_name": "Chrome",
		  "country": "Country",
		  "city": "City",
		  "sentEmail": true
		},
		"commission_status": true,
		"account_kyc": null,
		"kyc_reject_infos": [],
		"allow_order": 1,
		"disable_withdraw": 0,
		"referral_id": "XXXXXX",
		"favorite_pairs": [
		  "KNOW_ETH"
		],
		"chat_server": "wss://xxx.kryptono.exchange:xxxx/ws",
		"exchange_fee": {
		  "standard_fee": "0.1",
		  "know_fee": "0.05"
		}
	  }`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/api/v2/account/details", r.URL.String())
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assert.Equal(t, "XMLHttpRequest", r.Header.Get(HeaderXRequestedWith))

		assert.Equal(t, pseudoAPIKey.String(), r.Header.Get(HeaderAuthorization))
		assert.Equal(t, "f6f8463e8c41574ccb2bfc46938868b78b0b6a1a0cb75cf58a92eba3d6535b9f", r.Header.Get(HeaderSignature))

		expectedReqBody := `{
			"timestamp" : 1530682938651,
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
	request := &AccountInformationRequest{
		Timestamp:  1530682938651,
		RecvWindow: 5000,
	}
	resp, err := client.AccountInformation(request)

	assert.NotNil(t, resp, fmt.Sprintf("error: %v", err))

	respBytes, _ := json.Marshal(resp)
	equal, _ := isEqualJSON(body, string(respBytes))
	assert.True(t, equal, fmt.Sprintf("%s is not equal to %s", body, string(respBytes)))
}

func TestAccountBalances(t *testing.T) {
	pseudoAPIKey := uuid.NewV4()
	pseudoAPISecret := "4a894c5c-8a7e-4337-bb6b-9fde16e3dddd"
	body := `[
		{
		  "currency_code": "BTC",
		  "address": "2MxctvXExQofAVqakPfBjKqVipfwTqwyQyF",
		  "total": "1000.00275",
		  "available": "994.5022",
		  "in_order": "5.50055"
		}
	  ]`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/api/v2/account/balances", r.URL.String())
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assert.Equal(t, "XMLHttpRequest", r.Header.Get(HeaderXRequestedWith))

		assert.Equal(t, pseudoAPIKey.String(), r.Header.Get(HeaderAuthorization))
		assert.Equal(t, "f6f8463e8c41574ccb2bfc46938868b78b0b6a1a0cb75cf58a92eba3d6535b9f", r.Header.Get(HeaderSignature))

		expectedReqBody := `{
			"timestamp" : 1530682938651,
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
	request := &AccountBalancesRequest{
		Timestamp:  1530682938651,
		RecvWindow: 5000,
	}
	resp, err := client.AccountBalances(request)

	assert.NotNil(t, resp, fmt.Sprintf("error: %v", err))

	respBytes, _ := json.Marshal(resp)
	equal, _ := isEqualJSON(body, string(respBytes))
	assert.True(t, equal, fmt.Sprintf("%s is not equal to %s", body, string(respBytes)))
}
