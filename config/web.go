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
// time   : 2020-05-13 18:17
// version: 1.0.0
// desc   : 配置文件

package config

import "gox-simple/internal/db"

type Web struct {
	Server Server    `yaml:"server"`
	MySQL  db.Config `yaml:"mysql"`
}

type Server struct {
	Port        int    `yaml:"port"`
	ContextPath string `yaml:"context-path"`
}
