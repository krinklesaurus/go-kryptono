package kryptono

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	uuid "github.com/satori/go.uuid"

	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	pseudoAPIKey := uuid.NewV4()
	pseudoAPISecret := "4a894c5c-8a7e-4337-bb6b-9fde16e3dddd"
	body := `{
		"result": true
	}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/api/v2/ping", r.URL.String())
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

	resp, err := client.Ping()
	if err != nil {
		t.Error(err.Error())
	}

	assert.NotNil(t, resp, fmt.Sprintf("error: %v", err))
	assert.Equal(t, true, resp.Result)
}

func TestServerTime(t *testing.T) {
	pseudoAPIKey := uuid.NewV4()
	pseudoAPISecret := "4a894c5c-8a7e-4337-bb6b-9fde16e3dddd"
	body := `{
		"server_time": 1530682662257
	}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/api/v2/time", r.URL.String())
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

	resp, err := client.ServerTime()
	if err != nil {
		t.Error(err.Error())
	}

	assert.NotNil(t, resp, fmt.Sprintf("error: %v", err))
	assert.Equal(t, 1530682662257, resp.ServerTime)
}

func TestExchangeInformation(t *testing.T) {
	pseudoAPIKey := uuid.NewV4()
	pseudoAPISecret := "4a894c5c-8a7e-4337-bb6b-9fde16e3dddd"
	body := `{
		"timezone": "UTC",
		"server_time": 1530683054384,
		"rate_limits": [
		  {
			"type": "REQUESTS",
			"interval": "MINUTE",
			"limit": 1000
		  }
		],
		"base_currencies": [
		  {
			"currency_code": "KNOW",
			"minimum_total_order": "100"
		  }
		],
		"coins": [
		  {
			"currency_code": "USDT",
			"name": "Tether",
			"minimum_order_amount": "1"
		  }
		],
		"symbols": [
		  {
			"symbol": "GTO_ETH",
			"amount_limit_decimal": 0,
			"price_limit_decimal": 8,
			"allow_trading": true
		  }
		]
	  }`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/api/v2/exchange-info", r.URL.String())
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

	resp, err := client.ExchangeInformation()
	if err != nil {
		t.Error(err.Error())
	}

	assert.NotNil(t, resp, fmt.Sprintf("error: %v", err))
	assert.Equal(t, "UTC", resp.Timezone)
	assert.Equal(t, 1530683054384, resp.ServerTime)

	assert.Equal(t, 1, len(resp.RateLimits))
	assert.Equal(t, "REQUESTS", resp.RateLimits[0].Type)
	assert.Equal(t, "MINUTE", resp.RateLimits[0].Interval)
	assert.Equal(t, 1000, resp.RateLimits[0].Limit)

	assert.Equal(t, 1, len(resp.BaseCurrencies))
	assert.Equal(t, "KNOW", resp.BaseCurrencies[0].CurrencyCode)
	assert.Equal(t, 100.0, resp.BaseCurrencies[0].MinimumTotalOrder)

	assert.Equal(t, 1, len(resp.Coins))
	assert.Equal(t, "USDT", resp.Coins[0].CurrencyCode)
	assert.Equal(t, "Tether", resp.Coins[0].Name)
	assert.Equal(t, 1.0, resp.Coins[0].MinimumOrderAmount)

	assert.Equal(t, 1, len(resp.Symbols))
	assert.Equal(t, "GTO_ETH", resp.Symbols[0].Symbol)
	assert.Equal(t, 0.0, resp.Symbols[0].AmountLimitDecimal)
	assert.Equal(t, 8.0, resp.Symbols[0].PriceLimitDecimal)
	assert.Equal(t, true, resp.Symbols[0].AllowTrading)
}

func TestMarketPrice(t *testing.T) {
	pseudoAPIKey := uuid.NewV4()
	pseudoAPISecret := "4a894c5c-8a7e-4337-bb6b-9fde16e3dddd"
	body := `[
		{
			"symbol": "TRX_ETH",
			"price": "0.00009317",
			"updated_time": 1574515989114
		},
		{
			"symbol": "SPIKE_BTC",
			"price": "0.00000025",
			"updated_time": 1574515989127
		}
	]`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/api/v2/market-price", r.URL.String())
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

	resp, err := client.MarketPrice("")
	if err != nil {
		t.Error(err.Error())
	}

	assert.NotNil(t, resp, fmt.Sprintf("error: %v", err))
	assert.Equal(t, 2, len(resp))

	assert.Equal(t, "TRX_ETH", resp[0].Symbol)
	assert.Equal(t, 0.00009317, resp[0].Price)
	assert.Equal(t, 1574515989114, resp[0].UpdatedTime)

	assert.Equal(t, "SPIKE_BTC", resp[1].Symbol)
	assert.Equal(t, 0.00000025, resp[1].Price)
	assert.Equal(t, 1574515989127, resp[1].UpdatedTime)
}
