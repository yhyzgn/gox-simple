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
// time   : 2020-05-14 10:27
// version: 1.0.0
// desc   : 数据库实体基类

package table

import (
	"time"
)

type Entity struct {
	CreatTime       time.Time `gorm:"column:create_time;type:datetime;comment:'创建时间'" json:"createTime"`            // 创建时间
	CreatorId       string    `gorm:"column:creator_id;size:255;comment:'创建操作账号ID'" json:"creatorId"`               // 创建操作账号ID
	CreatorName     string    `gorm:"column:creator_name;size:255;comment:'创建操作账号'" json:"creatorName"`             // 创建操作账号
	DeleteTime      time.Time `gorm:"column:delete_time;type:datetime;comment:'删除时间'" json:"deleteTime"`            // 删除时间
	DeleterId       string    `gorm:"column:deleter_id;size:255;comment:'删除操作账号ID'" json:"deleterId"`               // 删除操作账号ID
	DeleterName     string    `gorm:"column:deleter_name;size:255;comment:'删除操作账号'" json:"deleterName"`             // 删除操作账号
	IsDeleted       bool      `gorm:"column:is_deleted;type:tinyint(1);comment:'删除标识 0正常 1删除'" json:"isDeleted"`    // 删除标识
	LastUpdateTime  time.Time `gorm:"column:last_update_time;type:datetime;comment:'最后操作时间'" json:"lastUpdateTime"` // 最后操作时间
	LastUpdateId    string    `gorm:"column:last_updater_id;size:255;comment:'最后操作账号ID'" json:"lastUpdaterId"`      // 最后操作账号ID
	LastUpdaterName string    `gorm:"column:last_updater_name;size:255;comment:'最后操作账号'" json:"lastUpdaterName"`    // 最后操作账号
}

func (Entity) NameWithPrefix(name string) string {
	return prefix + name
}
