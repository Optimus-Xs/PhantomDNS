// Copyright [2022] [Optimus-Xs@GitHub]. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package repositories

import (
	"phantomDNS/entities"
)

// QueryDeviceByID query device info by device id
func QueryDeviceByID(ID int) (dev entities.Device) {
	db.Find(&dev, ID)
	return dev
}

// QueryDeviceByRegisterName query device info by device RegisterName
func QueryDeviceByRegisterName(RegisterName string) (dev entities.Device) {
	db.Find(&dev, "register_name=?", RegisterName)
	return dev
}

// GetAllDevices query all device info
func GetAllDevices() (devices []entities.Device) {
	db.Find(&devices)
	return devices
}
