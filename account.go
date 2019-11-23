package kryptono

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type NewOrderRequest struct {
	OrderSymbol string  `json:"order_symbol"`
	OrderSide   string  `json:"order_side"`
	OrderPrice  float64 `json:"order_price,string"`
	OrderSize   float64 `json:"order_size,string"`
	Type        string  `json:"type"`
	Timestamp   int     `json:"timestamp"`
	RecvWindow  int     `json:"recvWindow"`
}

type NewOrderResp struct {
	OrderID     string `json:"order_id"`
	AccountID   string `json:"account_id"`
	OrderSymbol string `json:"order_symbol"`
	OrderSide   string `json:"order_side"`
	Status      string `json:"status"`
	CreateTime  int64  `json:"createTime"`
	Type        string `json:"type"`
	OrderPrice  string `json:"order_price"`
	OrderSize   string `json:"order_size"`
	Executed    string `json:"executed"`
	StopPrice   string `json:"stop_price"`
	Avg         string `json:"avg"`
	Total       string `json:"total"`
}

type TestNewOrderResp struct {
	Result bool `json:"result"`
}

func (c *client) NewOrder(request *NewOrderRequest) (*NewOrderResp, error) {
	url := fmt.Sprintf("%s/api/v2/order/add", c.accountAPIEndpoint)
	asJSON, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	resp, err := c.sendPost(url, nil, bytes.NewReader(asJSON))
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

	var result NewOrderResp
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *client) TestNewOrder(request *NewOrderRequest) (*TestNewOrderResp, error) {
	url := fmt.Sprintf("%s/api/v2/order/test", c.accountAPIEndpoint)
	asJSON, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	resp, err := c.sendPost(url, nil, bytes.NewReader(asJSON))
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

	var result TestNewOrderResp
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
