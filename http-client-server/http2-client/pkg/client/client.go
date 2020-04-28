package client

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"golang.org/x/net/http2"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

type H2client struct {
	*http.Client
}

func New() *H2client {
	hc := http.Client{}
	hc.Transport = &http2.Transport{
		AllowHTTP: true,
		DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
			return net.Dial(network, addr)
		},
	}
	return &H2client{&hc}
}

func (c *H2client) LogGet(url string) {
	if rs, err := c.Get(url); err != nil {
		log.Println(err)
	} else {
		s, _ := asString(*rs)
		log.Println(s)
	}
}

func (c *H2client) LogPostJsonRs(url string, json string) {
	data := []byte(json)
	r := bytes.NewReader(data)
	if rs, err := c.Post(url, "application/json", r); err != nil {
		log.Println(err)
	} else {
		s, _ := asString(*rs)
		log.Println(s)
	}
}

func asString(resp http.Response) (string, error) {
	if resp.StatusCode == http.StatusOK {
		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		return fmt.Sprintf("Rsponce Body:\n%s", string(bodyBytes)), nil
	}
	return "error", errors.New("not 200 status")
}
