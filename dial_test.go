package utp

import (
	"net"
	"testing"
)

func TestDial(t *testing.T) {
	addr, err := ResolveAddr("utp", ":0")
	if err != nil {
		t.Fatal(err)
	}

	l, err := Listen("utp", addr)
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()

	ch := make(chan int)
	go func() {
		c, err := l.Accept()
		if err != nil {
			t.Fatal(err)
		}
		defer c.Close()
		ch <- 0
	}()

	_, port, err := net.SplitHostPort(l.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	raddr, err := ResolveAddr("utp", net.JoinHostPort("::1", port))
	if err != nil {
		t.Fatal(err)
	}
	c, err := DialUTP("utp", nil, raddr)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	<-ch
}
