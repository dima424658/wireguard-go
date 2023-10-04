//go:build !linux

package device

import (
	"github.com/dima424658/wireguard-go/conn"
	"github.com/dima424658/wireguard-go/rwcancel"
)

func (device *Device) startRouteListener(bind conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
