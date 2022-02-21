package psql

import (
	"net"
	"time"

	"database/sql/driver"
	"github.com/lib/pq"
	"golang.org/x/crypto/ssh"
)

type ViaSSHDialer struct {
	Client *ssh.Client
}

func (dialer *ViaSSHDialer) Open(s string) (_ driver.Conn, err error) {
	return pq.DialOpen(dialer, s)
}

func (dialer *ViaSSHDialer) Dial(network, address string) (net.Conn, error) {
	return dialer.Client.Dial(network, address)
}

func (dialer *ViaSSHDialer) DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	return dialer.Client.Dial(network, address)
}
