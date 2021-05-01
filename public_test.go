package kraken

import (
	"fmt"
	"testing"
	"time"

	"github.com/hokaccha/go-prettyjson"
)

func TestClient_OHLC(t *testing.T) {
	key := "foo"
	secret := "bar"
	client := NewClient(key, secret)
	resp, err := client.OHLC("BTCUSD", 60, time.Now().Add(-time.Hour*2))
	if err != nil {
		t.Error(err)
	}

	j, err := prettyjson.Marshal(resp)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(j))
}

func TestClient_ServerTime(t *testing.T) {
	key := "foo"
	secret := "bar"
	client := NewClient(key, secret)
	resp, err := client.ServerTime()
	if err != nil {
		t.Error(err)
	}

	if resp.Result.UnixTime == 0 {
		t.Error("failed to get time")
	}
}

func Test_splitOHLC(t *testing.T) {
	_, err := parseOHLCResponse("")
	if err == nil {
		t.Errorf("excpted error")
	}

	_, err = parseOHLCResponse("[]")
	if err != nil {
		t.Error(err)
	}
	resp, err := parseOHLCResponse("[[1.6196436e+09 54450.0 54939.9 54115.0 54881.7 54537.1 179.88599954 2503] [1.6196472e+09 54901.0 54901.0 54578.6 54706.0 54761.0 30.33280582 644]]")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp)
}
