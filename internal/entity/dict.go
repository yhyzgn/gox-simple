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
// time   : 2020-05-15 16:13
// version: 1.0.0
// desc   : 企业库数据字典

package entity

import "gox-simple/internal/entity/table"

const (
	IndustryCategory = iota + 1 // 行业类别
	CertificateType             // 证件类别
	BusinessStatus              // 经营状态
	EnterpriseLevel             // 企业等级
)

type Dict struct {
	Code uint   `gorm:"column:code;type:int(11);comment:'编码'" json:"code"`
	Name string `gorm:"column:name;type:varchar(255);comment:'名称'" json:"name"`
	Type uint8  `gorm:"column:type;type:tinyint(8);comment:'类型：1-行业类别，2-证件类别，3-经营状态'" json:"type"`
	table.Entity
}

func (d Dict) TableName() string {
	return d.NameWithPrefix("dict")
}
