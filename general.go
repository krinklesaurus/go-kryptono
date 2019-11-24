package kryptono

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type PingResp struct {
	Result bool `json:"result"`
}

type ServerTimeResp struct {
	ServerTime int `json:"server_time"`
}

type ExchangeInformationResp struct {
	Timezone       string         `json:"timezone"`
	ServerTime     int            `json:"server_time"`
	RateLimits     []RateLimit    `json:"rate_limits"`
	BaseCurrencies []BaseCurrency `json:"base_currencies"`
	Coins          []Coin         `json:"coins"`
	Symbols        []Symbol       `json:"symbols"`
}

type BaseCurrency struct {
	CurrencyCode      string  `json:"currency_code"`
	MinimumTotalOrder float64 `json:"minimum_total_order,string"`
}

type Coin struct {
	CurrencyCode       string  `json:"currency_code"`
	Name               string  `json:"name"`
	MinimumOrderAmount float64 `json:"minimum_order_amount,string"`
}

type RateLimit struct {
	Type     string `json:"type"`
	Interval string `json:"interval"`
	Limit    int    `json:"limit"`
}

type Symbol struct {
	Symbol             string  `json:"symbol"`
	AmountLimitDecimal float64 `json:"amount_limit_decimal"`
	PriceLimitDecimal  float64 `json:"price_limit_decimal"`
	AllowTrading       bool    `json:"allow_trading"`
}

type MarketPriceResp []MarketPriceRespElement

type MarketPriceRespElement struct {
	Symbol      string  `json:"symbol"`
	Price       float64 `json:"price,string"`
	UpdatedTime int     `json:"updated_time"`
}

func (c *client) Ping() (*PingResp, error) {
	url := fmt.Sprintf("%s/api/v2/ping", c.generalAPIEndpoint)
	resp, err := c.sendGet(url, nil, nil)
	if err != nil {
		return nil, err
	}
	err = checkHTTPStatus(*resp, http.StatusOK)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result PingResp
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *client) ServerTime() (*ServerTimeResp, error) {
	url := fmt.Sprintf("%s/api/v2/time", c.generalAPIEndpoint)
	resp, err := c.sendGet(url, nil, nil)
	if err != nil {
		return nil, err
	}
	err = checkHTTPStatus(*resp, http.StatusOK)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result ServerTimeResp
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *client) ExchangeInformation() (*ExchangeInformationResp, error) {
	url := fmt.Sprintf("%s/api/v2/exchange-info", c.generalAPIEndpoint)
	resp, err := c.sendGet(url, nil, nil)
	if err != nil {
		return nil, err
	}
	err = checkHTTPStatus(*resp, http.StatusOK)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result ExchangeInformationResp
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *client) MarketPrice(symbol string) (MarketPriceResp, error) {
	url := fmt.Sprintf("%s/api/v2/market-price", c.generalAPIEndpoint)
	if symbol != "" {
		url = fmt.Sprintf("%s?symbol=%s", url, symbol)
	}
	resp, err := c.sendGet(url, nil, nil)
	if err != nil {
		return nil, err
	}
	err = checkHTTPStatus(*resp, http.StatusOK)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result MarketPriceResp
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
