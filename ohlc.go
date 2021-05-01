package kraken

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

type OHLC struct {
	Timestamp  int64
	Open       float64
	High       float64
	Low        float64
	Close      float64
	VWAP       float64
	Volume     float64
	TradeCount int64
}

type OHLCResponse struct {
	Current OHLC
	History []OHLC
}

// OHLC returns OHLC entries based on market pair, interval and start time
func (c Client) OHLC(pair string, interval int64, since time.Time) (OHLCResponse, error) {
	resp, err := c.getOHLC(pair, interval, since)
	if err != nil {
		return OHLCResponse{}, Wrap(err)
	}

	if len(resp.items) == 0 {
		return OHLCResponse{}, Wrap(fmt.Errorf("empty OHLC response"))
	}

	var o OHLCResponse
	for _, item := range resp.items {
		fmt.Println(item.Timestamp, resp.last, item.Timestamp <= resp.last)
		if item.Timestamp <= resp.last {
			o.History = append(o.History, item)
		} else {
			o.Current = item
		}
	}

	return o, nil
}

type ohlcKrakenResponse struct {
	Result map[string]interface{}
	Error  KrakenErrors
}

type ohlcResponse struct {
	items []OHLC
	last  int64
}

const ohlcPath = "/0/public/OHLC"

func (c Client) getOHLC(pair string, interval int64, since time.Time) (ohlcResponse, error) {
	url := fmt.Sprintf("%s/%s/public/OHLC?pair=%s&interval=%v&since=%v",
		baseURL, apiVersion, pair, interval, since.Unix())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ohlcResponse{}, Wrap(err)
	}

	// sendPublic request
	var resp ohlcKrakenResponse
	if err := c.sendPublic(req, &resp); err != nil {
		return ohlcResponse{}, Wrap(err)
	}

	// check for response errors
	if err := resp.Error.Errors(); err != nil {
		return ohlcResponse{}, Wrap(err)
	}

	// parse the response
	var o ohlcResponse
	for k, v := range resp.Result {
		if k == "last" {
			f, err := strconv.ParseFloat(fmt.Sprintf("%v", v), 10)
			if err != nil {
				return ohlcResponse{}, Wrap(err)
			}

			o.last = int64(f)
			continue
		} else {
			s := fmt.Sprintf("%v", v)
			o.items, err = parseOHLCResponse(s)
			if err != nil {
				return ohlcResponse{}, Wrap(err)
			}
		}
	}

	return o, nil
}

var ohlcSelect = regexp.MustCompile(`((?:[0-9\.]|e\+)+)`)

// because kraken's response is an array of arrays that cannot be converted to a Go struct, we need to parse the
// string response and assign it to fields in a struct.
// this is incredibly hacky but i dont want to spend too much time on this. still working on proof of concept and can
// come back and improve performance later.
func parseOHLCResponse(s string) ([]OHLC, error) {
	var out []OHLC
	if len(s) == 0 {
		return nil, errors.New("expected OHLC array, received empty string")
	}

	result := ohlcSelect.FindAllString(s, -1)
	var counter int // tracks which field we are dealing with
	var ohlc OHLC
	var err error
	for _, s := range result {
		switch counter {
		case 0:
			f, err := strconv.ParseFloat(s, 64)
			if err != nil {
				return nil, Wrap(err)
			}

			ohlc.Timestamp = int64(f)
			counter++
		case 1:
			ohlc.Open, err = strconv.ParseFloat(s, 64)
			if err != nil {
				return nil, Wrap(err)
			}
			counter++
		case 2:
			ohlc.High, err = strconv.ParseFloat(s, 64)
			if err != nil {
				return nil, Wrap(err)
			}
			counter++
		case 3:
			ohlc.Low, err = strconv.ParseFloat(s, 64)
			if err != nil {
				return nil, Wrap(err)
			}
			counter++
		case 4:
			ohlc.Close, err = strconv.ParseFloat(s, 64)
			if err != nil {
				return nil, Wrap(err)
			}
			counter++
		case 5:
			ohlc.VWAP, err = strconv.ParseFloat(s, 64)
			if err != nil {
				return nil, Wrap(err)
			}
			counter++
		case 6:
			ohlc.Volume, err = strconv.ParseFloat(s, 64)
			if err != nil {
				return nil, Wrap(err)
			}
			counter++
		case 7:
			f, err := strconv.ParseFloat(s, 64)
			if err != nil {
				return nil, Wrap(err)
			}

			ohlc.TradeCount = int64(f)

			out = append(out, ohlc)
			ohlc = OHLC{}
			counter = 0
		}
	}

	return out, nil
}
