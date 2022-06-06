// Copyright [2022] [Optimus-Xs@GitHub]. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package entities

type Device struct {
	ID               int
	Name             string
	RegisterName     string `gorm:"uniqueIndex"`
	RegisterPassword string
}
