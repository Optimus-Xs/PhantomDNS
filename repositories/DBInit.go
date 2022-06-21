// Copyright [2022] [Optimus-Xs@GitHub]. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package repositories

import (
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"phantomDNS/entities"
	"phantomDNS/utils"
)

var (
	db  *gorm.DB
	err error
)

// init db connection info, include read db Storage path from app.yml, set AutoMigrate
func init() {
	utils.ReadConfig()
	db, err = gorm.Open(sqlite.Open(viper.GetString("SqliteStorage")), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&entities.Device{})
	db.AutoMigrate(&entities.DnsRecord{})
	db.AutoMigrate(&entities.Client{})
}
