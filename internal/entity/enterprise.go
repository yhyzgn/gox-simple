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
// time   : 2020-05-15 10:37
// version: 1.0.0
// desc   : 企业信息，只允许这里的企业上报数据

package entity

import "gox-simple/internal/entity/table"

// 标准化的企业数据
// 达到上报标准的企业们
type Enterprise struct {
	ID            uint64 `gorm:"column:id;type:bigint(20);primary_key;auto_increment;comment:'主键';" json:"id"`
	Level         uint8  `gorm:"column:level;type:tinyint(8);default:0;comment:'级别：${level}星级酒店，${level}A级景区，0无等级';" json:"level"`
	EmployeeTotal uint   `gorm:"column:employee_total;type:int(11);comment:'员工总数';" json:"employeeTotal"`
	EnterpriseBasic
	table.Entity
}

func (e Enterprise) TableName() string {
	return e.NameWithPrefix("enterprise")
}
