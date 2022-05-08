package client

import (
	"net"
	"net/http"
	"runtime"
	"time"
)

// New ...
func New() *http.Client {
	var t = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		MaxIdleConnsPerHost:   runtime.GOMAXPROCS(0) + 1,
	}

	c := &http.Client{
		Timeout:   time.Second * time.Duration(30),
		Transport: t,
	}

	return c
}
