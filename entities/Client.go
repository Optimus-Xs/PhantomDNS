// Copyright [2022] [Optimus-Xs@GitHub]. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package entities

import "time"

type Client struct {
	ID        int
	DeviceID  int `gorm:"uniqueIndex"`
	Device    Device
	IpAddress string `gorm:"uniqueIndex"`
	UpdateAt  time.Time
}