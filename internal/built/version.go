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
// time   : 2020-05-14 10:58
// version: 1.0.0
// desc   : 版本信息

package built

import (
	"fmt"
	"gox-simple/util"
	"sync"
)

// 标记常量，服务名称
const (
	Name = "Simple"
)

// 版本和构建时间，编译时注入
var (
	once sync.Once
	// 版本
	Version = "dev"
	// 构建时间
	BuildDate = "unknown-build-date"
)

// 项目启动时加载系统信息
var (
	// 全名
	FullName string
	// 全名和版本
	FullNameWithBuildDate string
	// 启动时间
	StartedAt string
)

// 项目启动时，生成系统信息
func init() {
	once.Do(func() {
		FullName = fmt.Sprintf("%s-%s", Name, Version)
		FullNameWithBuildDate = fmt.Sprintf("%s_%s", FullName, BuildDate)
		StartedAt = util.Format(util.Now())
	})
}
