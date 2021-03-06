package main

import (
	"flag"
	"net"

	"zenhack.net/go/socks5"
)

var (
	addr = flag.String("addr", ":1080", "Network address to listen on")
)

func main() {
	flag.Parse()
	socks5.ListenAndServe(&net.Dialer{}, *addr)
}
