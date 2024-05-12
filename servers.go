package main

import "net"

type Server struct {
	Name     string
	Hostname string
	IP       net.IPAddr
}
