// Copyright [2022] [Optimus-Xs@GitHub]. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package repositories

import (
	"phantomDNS/entities"
	"time"
)

func QueryDnsByID(ID int) (dns entities.DnsRecord) {
	db.Find(&dns, ID)
	return dns
}

func QueryDnsByDomain(domain string) (dns entities.DnsRecord) {
	db.Preload("Device").Find(&dns, "domain=?", domain)
	return dns
}

func UpdateDns(record entities.DnsRecord) {
	dbRecord := entities.DnsRecord{}
	db.Find(&dbRecord, record.ID)
	dbRecord.IpAddress = record.IpAddress
	dbRecord.DnsType = record.DnsType
	dbRecord.UpdateAt = time.Now()
	db.Save(&dbRecord)
}

func CreatDns(record entities.DnsRecord) {
	db.Create(&record)
}

func DeleteDnsByID(record entities.DnsRecord) {
	db.Delete(&record)
}
