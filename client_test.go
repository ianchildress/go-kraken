package kraken

import (
	"fmt"
	"testing"
)

func TestClient_ServerTime(t *testing.T) {
	key := "foo"
	secret := "bar"
	client := NewClient(key, secret)
	resp, err := client.ServerTime()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(resp.Result)
}
