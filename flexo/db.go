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

package flexo

import (
	// stdlib
	"fmt"
	"os"

	// internal
	"flexo/model"

	// external
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func dbInit(user, pass, address, dbName string) error {
	db := dbCreate(user, pass, address, dbName)
	if db.Error != nil {
		fmt.Println(db.Error)
		fmt.Println("Could not create database")
		os.Exit(3)
	}

	return dbConnect(user, pass, address, dbName).AutoMigrate(&model.Team{}, &model.Category{}, &model.Target{})
}

func dbCreate(user, pass, address, dbName string) *gorm.DB {
	db := dbConnect(user, pass, address, "")

	return db.Raw(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName))
}

func dbConnect(user, pass, address, dbName string) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, address, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
