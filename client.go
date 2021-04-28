package kraken

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const baseURL = "https://api.kraken.com"

type Client struct {
	key    string
	secret string
}

func NewClient(key, secret string) Client {
	return Client{key: key, secret: secret}
}

func nonce() int64 {
	return time.Now().Unix()
}

const serverTimePath = "/0/public/Time"

type ServerTimeResponse struct {
	Error  []error
	Result ServerTimeResult
}

type ServerTimeResult struct {
	UnixTime int64
	RFC1123  time.Time
}

func (c Client) ServerTime() (ServerTimeResponse, error) {
	var out ServerTimeResponse

	req, err := http.NewRequest("GET", path(baseURL, serverTimePath), nil)
	if err != nil {
		return out, err
	}

	req.Header.Set("Cache-Control", "no-cache")

	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		return out, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return out, err
	}

	if err := json.Unmarshal(body, &out); err != nil {
		return out, err
	}

	return out, nil
}

func path(base, suffix string) string {
	return fmt.Sprintf("%s%s", base, suffix)
}
