/*
Copyright © 2021 Bren 'fraq' Briggs (code@fraq.io)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package util

import (
	"fmt"
	"net"
	"net/url"
	"os"
	"strings"

	"github.com/SECCDC/flexo/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBinit(user, pass, address, dbName, sslmode string) error {
	db := DBcreate(user, pass, address, dbName, sslmode)
	if db.Error != nil {
		fmt.Println(db.Error)
		fmt.Println("Could not create database")
		os.Exit(3)
	}

	return DBconnect(user, pass, address, dbName, sslmode).AutoMigrate(&model.Team{}, &model.Category{}, &model.Target{}, &model.Event{}, &model.EcomEvent{})
}

func DBcreate(user, pass, address, dbName, sslmode string) *gorm.DB {
	db := DBconnect(user, pass, address, "", sslmode)

	return db.Raw(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName))
}

func DBconnect(user, pass, address, dbName, sslmode string) *gorm.DB {
	host, port, err := extractHostAndPort(address)
	if err != nil {
		panic("DB address isn't of the expected form host:port")
	}

	dsn := url.URL{
		User:     url.UserPassword(user, pass),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%s", host, port),
		Path:     dbName,
		RawQuery: (&url.Values{"sslmode": []string{sslmode}}).Encode(),
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn.String(),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}

func extractHostAndPort(address string) (string, string, error) {
	// Is this a domain ?
	splitAddr := strings.Split(address, ":")
	_, e := net.LookupHost(splitAddr[0])
	if e == nil {
		return splitAddr[0], splitAddr[1], nil
	}

	// If not, is this an IP?
	addr, err := net.ResolveTCPAddr("tcp", address)
	return addr.IP.String(), fmt.Sprintf("%d", addr.Port), err
}
