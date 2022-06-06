// Copyright [2022] [Optimus-Xs@GitHub]. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package services

import (
	"github.com/gin-gonic/gin"
	"phantomDNS/entities"
	"phantomDNS/repositories"
	"phantomDNS/utils"
	"time"
)

type RegisterInterface interface {
	DoHostRegister(ip string, domain string, regName string)
	DoClientRegister(ip string, regName string)
	GetDevices() (deviceAccounts gin.Accounts)
}

type RegisterServices struct {
}

func (s RegisterServices) DoHostRegister(ip string, domain string, regName string) {
	dnsRecord := repositories.QueryDnsByDomain(domain)
	if dnsRecord.ID > 0 {
		dnsRecord.IpAddress = ip
		dnsRecord.DnsType = utils.GetIpType(ip)
		repositories.UpdateDns(dnsRecord)
	} else {
		device := repositories.QueryDeviceByRegisterName(regName)
		dnsRecord = entities.DnsRecord{
			DeviceID:    device.ID,
			Device:      device,
			IpAddress:   ip,
			Domain:      domain,
			DnsType:     utils.GetIpType(ip),
			StorageType: "ddns",
			UpdateAt:    time.Now(),
		}
		repositories.CreatDns(dnsRecord)
	}
}

func (s RegisterServices) DoClientRegister(ip string, regName string) {
	device := repositories.QueryDeviceByRegisterName(regName)
	client := repositories.QueryClientByDevice(device.ID)
	if client.ID > 0 {
		client.IpAddress = ip
		repositories.UpdateClient(client)
	} else {
		client = entities.Client{
			DeviceID:  device.ID,
			Device:    device,
			IpAddress: ip,
			UpdateAt:  time.Now(),
		}
		repositories.CreatClient(client)
	}
}

func (s RegisterServices) GetDevices() (deviceAccounts gin.Accounts) {
	devices := repositories.GetAllDevices()
	deviceAccounts = gin.Accounts{}
	for _, device := range devices {
		deviceAccounts[device.RegisterName] = device.RegisterPassword
	}
	return deviceAccounts
}
