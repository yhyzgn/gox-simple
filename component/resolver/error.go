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
// time   : 2020-05-15 21:54
// version: 1.0.0
// desc   : 全局异常处理器

package resolver

import (
	"gox-simple/internal/res"
	"net/http"
)

type ErrorResolver struct {
}

func NewErrorResolver() *ErrorResolver {
	return new(ErrorResolver)
}

func (er *ErrorResolver) Resolve(err error, writer http.ResponseWriter) interface{} {
	if err == nil {
		return res.Failed("未知错误")
	}
	// 默认为error类型
	return res.Failed(err.(error).Error())
}
