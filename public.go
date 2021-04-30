package kraken

import (
	"net/http"
)

const serverTimePath = "/0/public/Time"

type ServerTimeResponse struct {
	Error  KrakenErrors
	Result ServerTimeResult
}

type ServerTimeResult struct {
	UnixTime int64
}

func (c Client) ServerTime() (ServerTimeResponse, error) {
	var resp ServerTimeResponse

	req, err := http.NewRequest("GET", path(baseURL, serverTimePath), nil)
	if err != nil {
		return resp, Wrap(err)
	}

	// sendPublic request
	if err := c.sendPublic(req, &resp); err != nil {
		return resp, Wrap(err)
	}

	// check for response errors
	if err := resp.Error.Errors(); err != nil {
		return resp, Wrap(err)
	}

	return resp, nil
}
