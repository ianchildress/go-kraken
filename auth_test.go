package kraken

import (
	"encoding/base64"
	"net/url"
	"testing"
)

func Test_getKrakenSignature(t *testing.T) {
	want := "4/dpxb3iT4tp/ZCVEwSnEsLxx0bqyhLpdfOpc6fn7OR8+UClSV5n9E6aSS8MPtnRfp32bAb0nmbRn6H8ndwLUQ=="
	apiSecret := "kQH5HW/8p1uGOVjbgWA7FunAmGO8lsSUXNsu3eow76sz84Q18fWxnyRzBHCd3pd5nE9qa99HAZtuZuj6F1huXg=="

	payload := url.Values{}
	payload.Add("pair", "XBTUSD")
	payload.Add("type", "buy")
	payload.Add("ordertype", "limit")
	payload.Add("price", "37500")
	payload.Add("volume", "1.25")
	payload.Add("nonce", "1616492376594")

	b64DecodedSecret, _ := base64.StdEncoding.DecodeString(apiSecret)

	signature := getKrakenSignature("/0/private/AddOrder", payload, b64DecodedSecret)
	if signature != want {
		t.Errorf("wanted %v got %v", want, signature)
	}

}
