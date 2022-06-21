// Copyright [2022] [Optimus-Xs@GitHub]. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package entities

// Device is basic unit for phantomDNS to manage client and DNS record host, each device can be registered as a client
// also a DNS record host, phantomDNS will take different action of dns query and record update base on different client
// and DNSRecord info
type Device struct {
	ID               int    // device db id
	Name             string // device name for readable notification
	RegisterName     string `gorm:"uniqueIndex"` // device register id, also used as auth username in interface for update device related info
	RegisterPassword string // device register password, used as auth password in interface for update device related info
}
