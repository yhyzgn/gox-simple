// Copyright 2019 yhyzgn enp-simple
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2020-05-14 9:40
// version: 1.0.0
// desc   : 数据库配置类

package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gox-simple/internal/entity"
	"gox-simple/internal/entity/table"
	"gox-simple/internal/logger"
	"time"
)

type Config struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	Database    string `yaml:"database"`
	Params      string `yaml:"params"`
	MaxLifetime int    `yaml:"max-lifetime"`
	MaxOpenConn int    `yaml:"max-open-conn"`
	MaxIdleConn int    `yaml:"max-idle-conn"`
	TablePrefix string `yaml:"table-prefix"`
}

var (
	DB *gorm.DB
)

func Init(c Config) {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", c.Username, c.Password, c.Host, c.Port, c.Database, c.Params))
	if err != nil {
		logger.Fatal(err)
	}

	// 表名前缀
	table.Prefix(c.TablePrefix)

	// 取消默认的表名复数形式
	db.SingularTable(true)
	db.LogMode(true)
	db.SetLogger(logger.NewGorm())
	db.DB().SetConnMaxLifetime(time.Duration(c.MaxLifetime) * time.Second)
	db.DB().SetMaxIdleConns(c.MaxIdleConn)
	db.DB().SetMaxOpenConns(c.MaxOpenConn)

	DB = db
}

func AutoMigration() {
	log(DB.AutoMigrate(new(entity.Dict)).Error)
	log(DB.AutoMigrate(new(entity.Staff)).Error)
	log(DB.AutoMigrate(new(entity.EnterpriseSynced)).Error)
	log(DB.AutoMigrate(new(entity.Enterprise)).Error)
	log(DB.AutoMigrate(new(entity.Report)).Error)
	log(DB.AutoMigrate(new(entity.ReportLog)).Error)
}

func log(err error) {
	if err != nil {
		logger.Error(err)
	}
}
