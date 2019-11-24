package kryptono

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type TradeHistoryResp struct {
	Symbol  string    `json:"symbol"`
	Limit   int       `json:"limit"`
	History []History `json:"history"`
	Time    int       `json:"time"`
}

type History struct {
	ID           int     `json:"id"`
	Price        float64 `json:"price,string"`
	Qty          float64 `json:"qty,string"`
	IsBuyerMaker bool    `json:"isBuyerMaker"`
	Time         int     `json:"time"`
}

type OrderBookResp struct {
	Symbol string        `json:"symbol"`
	Asks   []Float64Pair `json:"asks"`
	Limit  int           `json:"limit"`
	Bids   []Float64Pair `json:"bids"`
	Time   int           `json:"time"`
}

func (c *client) TradeHistory(symbol string) (*TradeHistoryResp, error) {
	url := fmt.Sprintf("%s/api/v1/ht?symbol=%s", c.marketAPIEndpoint, symbol)
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

	var result TradeHistoryResp
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *client) OrderBook(symbol string) (*OrderBookResp, error) {
	url := fmt.Sprintf("%s/api/v1/dp?symbol=%s", c.marketAPIEndpoint, symbol)
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

	var result OrderBookResp
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
