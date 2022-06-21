// Copyright [2022] [Optimus-Xs@GitHub]. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package entities

import (
	"time"
)

// DnsRecord is DNS record object which defined DNS record info for custom domain resolve and DDNS
type DnsRecord struct {
	ID          int       // DnsRecord db id
	DeviceID    int       // device db id that bind on current DnsRecord info
	Device      Device    // device is the basic unit to phantomDNS, it can be registered as client also DNSRecord Host
	IpAddress   string    // dns record ip address, this is the phantomDNS custom domain resolve result
	Domain      string    `gorm:"uniqueIndex"` // dns record domain, the user custom domain that point to custom IpAddress
	DnsType     uint16    // define current DNS record is an ipv4 dns record or an ipv6 dns record
	StorageType string    // source of current DNS record, it only has ddns for now, it for support dns cache in later version
	UpdateAt    time.Time // the timestamp when this record last update
}
