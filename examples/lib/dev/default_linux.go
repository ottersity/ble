package dev

import (
	"github.com/ottersity/ble"
	"github.com/ottersity/ble/linux"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return linux.NewDevice(opts...)
}
