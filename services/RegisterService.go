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

// RegisterInterface define device, client and DNSRecord register related methods
type RegisterInterface interface {
	DoHostRegister(ip string, domain string, regName string) // creat or update DNSRecord info
	DoClientRegister(ip string, regName string)              // creat or update Client info
	GetDevices() (deviceAccounts gin.Accounts)               // get all Device and register into gin.Accounts to implant basic auth
}

// RegisterServices is implantation of RegisterInterface
type RegisterServices struct {
}

// DoHostRegister is implantation of RegisterInterface.DoHostRegister it creat a new DNSRecord when domain is not exist
// in db, and update it when domain is already there
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

// DoClientRegister is implantation of RegisterInterface.DoClientRegister it creat a new client info when client id is
// not exist in db, and update it when client is already there
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

// GetDevices is implantation of RegisterInterface.GetDevices it read all device RegisterName and RegisterPassword for
// use gin basic auth service
func (s RegisterServices) GetDevices() (deviceAccounts gin.Accounts) {
	devices := repositories.GetAllDevices()
	deviceAccounts = gin.Accounts{}
	for _, device := range devices {
		deviceAccounts[device.RegisterName] = device.RegisterPassword
	}
	return deviceAccounts
}
