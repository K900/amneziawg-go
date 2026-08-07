//go:build !linux

package device

import (
	"github.com/K900/amneziawg-go/conn"
	"github.com/K900/amneziawg-go/rwcancel"
)

func (device *Device) startRouteListener(_ conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
