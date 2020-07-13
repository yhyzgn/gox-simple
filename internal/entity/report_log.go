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
// time   : 2020-05-15 20:30
// version: 1.0.0
// desc   : 上报日志表

package entity

import "gox-simple/internal/entity/table"

type ReportLog struct {
	ID            uint64 `gorm:"column:id;type:bigint(20);primary_key;auto_increment;comment:'主键'" json:"id"`
	ReportID      uint64 `gorm:"column:report_id;type:bigint(20);comment:'上报记录ID'" json:"reportId"`
	ReporterCode  string `gorm:"column:reporter_code;type:varchar(255);comment:'上报人code'" json:"reporterCode"`
	ReporterName  string `gorm:"column:reporter_name;type:varchar(255);comment:'上报人姓名'" json:"reporterName"`
	TodayBackWork uint   `gorm:"column:today_back_work;type:int(11);comment:'当日返岗人数'" json:"todayBackWork"`
	table.Entity
}

func (r ReportLog) TableName() string {
	return r.NameWithPrefix("report_log")
}
