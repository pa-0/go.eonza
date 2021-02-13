// Copyright 2021 Alexey Krivonogov. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// +build pro

package main

import (
	"eonza/users"
	"net/http"

	pro "github.com/gentee/eonza-pro"
	"github.com/labstack/echo/v4"
)

type ProOptions struct {
	Active   bool         `json:"active"`
	Settings pro.Settings `json:"settings"`
	Trial    Trial        `json:"trial"`
}

const (
	Pro = true
)

func SetActive(active bool) error {
	return pro.SetActive(active)
}

func GetUser(id uint32) (user users.User, ok bool) {
	return pro.GetUser(id)
}

func GetUsers() []users.User {
	return pro.GetUsers()
}

func SetUserPassword(id uint32, hash []byte) error {
	return pro.SetUserPassword(id, hash)
}

func IncPassCounter(id uint32) error {
	return pro.IncPassCounter(id)
}

func ProInit(psw []byte, counter uint32) {
	pro.LoadPro(storage.Trial.Mode > TrialOff, psw, counter, cfg.path)
}

func proSettingsHandle(c echo.Context) error {
	var response ProOptions

	response.Active = pro.Active
	response.Trial = storage.Trial
	return c.JSON(http.StatusOK, &response)
}

func ProApi(e *echo.Echo) {
	pro.ProApi(e)
}
