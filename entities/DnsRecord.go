// Copyright [2022] [Optimus-Xs@GitHub]. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package entities

import (
	"time"
)

type DnsRecord struct {
	ID          int
	DeviceID    int
	Device      Device
	IpAddress   string
	Domain      string `gorm:"uniqueIndex"`
	DnsType     uint16
	StorageType string
	UpdateAt    time.Time
}
