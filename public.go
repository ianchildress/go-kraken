package kraken

import "net/http"

const serverTimePath = "/0/public/Time"

type ServerTimeResponse struct {
	Error  []error
	Result ServerTimeResult
}

type ServerTimeResult struct {
	UnixTime int64
}

func (c Client) ServerTime() (ServerTimeResponse, error) {
	var resp ServerTimeResponse

	req, err := http.NewRequest("GET", path(baseURL, serverTimePath), nil)
	if err != nil {
		return resp, wrap(err)
	}

	// sendPublic request
	if err := c.sendPublic(req, &resp); err != nil {
		return resp, wrap(err)
	}

	// check for response errors
	if err := mergeErrors(resp.Error); err != nil {
		return resp, wrap(err)
	}

	return resp, nil
}
