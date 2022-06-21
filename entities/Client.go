// Copyright [2022] [Optimus-Xs@GitHub]. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package entities

import "time"

// Client Dns server client Object which defined client info for server to do identification
type Client struct {
	ID        int       // Client database id
	DeviceID  int       `gorm:"uniqueIndex"` // device db id that bind on current Client info
	Device    Device    // device is the basic unit to phantomDNS, it can be registered as client also DNSRecord Host
	IpAddress string    `gorm:"uniqueIndex"` //	client's internet ip address, phantomDNS only answer dns query after client ip address is validated
	UpdateAt  time.Time // when did the client info update last time which often means ip address update
}
