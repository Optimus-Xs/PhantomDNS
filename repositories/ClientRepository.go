// Copyright [2022] [Optimus-Xs@GitHub]. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package repositories

import (
	"phantomDNS/entities"
	"time"
)

// QueryClientByID query client info by client id
func QueryClientByID(ID int) (client entities.Client) {
	db.Find(&client, ID)
	return client
}

// QueryClientByDevice query client info with bind device id
func QueryClientByDevice(devID int) (client entities.Client) {
	db.Preload("Device").Find(&client, "device_id=?", devID)
	return client
}

// QueryClientByIp query client info by client ip address
func QueryClientByIp(ip string) (client entities.Client) {
	db.Preload("Device").Find(&client, "ip_address=?", ip)
	return client
}

// CreatClient creat a new client info
func CreatClient(client entities.Client) {
	db.Create(&client)
}

// UpdateClient update a exist client info
func UpdateClient(client entities.Client) {
	dbClient := entities.Client{}
	db.Find(&dbClient, client.ID)
	dbClient.IpAddress = client.IpAddress
	dbClient.UpdateAt = time.Now()
	db.Save(&dbClient)
}

// DeleteClient delete a exist client info
func DeleteClient(client entities.Client) {
	db.Delete(&client)
}
