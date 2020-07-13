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
// time   : 2020-05-14 10:26
// version: 1.0.0
// desc   : 工作人员信息

package entity

import "gox-simple/internal/entity/table"

type Staff struct {
	ID     uint64 `gorm:"column:id;type:bigint(20);primary_key;auto_increment;comment:'主键';" json:"id"`
	Code   string `gorm:"column:code;type:varchar(255);comment:'编号';" json:"code"`
	Name   string `gorm:"column:name;type:varchar(255);comment:'名字';" json:"name"`
	Status uint8  `gorm:"column:status;type:tinyint(8);comment:'状态：1-正常，2-禁用，3-锁定'" json:"status"`
	table.Entity
}

// 表名
func (s Staff) TableName() string {
	return s.NameWithPrefix("staff")
}
