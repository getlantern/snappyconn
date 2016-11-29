package snappyconn

import (
	"github.com/golang/snappy"
	"net"
)

// Wrap wraps a connection and adds snappy compression on reading and writing.
func Wrap(wrapped net.Conn) net.Conn {
	return &snappyconn{wrapped, snappy.NewReader(wrapped), snappy.NewWriter(wrapped)}
}

type snappyconn struct {
	net.Conn
	r *snappy.Reader
	w *snappy.Writer
}

func (c *snappyconn) Read(b []byte) (n int, err error) {
	return c.r.Read(b)
}

func (c *snappyconn) Write(b []byte) (n int, err error) {
	return c.w.Write(b)
}
