// +build cgo

package motorrent

import (
	"github.com/anacrolix/go-libutp"
)

func NewUtpSocket(network, addr string) (utpSocket, error) {
	return utp.NewSocket(network, addr)
}
