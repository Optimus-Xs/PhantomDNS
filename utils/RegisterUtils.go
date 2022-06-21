// Copyright [2022] [Optimus-Xs@GitHub]. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// GetCurrentAuthDev provide current authed device info allow register service access to bind register info
func GetCurrentAuthDev(context *gin.Context) string {
	user, exist := context.Get("user")
	registerName := ""
	if exist {
		registerName = user.(string)
	}
	return registerName
}

// ReadConfig set all config info source from app.yml
func ReadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}
