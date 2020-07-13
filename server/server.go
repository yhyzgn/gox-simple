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
// time   : 2020-05-13 18:13
// version: 1.0.0
// desc   : 服务启动

package server

import (
	"fmt"
	"github.com/yhyzgn/gox"
	"gox-simple/config"
	"gox-simple/internal/built"
	"gox-simple/internal/logger"
	"net/http"
)

// 运行服务
func Run(filename string) {
	// 创建路由
	x := gox.NewGoX()

	// 读取配置文件
	web := &config.Web{}
	err := x.Load(filename, web)
	if err != nil {
		panic(err)
	}

	// 初始化数据库
	initDB(web.MySQL)

	// 配置路由
	initRouter(x, web.Server.ContextPath)

	// 启动服务
	svr := &http.Server{
		Addr:    fmt.Sprintf(":%d", web.Server.Port),
		Handler: x,
	}

	logger.InfoF("Reporter [{}] started at [{}] which built at [{}].", built.Version, built.StartedAt, built.FullNameWithBuildDate)

	x.Run(svr)
}
