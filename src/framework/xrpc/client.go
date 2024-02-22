package xrpc

import (
	"bufio"
	"net"
	"time"
)

const (
	defaultConnectTimeout = 100 * time.Millisecond
	defaultReadTimeout    = 500 * time.Millisecond
	defaultWriteTimeout   = 500 * time.Millisecond
)

type Client struct {
	ConnectTimeout time.Duration
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration

	selector ServerSelector
}

func NewClient(servers []string) *Client {
	s := new(RoundRobinSelector)
	s.SetServers(servers)
	return &Client{
		selector: s,
	}
}

func (c *Client) connectTimeout() time.Duration {
	if c.ConnectTimeout == 0 {
		return defaultConnectTimeout
	}
	return c.ConnectTimeout
}

func (c *Client) readTimeout() time.Duration {
	if c.ReadTimeout == 0 {
		return defaultReadTimeout
	}
	return c.ReadTimeout
}

func (c *Client) writeTimeout() time.Duration {
	if c.WriteTimeout == 0 {
		return defaultWriteTimeout
	}
	return c.WriteTimeout
}

func (c *Client) Do(req *Request) (*Response, error) {
	//chose a server
	addr, err := c.selector.PickServer()
	if err != nil {
		return nil, err
	}
	//connect to server
	connect, err := net.DialTimeout(addr.Network(), addr.String(), c.connectTimeout())
	if err != nil {
		return nil, err
	}
	connect.SetReadDeadline(time.Now().Add(c.readTimeout()))
	connect.SetWriteDeadline(time.Now().Add(c.writeTimeout()))

	rw := bufio.NewReadWriter(bufio.NewReader(connect), bufio.NewWriter(connect))
	if _, err := req.Write(rw); err != nil {
		return nil, err
	}
	if err := rw.Flush(); err != nil {
		return nil, err
	}
	return &Response{}, nil
}
