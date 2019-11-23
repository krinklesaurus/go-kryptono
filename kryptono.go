package kryptono

import (
	"encoding/json"
	"net/http"
)

const (
	generalAPIEndpoint = "https://p.kryptono.exchange/k"
	marketAPIEndpoint  = "https://engine2.kryptono.exchange"
	accountAPIEndpoint = ""
)

const (
	OrderStatusOpen        = "OPEN"
	OrderStatusFilled      = "FILLED"
	OrderStatusPartialFill = "PARTIAL_FILL"
	OrderStatusCanceled    = "CANCELED"
	OrderStatusCanceling   = "CANCELING"

	OrderTypeLimit      = "LIMIT"
	OrderTypeMarket     = "MARKET"
	OrderTypeStopLoss   = "STOP_LOSS"
	OrderTypeTakeProfit = "TAKE_PROFIT"

	OrderSideBuy  = "BUY"
	OrderSideSell = "SELL"
)

// for testing purposes only
func newClientWithURL(url string, apiKey string, apiSecret string) (Client, error) {
	return newClientWithURLs(apiKey, apiSecret, url, url, url)
}

func newClientWithURLs(apiKey string, apiSecret string, generalAPIEndpoint string, marketAPIEndpoint string, accountAPIEndpoint string) (Client, error) {
	return &client{
		http: http.DefaultClient,
		auth: &auth{
			APIKey:    apiKey,
			APISecret: apiSecret,
		},
		generalAPIEndpoint: generalAPIEndpoint,
		marketAPIEndpoint:  marketAPIEndpoint,
		accountAPIEndpoint: accountAPIEndpoint,
	}, nil
}

// NewClient creates a new kryptono client with apiKey and apiSecret
func NewClient(apiKey string, apiSecret string) (Client, error) {
	return newClientWithURLs(apiKey, apiSecret, generalAPIEndpoint, marketAPIEndpoint, accountAPIEndpoint)
}

type Client interface {
	Ping() (*PingResp, error)
	ServerTime() (*ServerTimeResp, error)
	ExchangeInformation() (*ExchangeInformationResp, error)
	MarketPrice() (MarketPriceResp, error)
	TradeHistory(symbol string) (*TradeHistoryResp, error)
	OrderBook(symbol string) (*OrderBookResp, error)
	NewOrder(request *NewOrderRequest) (*NewOrderResp, error)
	TestNewOrder(request *NewOrderRequest) (*TestNewOrderResp, error)
}

type auth struct {
	APIKey    string
	APISecret string
}

type client struct {
	http               *http.Client
	auth               *auth
	generalAPIEndpoint string
	marketAPIEndpoint  string
	accountAPIEndpoint string
}

type Float64Pair [2]float64

func (pair *Float64Pair) UnmarshalJSON(b []byte) error {
	tmp := []json.Number{}
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}

	left, err := tmp[0].Float64()
	if err != nil {
		return err
	}
	right, err := tmp[1].Float64()
	if err != nil {
		return err
	}
	*pair = Float64Pair{left, right}

	return nil
}