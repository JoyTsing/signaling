package xrpc

import "time"

type Client struct {
	ConnectTimeout time.Duration
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
}

func NewClient(servers []string) *Client {
	return &Client{}
}
