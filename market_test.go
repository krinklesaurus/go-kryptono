package kryptono

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestTradeHistory(t *testing.T) {
	pseudoAPIKey := uuid.NewV4()
	pseudoAPISecret := "4a894c5c-8a7e-4337-bb6b-9fde16e3dddd"
	body := `{
		"symbol":"KNOW_BTC",
		"limit":100,
		"history":[
		  {
			"id":139638,
			"price":"0.00001723",
			"qty":"81.00000000",
			"isBuyerMaker":false,
			"time":1529262196270
		  }
		],
		"time":1529298130192
	  }`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/api/v1/ht?symbol=KNOW_BTC", r.URL.String())
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assert.Equal(t, "XMLHttpRequest", r.Header.Get(HeaderXRequestedWith))

		assert.Equal(t, pseudoAPIKey.String(), r.Header.Get(HeaderAuthorization))
		assert.Empty(t, r.Header.Get(HeaderSignature))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(body))
	}))
	defer ts.Close()

	client, err := newClientWithURL(ts.URL, pseudoAPIKey.String(), pseudoAPISecret)
	if err != nil {
		t.Error(err.Error())
	}

	resp, err := client.TradeHistory("KNOW_BTC")
	if err != nil {
		t.Error(err.Error())
	}

	assert.NotNil(t, resp, fmt.Sprintf("error: %v", err))
	assert.Equal(t, "KNOW_BTC", resp.Symbol)
	assert.Equal(t, 100, resp.Limit)
	assert.Equal(t, 1529298130192, resp.Time)

	assert.Equal(t, 1, len(resp.History))
	assert.Equal(t, 139638, resp.History[0].ID)
	assert.Equal(t, 0.00001723, resp.History[0].Price)
	assert.Equal(t, 81.00000000, resp.History[0].Qty)
	assert.Equal(t, false, resp.History[0].IsBuyerMaker)
	assert.Equal(t, 1529262196270, resp.History[0].Time)
}

func TestOrderBook(t *testing.T) {
	pseudoAPIKey := uuid.NewV4()
	pseudoAPISecret := "4a894c5c-8a7e-4337-bb6b-9fde16e3dddd"
	body := `{
		"symbol": "KNOW_BTC",
		"limit": 100,
		"asks": [
			[
				"0.00000035",
				"17790.00000000"
			]
		],
		"bids": [
			[
				"0.00000019",
				"21052.00000000"
			]
		],
		"time": 1574517091326
	}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/api/v1/dp?symbol=KNOW_BTC", r.URL.String())
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assert.Equal(t, "XMLHttpRequest", r.Header.Get(HeaderXRequestedWith))

		assert.Equal(t, pseudoAPIKey.String(), r.Header.Get(HeaderAuthorization))
		assert.Empty(t, r.Header.Get(HeaderSignature))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(body))
	}))
	defer ts.Close()

	client, err := newClientWithURL(ts.URL, pseudoAPIKey.String(), pseudoAPISecret)
	if err != nil {
		t.Error(err.Error())
	}

	resp, err := client.OrderBook("KNOW_BTC")
	if err != nil {
		t.Error(err.Error())
	}

	assert.NotNil(t, resp, fmt.Sprintf("error: %v", err))
	assert.Equal(t, "KNOW_BTC", resp.Symbol)
	assert.Equal(t, 100, resp.Limit)
	assert.Equal(t, 1574517091326, resp.Time)

	assert.Equal(t, 1, len(resp.Asks))
	assert.Equal(t, 0.00000035, resp.Asks[0][0])
	assert.Equal(t, 17790.00000000, resp.Asks[0][1])

	assert.Equal(t, 1, len(resp.Bids))
	assert.Equal(t, 0.00000019, resp.Bids[0][0])
	assert.Equal(t, 21052.00000000, resp.Bids[0][1])
}

func TestMarketSummaries(t *testing.T) {
	pseudoAPIKey := uuid.NewV4()
	pseudoAPISecret := "4a894c5c-8a7e-4337-bb6b-9fde16e3dddd"
	body := `{
		"success": true,
		"message": "",
		"result": [
			{
				"MarketName": "EOS-BTC",
				"High": 0.0003664,
				"Low": 0.0003539,
				"BaseVolume": 187.39362915,
				"Last": 0.0003599,
				"TimeStamp": "2019-11-24T11:06:16.081Z",
				"Volume": 520682.49277577107,
				"Ask": "0.00036880",
				"Bid": "0.00035120",
				"PrevDay": 0.0003568
			},
			{
				"MarketName": "LYL-BTC",
				"High": 4e-8,
				"Low": 4e-8,
				"BaseVolume": 0,
				"Last": 4e-8,
				"TimeStamp": "2019-11-24T11:06:16.081Z",
				"Volume": 0,
				"Ask": 4e-8,
				"Bid": 4e-8,
				"PrevDay": 4e-8
			}
		],
		"volumes": [
			{
				"CoinName": "BTC",
				"Volume": 572.2267145000001
			}
		],
		"t": 1574593579082
	}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/v1/getmarketsummaries", r.URL.String())
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assert.Equal(t, "XMLHttpRequest", r.Header.Get(HeaderXRequestedWith))

		assert.Equal(t, pseudoAPIKey.String(), r.Header.Get(HeaderAuthorization))
		assert.Empty(t, r.Header.Get(HeaderSignature))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(body))
	}))
	defer ts.Close()

	client, err := newClientWithURL(ts.URL, pseudoAPIKey.String(), pseudoAPISecret)
	if err != nil {
		t.Error(err.Error())
	}

	resp, err := client.MarketSummaries()
	if err != nil {
		t.Error(err.Error())
	}

	assert.NotNil(t, resp, fmt.Sprintf("error: %v", err))
	assert.Equal(t, true, resp.Success)
	assert.Equal(t, "", resp.Message)
}
