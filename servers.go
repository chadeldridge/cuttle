package main

import "net"

type Protocol int

const (
	SSH Protocol = iota + 1
	SSHPwd
	RDP
	Telnet
)

type Server struct {
	Name     string
	Hostname string
	IP       net.IPAddr
	Protocol
}

func NewServer(name string, proto Protocol) Server {
	return Server{Name: name, Protocol: proto}
}
