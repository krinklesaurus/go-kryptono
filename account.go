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
	StopPrice   string  `json:"stop_price,omitempty"`
	Type        string  `json:"type"`
	Timestamp   int     `json:"timestamp"`
	RecvWindow  int     `json:"recvWindow,omitempty"`
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

type OrderDetailRequest struct {
	OrderID    string `json:"order_id"`
	Timestamp  int64  `json:"timestamp"`
	RecvWindow int64  `json:"recvWindow,omitempty"`
}

type OrderDetailResp struct {
	OrderID     string  `json:"order_id"`
	AccountID   string  `json:"account_id"`
	OrderSymbol string  `json:"order_symbol"`
	OrderSide   string  `json:"order_side"`
	Status      string  `json:"status"`
	CreateTime  int     `json:"createTime"`
	Type        string  `json:"type"`
	OrderPrice  float64 `json:"order_price,string"`
	OrderSize   float64 `json:"order_size,string"`
	Executed    float64 `json:"executed,string"`
	StopPrice   float64 `json:"stop_price,string"`
	Avg         float64 `json:"avg,string"`
	Total       string  `json:"total"`
}

type CancelOrderRequest struct {
	OrderID     string `json:"order_id"`
	OrderSymbol string `json:"order_symbol"`
	Timestamp   int    `json:"timestamp"`
	RecvWindow  int    `json:"recvWindow,omitempty"`
}

type CancelOrderResp struct {
	OrderID     string `json:"order_id"`
	OrderSymbol string `json:"order_symbol"`
}

type TradeDetailsRequest struct {
	OrderID    string `json:"order_id"`
	Timestamp  int64  `json:"timestamp"`
	RecvWindow int64  `json:"recvWindow,omitempty"`
}

type TradeDetailsResp []TradeDetailsRespElement

type TradeDetailsRespElement struct {
	HexID     string  `json:"hex_id"`
	Symbol    string  `json:"symbol"`
	OrderID   string  `json:"order_id"`
	OrderSide string  `json:"order_side"`
	Price     float64 `json:"price,string"`
	Quantity  float64 `json:"quantity,string"`
	Fee       string  `json:"fee"`
	Total     string  `json:"total"`
	Timestamp int     `json:"timestamp"`
}

type OpenOrdersRequest struct {
	Symbol     string `json:"symbol"`
	FromID     string `json:"from_id,omitempty"`
	Limit      int    `json:"limit,omitempty"`
	Page       int    `json:"page"`
	Timestamp  int    `json:"timestamp"`
	RecvWindow int    `json:"recvWindow,omitempty"`
}

type OpenOrdersResp struct {
	Total int                     `json:"total"`
	List  []OpenOrdersRespElement `json:"list"`
}

type OpenOrdersRespElement struct {
	OrderID     string  `json:"order_id"`
	AccountID   string  `json:"account_id"`
	OrderSymbol string  `json:"order_symbol"`
	OrderSide   string  `json:"order_side"`
	Status      string  `json:"status"`
	CreateTime  int     `json:"createTime"`
	Type        string  `json:"type"`
	OrderPrice  float64 `json:"order_price,string"`
	OrderSize   float64 `json:"order_size,string"`
	Executed    float64 `json:"executed,string"`
	StopPrice   float64 `json:"stop_price,string"`
	Avg         float64 `json:"avg,string"`
	Total       string  `json:"total"`
}

type CompletedOrdersRequest struct {
	Symbol     string `json:"symbol"`
	FromID     string `json:"from_id,omitempty"`
	Limit      int    `json:"limit,omitempty"`
	Page       int    `json:"page"`
	Timestamp  int    `json:"timestamp"`
	RecvWindow int    `json:"recvWindow"`
}

type CompletedOrdersResp struct {
	Total int                          `json:"total"`
	List  []CompletedOrdersRespElement `json:"list"`
}

type CompletedOrdersRespElement struct {
	OrderID     string  `json:"order_id"`
	AccountID   string  `json:"account_id"`
	OrderSymbol string  `json:"order_symbol"`
	OrderSide   string  `json:"order_side"`
	Status      string  `json:"status"`
	CreateTime  int     `json:"createTime"`
	Type        string  `json:"type"`
	OrderPrice  float64 `json:"order_price,string"`
	OrderSize   float64 `json:"order_size,string"`
	Executed    float64 `json:"executed,string"`
	StopPrice   float64 `json:"stop_price,string"`
	Avg         float64 `json:"avg,string"`
	Total       string  `json:"total"`
}

type AllOrdersRequest struct {
	Symbol     string `json:"symbol"`
	FromID     string `json:"from_id,omitempty"`
	Limit      int64  `json:"limit,omitempty"`
	Timestamp  int64  `json:"timestamp"`
	RecvWindow int64  `json:"recvWindow,omitempty"`
}

type AllOrdersResp []AllOrdersRespElement

type AllOrdersRespElement struct {
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

type TradeListRequest struct {
	Symbol     string `json:"symbol"`
	FromID     string `json:"from_id,omitempty"`
	Limit      int    `json:"limit,omitempty"`
	Timestamp  int    `json:"timestamp"`
	RecvWindow int    `json:"recvWindow,omitempty"`
}

type TradeListResp []TradeListRespElement

type TradeListRespElement struct {
	HexID     string `json:"hex_id"`
	Symbol    string `json:"symbol"`
	OrderID   string `json:"order_id"`
	OrderSide string `json:"order_side"`
	Price     string `json:"price"`
	Quantity  string `json:"quantity"`
	Fee       string `json:"fee"`
	Total     string `json:"total"`
	Timestamp int64  `json:"timestamp"`
}

type AccountInformationRequest struct {
	Timestamp  int `json:"timestamp"`
	RecvWindow int `json:"recvWindow,omitempty"`
}

type AccountInformationResp struct {
	AccountID        string           `json:"account_id"`
	Email            string           `json:"email"`
	Phone            interface{}      `json:"phone"`
	EnableGoogle2Fa  bool             `json:"enable_google_2fa"`
	Status           string           `json:"status"`
	CreateAt         int              `json:"create_at"`
	NickName         string           `json:"nick_name"`
	ChatID           string           `json:"chat_id"`
	ChatPassword     string           `json:"chat_password"`
	Banks            []interface{}    `json:"banks"`
	Country          string           `json:"country"`
	Language         string           `json:"language"`
	KycStatus        interface{}      `json:"kyc_status"`
	KycLevel         string           `json:"kyc_level"`
	LastLoginHistory LastLoginHistory `json:"last_login_history"`
	CommissionStatus bool             `json:"commission_status"`
	AccountKyc       interface{}      `json:"account_kyc"`
	KycRejectInfos   []interface{}    `json:"kyc_reject_infos"`
	AllowOrder       int              `json:"allow_order"`
	DisableWithdraw  int              `json:"disable_withdraw"`
	ReferralID       string           `json:"referral_id"`
	FavoritePairs    []string         `json:"favorite_pairs"`
	ChatServer       string           `json:"chat_server"`
	ExchangeFee      ExchangeFee      `json:"exchange_fee"`
}

type ExchangeFee struct {
	StandardFee float64 `json:"standard_fee,string"`
	KnowFee     float64 `json:"know_fee,string"`
}

type LastLoginHistory struct {
	ID          ID     `json:"id"`
	AccountID   string `json:"account_id"`
	NickName    string `json:"nick_name"`
	Email       string `json:"email"`
	IPAddress   string `json:"ip_address"`
	LoginAt     int    `json:"login_at"`
	OSName      string `json:"os_name"`
	BrowserName string `json:"browser_name"`
	Country     string `json:"country"`
	City        string `json:"city"`
	SentEmail   bool   `json:"sentEmail"`
}

type ID struct {
	Timestamp         int `json:"timestamp"`
	MachineIdentifier int `json:"machineIdentifier"`
	ProcessIdentifier int `json:"processIdentifier"`
	Counter           int `json:"counter"`
	Time              int `json:"time"`
	Date              int `json:"date"`
	TimeSecond        int `json:"timeSecond"`
}

type AccountBalancesRequest struct {
	Timestamp  int `json:"timestamp"`
	RecvWindow int `json:"recvWindow,omitempty"`
}

type AccountBalancesResp []AccountBalancesRespElement

type AccountBalancesRespElement struct {
	CurrencyCode string  `json:"currency_code"`
	Address      string  `json:"address"`
	Total        float64 `json:"total,string"`
	Available    float64 `json:"available,string"`
	InOrder      float64 `json:"in_order,string"`
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

func (c *client) OrderDetail(request *OrderDetailRequest) (*OrderDetailResp, error) {
	url := fmt.Sprintf("%s/api/v2/order/details", c.accountAPIEndpoint)
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

	var result OrderDetailResp
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *client) CancelOrder(request *CancelOrderRequest) (*CancelOrderResp, error) {
	url := fmt.Sprintf("%s/api/v2/order/cancel", c.accountAPIEndpoint)
	asJSON, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	resp, err := c.sendDelete(url, nil, bytes.NewReader(asJSON))
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

	var result CancelOrderResp
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *client) TradeDetails(request *TradeDetailsRequest) (*TradeDetailsResp, error) {
	url := fmt.Sprintf("%s/api/v2/order/trade-detail", c.accountAPIEndpoint)
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

	var result TradeDetailsResp
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *client) OpenOrders(request *OpenOrdersRequest) (*OpenOrdersResp, error) {
	url := fmt.Sprintf("%s/api/v2/order/list/open", c.accountAPIEndpoint)
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

	var result OpenOrdersResp
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *client) CompletedOrders(request *CompletedOrdersRequest) (*CompletedOrdersResp, error) {
	url := fmt.Sprintf("%s/api/v2/order/list/completed", c.accountAPIEndpoint)
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

	var result CompletedOrdersResp
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *client) AllOrders(request *AllOrdersRequest) (*AllOrdersResp, error) {
	url := fmt.Sprintf("%s/api/v2/order/list/all", c.accountAPIEndpoint)
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

	var result AllOrdersResp
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *client) TradeList(request *TradeListRequest) (*TradeListResp, error) {
	url := fmt.Sprintf("%s/api/v2/order/list/trades", c.accountAPIEndpoint)
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

	var result TradeListResp
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *client) AccountInformation(request *AccountInformationRequest) (*AccountInformationResp, error) {
	url := fmt.Sprintf("%s/api/v2/account/details", c.accountAPIEndpoint)
	asJSON, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	resp, err := c.sendGet(url, nil, bytes.NewReader(asJSON))
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

	var result AccountInformationResp
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *client) AccountBalances(request *AccountBalancesRequest) (*AccountBalancesResp, error) {
	url := fmt.Sprintf("%s/api/v2/account/balances", c.accountAPIEndpoint)
	asJSON, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	resp, err := c.sendGet(url, nil, bytes.NewReader(asJSON))
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

	var result AccountBalancesResp
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
