package xrpc

import (
	"errors"
	"net"
	"sync"
)

type ServerSelector interface {
	PickServer() (net.Addr, error)
}

type RoundRobinSelector struct {
	sync.RWMutex
	addrs    []net.Addr
	curIndex int
}

func (r *RoundRobinSelector) SetServers(ad []string) error {
	if len(ad) == 0 {
		return errors.New("new servers is empty")
	}

	addrs := make([]net.Addr, len(ad))
	for i, server := range ad {
		tcpAddr, err := net.ResolveTCPAddr("tcp", server)
		if err != nil {
			return err
		}
		addrs[i] = tcpAddr
	}
	r.Lock()
	defer r.Unlock()
	r.addrs = addrs
	r.curIndex = 0
	return nil
}

func (r *RoundRobinSelector) PickServer() (net.Addr, error) {
	r.Lock()
	index := r.curIndex
	r.curIndex = (r.curIndex + 1) % len(r.addrs)
	r.Unlock()
	r.RLock()
	defer r.RUnlock()
	if len(r.addrs) == 0 {
		return nil, errors.New("no server available")
	}
	return r.addrs[index], nil
}
