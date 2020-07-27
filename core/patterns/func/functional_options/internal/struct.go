package internal

import "time"

type Server struct {
	addr         string
	timeout      time.Duration
	whitelistIPs []string
}

type Option func(s *Server)

func Timeout(timeout time.Duration) func(*Server) {
	return func(s *Server) {
		s.timeout = timeout
	}
}

func WithWhitelistedIP(ip string) Option {
	return func(s *Server) {
		s.whitelistIPs = append(s.whitelistIPs, ip)
	}
}

func NewServer(addr string, opts ...Option) *Server {
	server := &Server{
		addr: addr,
	}

	// apply the list of options to Server
	for _, opt := range opts {
		opt(server)
	}

	return server
}
