package netdb

import (
	"io"

	"github.com/bonedaddy/go-i2p/lib/common"
)

// netdb entry
// wraps a router info and provides serialization
type Entry struct {
	ri common.RouterInfo
}

func (e *Entry) WriteTo(w io.Writer) (err error) {
	return
}

func (e *Entry) ReadFrom(r io.Reader) (err error) {
	return
}
