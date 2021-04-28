package kraken

import (
	"time"
)

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
