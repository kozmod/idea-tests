package test

import (
	"fmt"
	"net"
	"net/url"
	"testing"
)

func TestUrlParse(t *testing.T) {
	s := "http://user:pass@host.com:5432/path?k=v#f"
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme)

	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}

func TestUrlPars2(t *testing.T) {
	s := "http://user:pass@host.com:5432/one/two"
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme)
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)
}
