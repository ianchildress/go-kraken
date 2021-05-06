package kraken

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Order struct {
	OrderType string
	Direction string
	Amount    float64
	Pair      string
	Price     float64
	Leverage  string
	Validate  bool
}

type krakenAddOrder struct {
	Nonce     int64  `json:"nonce"`
	OrderType string `json:"ordertype"`
	Type      string `json:"type"`
	Volume    string `json:"volume"`
	Pair      string `json:"pair"`
	Price     string `json:"price"`
	Leverage  string `json:"leverage"`
	Validate  bool   `json:"validate"`
}

type krakenAddOrderResponse struct {
	Result map[string]string
	Error  []string
}

const addOrderPath = "/0/private/AddOrder"

func (c Client) OpenOrder(o Order) (string, error) {

	nonce := time.Now().Unix()
	//ko := krakenAddOrder{
	//	Nonce:     time.Now().Unix(),
	//	OrderType: o.OrderType,
	//	Type:      o.Direction,
	//	Volume:    fmt.Sprintf("%.6f", o.Amount),
	//	Pair:      o.Pair,
	//	Price:     fmt.Sprintf("%.6f", o.Price),
	//	Leverage:  o.Leverage,
	//	Validate:  false,
	//}

	payload := url.Values{}
	payload.Add("pair", o.Pair)
	payload.Add("type", o.Direction)
	payload.Add("ordertype", o.OrderType)
	payload.Add("price", fmt.Sprintf("%.6f", o.Price))
	payload.Add("volume", fmt.Sprintf("%.6f", o.Amount))
	payload.Add("nonce", fmt.Sprintf("%v", nonce))

	b64DecodedSecret, _ := base64.StdEncoding.DecodeString(c.secret)
	sig := getKrakenSignature(addOrderPath, payload, b64DecodedSecret)

	req, err := http.NewRequest("POST", path(baseURL, addOrderPath), strings.NewReader(payload.Encode()))
	if err != nil {
		return "", Wrap(err)
	}
	req = c.setPrivateHeaders(req, sig)

	var foo interface{}
	err = c.sendPrivate(req, &foo)
	return "", nil
}
