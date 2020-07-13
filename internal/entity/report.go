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
// time   : 2020-05-15 20:26
// version: 1.0.0
// desc   : 最新上报数据表

package entity

import "gox-simple/internal/entity/table"

type Report struct {
	ID             uint64 `gorm:"column:id;type:bigint(20);primary_key;auto_increment;comment:'主键'" json:"id"`
	EnterpriseID   uint64 `gorm:"column:enterprise_id;type:bigint(20);comment:'企业ID'"`
	AtWork         bool   `gorm:"column:at_work;type:tinyint(1);comment:'是否已复工：0未复工，1已复工'" json:"atWork"`
	EmployeeAtWork uint   `gorm:"column:employee_at_work;type:int(11);comment:'已复工的员工数量';" json:"employeeAtWork"`
	table.Entity
}

func (r Report) TableName() string {
	return r.NameWithPrefix("report")
}
