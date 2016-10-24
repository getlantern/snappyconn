package snappyconn

import (
	"github.com/golang/snappy"
	"net"
	"time"
)

func Wrap(wrapped net.Conn) net.Conn {
	return &snappyconn{wrapped, snappy.NewReader(wrapped), snappy.NewBufferedWriter(wrapped)}
}

type snappyconn struct {
	wrapped net.Conn
	r       *snappy.Reader
	w       *snappy.Writer
}

func (c *snappyconn) Read(b []byte) (n int, err error) {
	return c.r.Read(b)
}

func (c *snappyconn) Write(b []byte) (n int, err error) {
	n, err = c.w.Write(b)
	if err == nil {
		err = c.w.Flush()
	}
	return n, err
}

func (c *snappyconn) Close() error {
	return c.wrapped.Close()
}

func (c *snappyconn) LocalAddr() net.Addr {
	return c.wrapped.LocalAddr()
}

func (c *snappyconn) RemoteAddr() net.Addr {
	return c.wrapped.RemoteAddr()
}

func (c *snappyconn) SetDeadline(t time.Time) error {
	return c.wrapped.SetDeadline(t)
}

func (c *snappyconn) SetReadDeadline(t time.Time) error {
	return c.wrapped.SetReadDeadline(t)
}

func (c *snappyconn) SetWriteDeadline(t time.Time) error {
	return c.wrapped.SetWriteDeadline(t)
}
