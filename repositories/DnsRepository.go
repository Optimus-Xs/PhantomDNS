// Copyright [2022] [Optimus-Xs@GitHub]. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package repositories

import (
	"phantomDNS/entities"
	"time"
)

// QueryDnsByID query DnsRecord info by dns id
func QueryDnsByID(ID int) (dns entities.DnsRecord) {
	db.Find(&dns, ID)
	return dns
}

// QueryDnsByDomain query DnsRecord info by dns record domain
func QueryDnsByDomain(domain string) (dns entities.DnsRecord) {
	db.Preload("Device").Find(&dns, "domain=?", domain)
	return dns
}

// UpdateDns update a exist DnsRecord info
func UpdateDns(record entities.DnsRecord) {
	dbRecord := entities.DnsRecord{}
	db.Find(&dbRecord, record.ID)
	dbRecord.IpAddress = record.IpAddress
	dbRecord.DnsType = record.DnsType
	dbRecord.UpdateAt = time.Now()
	db.Save(&dbRecord)
}

// CreatDns creat a new DnsRecord info
func CreatDns(record entities.DnsRecord) {
	db.Create(&record)
}

// DeleteDnsByID delete a exist DnsRecord info
func DeleteDnsByID(record entities.DnsRecord) {
	db.Delete(&record)
}
