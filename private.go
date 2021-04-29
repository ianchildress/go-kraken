package kraken

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// sendPrivate will send the http request and attempt to unmarshal the json response into the given interface
func (c Client) sendPrivate(req *http.Request, v interface{}) error {

	resp, err := c.client.Do(req)
	if err != nil {
		return Wrap(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Wrap(err)
	}

	if err := json.Unmarshal(body, v); err != nil {
		return Wrap(err)
	}

	return nil
}

func nonce() int64 {
	return time.Now().Unix()
}

func (c Client) setPrivateHeaders(req *http.Request) (*http.Request, int64) {
	nonce := time.Now().Unix()
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	req.Header.Set("API-Key", c.key)

	return req, nonce
}

type GetOpenOrdersResponse struct {
}

//func (c Client) GetOpenOrders() (GetOpenOrdersResponse, error) {
//	var resp GetOpenOrdersResponse
//	req, err := http.NewRequest("POST", path(baseURL, serverTimePath), nil)
//	if err != nil {
//		return resp, Wrap(err)
//	}
//
//}
