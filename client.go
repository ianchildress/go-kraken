package kraken

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const baseURL = "https://api.kraken.com"
const apiVersion = "0"

type Client struct {
	key    string
	secret string
	client *http.Client
}

func NewClient(key, secret string) Client {
	return Client{key: key, secret: secret, client: &http.Client{Timeout: time.Second * 10}}
}

// path builds and returns the full url needed to reach the intended endpoint
func path(base, suffix string) string {
	return fmt.Sprintf("%s%s", base, suffix)
}

// sendPublic will send the http request and attempt to unmarshal the json response into the given interface
func (c Client) sendPublic(req *http.Request, v interface{}) error {
	resp, err := c.client.Do(req)
	if err != nil {
		return wrap(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return wrap(err)
	}

	if err := json.Unmarshal(body, v); err != nil {
		return wrap(err)
	}

	return nil
}
